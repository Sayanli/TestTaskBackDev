package service

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/Sayanli/TestTaskBackDev/internal/entity"
	"github.com/Sayanli/TestTaskBackDev/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   repository.Auth
	secret string
}

func NewAuthService(repo repository.Auth, secret string) *AuthService {
	return &AuthService{
		repo:   repo,
		secret: secret,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, guid string) (entity.Token, error) {
	if guid == "" {
		return entity.Token{}, entity.ErrEmptyGuid
	}
	err := s.isUserExists(ctx, guid)
	if err != nil {
		return entity.Token{}, err
	}
	tokens, err := s.generateTokens(guid)
	if err != nil {
		return entity.Token{}, err
	}
	refreshTokenHash, err := s.hashRefreshToken(tokens.RefreshToken)
	if err != nil {
		return entity.Token{}, err
	}
	user := entity.User{
		Guid:         guid,
		RefreshToken: refreshTokenHash,
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return entity.Token{}, err
	}

	return tokens, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, guid string, refreshToken string) (entity.Token, error) {
	if guid == "" {
		return entity.Token{}, entity.ErrEmptyGuid
	}
	user, err := s.repo.GetByGuid(ctx, guid)
	if err != nil {
		return entity.Token{}, entity.ErrUserNotFound
	}
	if !s.verifyRefreshTokens(refreshToken, user.RefreshToken) {
		return entity.Token{}, entity.ErrInvalidRefreshToken
	}
	tokens, err := s.generateTokens(guid)
	if err != nil {
		return entity.Token{}, err
	}

	refreshTokenHash, err := s.hashRefreshToken(tokens.RefreshToken)
	if err != nil {
		return entity.Token{}, err
	}
	user.RefreshToken = refreshTokenHash
	err = s.repo.RefreshToken(ctx, user)
	if err != nil {
		return entity.Token{}, err
	}

	return tokens, nil
}

func (s *AuthService) isUserExists(ctx context.Context, guid string) error {
	flag, err := s.repo.IsUserExists(ctx, guid)
	if err != nil {
		return err
	}
	if flag {
		return entity.ErrUserDuplicate
	}
	return nil
}

func (s *AuthService) generateTokens(guid string) (entity.Token, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"guid": guid,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		return entity.Token{}, err
	}

	refreshToken := uuid.New()
	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))

	tokens := entity.Token{
		AccessToken:  tokenString,
		RefreshToken: refreshTokenBase64,
	}

	return tokens, nil
}

func (s *AuthService) hashRefreshToken(refreshToken string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 14)
	return string(bytes), err
}

func (s *AuthService) verifyRefreshTokens(refreshToken string, hashedRefreshToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedRefreshToken), []byte(refreshToken))
	return err == nil
}
