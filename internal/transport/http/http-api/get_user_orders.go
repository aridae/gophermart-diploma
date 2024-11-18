package http_api

import (
	"encoding/json"
	"github.com/aridae/gophermart-diploma/pkg/slice"
	"net/http"
)

func (s *ApiService) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	domainOrders, err := s.getOrdersHandler.Handle(ctx)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	apiOrders, err := slice.MapBatch(domainOrders, mapDomainToAPIOrder)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(apiOrders)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}
}
