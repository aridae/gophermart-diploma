package httpapi

import (
	"encoding/json"
	oapispec "github.com/aridae/gophermart-diploma/internal/transport/http/http-api/oapi-spec"
	"net/http"
)

func (s *ApiService) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginReq := oapispec.PostUserLoginJSONBody{}

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		mustPresentJSONErrorWithCode(err, w, http.StatusBadRequest)
		return
	}

	domainReq := mapAPIToDomainLoginRequest(loginReq)

	loginResp, err := s.loginUserHandler.Handle(ctx, domainReq)
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}

	w.Header().Add("Authorization", loginResp.JWT)
	_, err = w.Write([]byte(loginResp.JWT))
	if err != nil {
		mustPresentJSONError(err, w)
		return
	}
}
