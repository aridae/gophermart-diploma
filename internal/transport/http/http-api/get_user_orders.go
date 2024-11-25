package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/aridae/gophermart-diploma/pkg/slice"
)

func (s *APIService) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
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

	if len(apiOrders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
}
