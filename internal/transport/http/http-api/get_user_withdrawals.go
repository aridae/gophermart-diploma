package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/aridae/gophermart-diploma/pkg/slice"
)

func (s *ApiService) GetUserWithdrawals(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
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

	if len(apiWithdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
}
