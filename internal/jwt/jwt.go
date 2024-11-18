package jwt

import (
	"context"
	"fmt"
	"github.com/aridae/gophermart-diploma/internal/model"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Service struct {
	secretKeyProvider func(ctx context.Context) []byte
}

func NewService(
	secretKeyProvider func(ctx context.Context) []byte,
) *Service {
	return &Service{secretKeyProvider: secretKeyProvider}
}

func (s *Service) GenerateToken(ctx context.Context, user model.User) (string, error) {
	claims := jwtv5.MapClaims{
		"sub": user.Login,
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)

	signedTokenString, err := token.SignedString(s.secretKeyProvider(ctx))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedTokenString, nil
}

func (s *Service) ParseToken(ctx context.Context, tokenString string) (model.User, error) {
	claims := jwtv5.MapClaims{}

	token, err := jwtv5.ParseWithClaims(tokenString, &claims, func(token *jwtv5.Token) (interface{}, error) { return s.secretKeyProvider(ctx), nil })
	if err != nil {
		return model.User{}, fmt.Errorf("failed to parse token: %w", err)
	}

	login, err := token.Claims.GetSubject()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get subject from token claims: %w", err)
	}

	return model.User{Login: login}, nil
}
