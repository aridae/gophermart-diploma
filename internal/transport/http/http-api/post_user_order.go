package httpapi

import (
	submitorder "github.com/aridae/gophermart-diploma/internal/usecases/submit-order"
	"io"
	"net/http"
)

func (s *ApiService) PostUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	apiOrderNum, err := io.ReadAll(r.Body)
	if err != nil {
		mustPresentJSONErrorWithCode(err, w, http.StatusBadRequest)
		return
	}

	err = s.submitOrderHandler.Handle(ctx, submitorder.Request{OrderNumber: string(apiOrderNum)})
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
