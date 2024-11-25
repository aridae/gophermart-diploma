package httpapi

import (
	"encoding/json"
	"net/http"

	oapispec "github.com/aridae/gophermart-diploma/internal/transport/http/http-api/oapi-spec"
)

func (s *APIService) PostUserBalanceWithdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	apiReq := oapispec.PostUserBalanceWithdrawJSONBody{}

	err := json.NewDecoder(r.Body).Decode(&apiReq)
	if err != nil {
		mustPresentJSONErrorWithCode(err, w, http.StatusBadRequest)
		return
	}

	domainReq := mapAPIToDomainWithdrawalRequest(apiReq)

	err = s.requestWithdrawalHandler.Handle(ctx, domainReq)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
