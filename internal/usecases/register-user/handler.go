package registeruser

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
	userrepo "github.com/aridae/gophermart-diploma/internal/repos/user-repo"
)

type userRepository interface {
	CreateUser(ctx context.Context, user model.UserCredentials, now time.Time) error
}

type jwtService interface {
	GenerateToken(ctx context.Context, user model.User) (string, error)
}

type Handler struct {
	userRepository userRepository
	jwtService     jwtService
	now            func() time.Time
}

func NewHandler(
	usersRepository userRepository,
	jwtService jwtService,
) *Handler {
	return &Handler{
		userRepository: usersRepository,
		jwtService:     jwtService,
		now:            time.Now().UTC,
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
	now := h.now()

	creds, err := model.NewUserCredentials(req.Login, req.Password)
	if err != nil {
		return Response{}, fmt.Errorf("failed to create user credentials: %w", err)
	}

	err = h.userRepository.CreateUser(ctx, creds, now)
	if err != nil {
		if errors.Is(err, userrepo.ErrLoginUniqueConstraintViolated) {
			return Response{}, domainerrors.ErrLoginAlreadyTaken(req.Login)
		}

		return Response{}, fmt.Errorf("error creating user: %w", err)
	}

	user := model.User{Login: creds.Login}

	token, err := h.jwtService.GenerateToken(ctx, user)
	if err != nil {
		return Response{}, fmt.Errorf("failed to create JWT token: %w", err)
	}

	return Response{JWT: token}, nil
}
