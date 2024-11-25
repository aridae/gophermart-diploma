package httpapi

import (
	"io"
	"net/http"

	submitorder "github.com/aridae/gophermart-diploma/internal/usecases/submit-order"
)

func (s *APIService) PostUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	apiOrderNum, err := io.ReadAll(r.Body)
	if err != nil {
		mustPresentJSONErrorWithCode(err, w, http.StatusBadRequest)
		return
	}

	resp, err := s.submitOrderHandler.Handle(ctx, submitorder.Request{OrderNumber: string(apiOrderNum)})
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	httpStatusCode := mapResponseCode(resp.Code)
	w.WriteHeader(httpStatusCode)
}

func mapResponseCode(code submitorder.ResponseCode) int {
	switch code {
	case submitorder.OrderNumberAlreadyLoadedByThisOwner:
		return http.StatusOK
	case submitorder.OrderNumberAccepted:
		return http.StatusAccepted
	}

	return http.StatusOK
}
