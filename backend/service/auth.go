package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chirag3003/go-backend-template/dto/request"
	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/pkg/auth"
	"github.com/chirag3003/go-backend-template/pkg/idgen"
	"github.com/chirag3003/go-backend-template/pkg/validate"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/rs/zerolog"
)

// AuthService handles authentication business logic.
type AuthService struct {
	userRepo   repository.UserRepository
	jwtService *auth.JWTService
	log        zerolog.Logger
}

// NewAuthService creates a new AuthService.
func NewAuthService(userRepo repository.UserRepository, jwtService *auth.JWTService, log zerolog.Logger) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
		log:        log.With().Str("service", "auth").Logger(),
	}
}

// Login authenticates a user and returns a JWT token.
func (s *AuthService) Login(ctx context.Context, req *request.LoginRequest) (*response.LoginResponse, error) {
	if err := validate.Struct(req); err != nil {
		return nil, apperror.ValidationError(err.Error())
	}

	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.log.Error().Err(err).Str("email", req.Email).Msg("failed to fetch user by email")
		return nil, apperror.Internal(err)
	}
	if user == nil {
		return nil, apperror.ErrInvalidCredentials
	}

	if !auth.VerifyPassword(req.Password, user.Hash) {
		return nil, apperror.ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(strconv.FormatInt(user.ID, 10), user.Name, user.Email, user.PhoneNo)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to generate JWT")
		return nil, apperror.Internal(err)
	}

	return &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			ID:      strconv.FormatInt(user.ID, 10),
			Name:    user.Name,
			Email:   user.Email,
			PhoneNo: user.PhoneNo,
		},
	}, nil
}

// Register creates a new user account.
func (s *AuthService) Register(ctx context.Context, req *request.RegisterRequest) error {
	if err := validate.Struct(req); err != nil {
		return apperror.ValidationError(err.Error())
	}

	existing, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.log.Error().Err(err).Str("email", req.Email).Msg("failed to check existing user")
		return apperror.Internal(err)
	}
	if existing != nil {
		return apperror.ErrUserAlreadyExists
	}

	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to hash password")
		return apperror.Internal(err)
	}

	user := &models.User{
		ID:    0,
		Name:  req.Name,
		Email: req.Email,
		Hash:  hash,
	}

	userID, err := idgen.NewID()
	if err != nil {
		s.log.Error().Err(err).Msg("failed to generate user id")
		return apperror.Internal(err)
	}
	user.ID = userID

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		s.log.Error().Err(err).Msg("failed to create user")
		return apperror.Internal(fmt.Errorf("creating user: %w", err))
	}

	return nil
}
