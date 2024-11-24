package httpapi

import (
	"encoding/json"
	oapispec "github.com/aridae/gophermart-diploma/internal/transport/http/http-api/oapi-spec"
	"net/http"
)

func (s *ApiService) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	signupReq := oapispec.PostUserRegisterJSONBody{}

	err := json.NewDecoder(r.Body).Decode(&signupReq)
	if err != nil {
		mustPresentJSONErrorWithCode(err, w, http.StatusBadRequest)
		return
	}

	domainReq := mapAPIToDomainRegisterRequest(signupReq)

	signupResp, err := s.registerUserHandler.Handle(ctx, domainReq)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	w.Header().Add("Authorization", signupResp.JWT)
	_, err = w.Write([]byte(signupResp.JWT))
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}
}
