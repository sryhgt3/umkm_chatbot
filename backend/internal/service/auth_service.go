package service

import (
	"context"
	"errors"
	"umkm-chatbot/internal/model"
	"umkm-chatbot/internal/repository"
	"umkm-chatbot/internal/utils"
)

type AuthService interface {
	RegisterStoreOwner(ctx context.Context, name, email, password, storeName string) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetMe(ctx context.Context, userID int) (*model.User, error)
}

type authService struct {
	userRepo  repository.UserRepository
	storeRepo repository.StoreRepository
	jwtSecret string
}

func NewAuthService(userRepo repository.UserRepository, storeRepo repository.StoreRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		storeRepo: storeRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) RegisterStoreOwner(ctx context.Context, name, email, password, storeName string) (*model.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetUserByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create Store
	store := &model.Store{Name: storeName}
	if err := s.storeRepo.CreateStore(ctx, store); err != nil {
		return nil, err
	}

	// Create User
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     model.RoleStoreOwner,
		StoreID:  &store.ID,
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Role, user.StoreID, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) GetMe(ctx context.Context, userID int) (*model.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}
