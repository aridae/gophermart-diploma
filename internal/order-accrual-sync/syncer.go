package orderaccrualsync

import (
	"context"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/downstream/accrual"
	"github.com/aridae/gophermart-diploma/internal/logger"
	"github.com/aridae/gophermart-diploma/internal/model"
	orderdb "github.com/aridae/gophermart-diploma/internal/repos/order-db"
)

var (
	nonTerminalOrderStatuses = []model.OrderStatus{model.OrderStatusNew, model.OrderStatusProcessing}
)

type orderAccrualService interface {
	GetOrderByNumber(ctx context.Context, orderNumber string) (accrual.Order, error)
}

type ordersService interface {
	Search(ctx context.Context, filter orderdb.Filter, pagination orderdb.Pagination) ([]model.Order, error)
	UpdateOrder(ctx context.Context, orderNumber string, setters ...orderdb.Setter) error
}

type Syncer struct {
	orderAccrualService orderAccrualService
	ordersService       ordersService
	interval            time.Duration
	workersPoolSize     int
}

func New(
	orderAccrualService orderAccrualService,
	ordersService ordersService,
	interval time.Duration,
	workersPoolSize int,
) *Syncer {
	return &Syncer{
		orderAccrualService: orderAccrualService,
		ordersService:       ordersService,
		workersPoolSize:     workersPoolSize,
		interval:            interval,
	}
}

func (s *Syncer) Run(ctx context.Context) {
	ticker := time.NewTicker(s.interval)

	syncQueue := make(chan model.Order, s.workersPoolSize)
	go func() {
		<-ctx.Done()
		logger.Errorf("stopping Syncer due to context cancellation: %v", ctx.Err())
		close(syncQueue)
	}()

	for i := range s.workersPoolSize {
		go runSyncerWorker(ctx, syncQueue, i, s.syncOrder)
	}

	for {
		select {
		case <-ctx.Done():
			logger.Infof("[orderaccrualsync.Run] exiting syncer due to context cancel")
		case <-ticker.C:
			if err := s.runSyncing(ctx, syncQueue); err != nil {
				logger.Errorf("orderaccrualsync.Syncer sync failed: %v", err)
			}
		}
	}
}

func (s *Syncer) runSyncing(ctx context.Context, queue chan model.Order) error {
	ordersToSync, err := s.loadOrders(ctx, orderdb.Filter{Statuses: nonTerminalOrderStatuses})
	if err != nil {
		return fmt.Errorf("loadOrders: %w", err)
	}

	for _, orderToSync := range ordersToSync {
		queueOrderToSync(ctx, queue, orderToSync)
	}

	return nil
}

func (s *Syncer) loadOrders(ctx context.Context, filter orderdb.Filter) ([]model.Order, error) {
	limit := 100
	orders := make([]model.Order, 0)

	for page := 1; ; page++ {
		ordersPage, err := s.ordersService.Search(ctx, filter, orderdb.Pagination{Page: page, Limit: limit})
		if err != nil {
			return nil, fmt.Errorf("ordersService.Search: %w", err)
		}

		orders = append(orders, ordersPage...)

		if len(ordersPage) < limit {
			break
		}
	}

	return orders, nil
}

func (s *Syncer) syncOrder(ctx context.Context, order model.Order) error {
	logger.Infof("[orderaccrualsync.syncOrder] syncing order number %s", order.Number)
	orderAccrual, err := s.orderAccrualService.GetOrderByNumber(ctx, order.Number)
	if err != nil {
		return fmt.Errorf("orderAccrual.GetOrderByNumber: %w", err)
	}

	err = s.ordersService.UpdateOrder(ctx, order.Number,
		orderdb.SetOrderStatus(orderAccrual.Status),
		orderdb.SetOrderAccrual(orderAccrual.Accrual),
	)
	if err != nil {
		return fmt.Errorf("ordersService.UpdateOrder: %w", err)
	}

	return nil
}

func queueOrderToSync(ctx context.Context, queue chan<- model.Order, order model.Order) {
	select {
	case queue <- order:
	case <-ctx.Done():
		logger.Infof("terminating queueOrderToSync due to context cancel")
	}
}

func runSyncerWorker(ctx context.Context, queue <-chan model.Order, workerNumber int, syncFn func(ctx context.Context, order model.Order) error) {
	logger.Infof("starting syncer worker #%d", workerNumber)
	for order := range queue {
		if err := syncFn(ctx, order); err != nil {
			logger.Errorf("syncer worker #%d failed: %v", workerNumber, err)
		}
	}
}
