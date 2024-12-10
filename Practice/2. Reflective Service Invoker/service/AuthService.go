package service

import (
	"context"
	"test/errors"
	"test/models"
)

type AuthService struct {
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	if req.Username == "admin" && req.Password == "password" {
		return &models.LoginResponse{Message: "Login successful", Status: "OK"}, nil
	}
	return nil, errors.ErrInvalidCredentials
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.RegisterResponse, error) {
	return &models.RegisterResponse{Message: "Registered successful", Status: "OK"}, nil
}
