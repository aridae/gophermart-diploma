package http_api

import (
	"encoding/json"
	"github.com/aridae/gophermart-diploma/pkg/slice"
	"net/http"
)

func (s *ApiService) GetUserWithdrawals(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	domainWithdrawals, err := s.getWithdrawalsHistoryHandler.Handle(ctx)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	apiWithdrawals := slice.MapBatchNoErr(domainWithdrawals, mapDomainToAPIWithdrawal)

	err = json.NewEncoder(w).Encode(apiWithdrawals)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}
}
