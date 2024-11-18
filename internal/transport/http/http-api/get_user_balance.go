package http_api

import (
	"encoding/json"
	"net/http"
)

func (s *ApiService) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	domainBalance, err := s.getBalanceHandler.Handle(ctx)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	apiBalance := mapDomainToAPIBalance(domainBalance)

	err = json.NewEncoder(w).Encode(apiBalance)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}
}
