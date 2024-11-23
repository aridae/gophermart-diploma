package loginuser

import (
	"context"
	"fmt"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type userRepository interface {
	GetUserCredentials(ctx context.Context, login string) (*model.UserCredentials, error)
}

type jwtService interface {
	GenerateToken(ctx context.Context, user model.User) (string, error)
}

type Handler struct {
	userRepository userRepository
	jwtService     jwtService
}

func NewHandler(
	usersRepository userRepository,
	jwtService jwtService,
) *Handler {
	return &Handler{
		userRepository: usersRepository,
		jwtService:     jwtService,
	}
}

type Request struct {
	Login    string
	Password string
}

type Response struct {
	JWT string
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	creds, err := h.userRepository.GetUserCredentials(ctx, req.Login)
	if err != nil {
		return Response{}, fmt.Errorf("error getting user credentials: %w", err)
	}

	if creds == nil {
		return Response{}, domainerrors.InvalidUserCredentialsError()
	}

	if !creds.Equal(req.Login, req.Password) {
		return Response{}, domainerrors.InvalidUserCredentialsError()
	}

	user := model.User{Login: creds.Login}

	token, err := h.jwtService.GenerateToken(ctx, user)
	if err != nil {
		return Response{}, fmt.Errorf("failed to create JWT token: %w", err)
	}

	return Response{JWT: token}, nil
}
