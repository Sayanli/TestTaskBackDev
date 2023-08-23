package service

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/Sayanli/TestTaskBackDev/internal/domain"
	"github.com/Sayanli/TestTaskBackDev/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Tokens struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type UserService struct {
	repo   repository.User
	secret string
}

func NewUserService(repo repository.User, secret string) *UserService {
	return &UserService{
		repo:   repo,
		secret: secret,
	}
}

func (s *UserService) CreateUser(ctx context.Context, guid string) (Tokens, error) {
	if guid == "" {
		return Tokens{}, domain.ErrEmptyGuid
	}
	err := s.checkDublicateUser(ctx, guid)
	if err != nil {
		return Tokens{}, err
	}
	tokens, err := s.generateTokens(guid)
	if err != nil {
		return Tokens{}, err
	}
	refreshTokenHash, err := s.hashRefreshToken(tokens.RefreshToken)
	if err != nil {
		return Tokens{}, err
	}
	user := domain.User{
		Guid:         guid,
		RefreshToken: refreshTokenHash,
	}
	err = s.repo.Create(ctx, user)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (s *UserService) RefreshToken(ctx context.Context, guid string, refreshToken string) (Tokens, error) {
	if guid == "" {
		return Tokens{}, domain.ErrEmptyGuid
	}
	user, err := s.repo.GetByGuid(ctx, guid)
	if err != nil {
		return Tokens{}, domain.ErrUserNotFound
	}
	if !s.verifyRefreshTokens(refreshToken, user.RefreshToken) {
		return Tokens{}, domain.ErrInvalidRefreshToken
	}
	tokens, err := s.generateTokens(guid)
	if err != nil {
		return Tokens{}, err
	}

	refreshTokenHash, err := s.hashRefreshToken(tokens.RefreshToken)
	if err != nil {
		return Tokens{}, err
	}
	user.RefreshToken = refreshTokenHash
	err = s.repo.RefreshToken(ctx, user)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (s *UserService) checkDublicateUser(ctx context.Context, guid string) error {
	flag, err := s.repo.CheckDublicateUser(ctx, guid)
	if err != nil {
		return err
	}
	if flag {
		return domain.ErrUserDuplicate
	}
	return nil
}

func (s *UserService) generateTokens(guid string) (Tokens, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"guid": guid,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		return Tokens{}, err
	}

	refreshToken := uuid.New()
	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))

	tokens := Tokens{
		AccessToken:  tokenString,
		RefreshToken: refreshTokenBase64,
	}

	return tokens, nil
}

func (s *UserService) hashRefreshToken(refreshToken string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 14)
	return string(bytes), err
}

func (s *UserService) verifyRefreshTokens(refreshToken string, hashedRefreshToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedRefreshToken), []byte(refreshToken))
	return err == nil
}
