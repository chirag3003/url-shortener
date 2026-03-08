package service

import (
	"context"
	"strconv"

	"github.com/chirag3003/go-backend-template/dto/response"
	"github.com/chirag3003/go-backend-template/pkg/apperror"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/rs/zerolog"
)

// UserService handles user business logic.
type UserService struct {
	userRepo repository.UserRepository
	log      zerolog.Logger
}

// NewUserService creates a new UserService.
func NewUserService(userRepo repository.UserRepository, log zerolog.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		log:      log.With().Str("service", "user").Logger(),
	}
}

// GetByID retrieves a user by their ID and returns a safe response DTO.
func (s *UserService) GetByID(ctx context.Context, userID string) (*response.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		s.log.Error().Err(err).Str("userID", userID).Msg("failed to fetch user")
		return nil, apperror.Internal(err)
	}
	if user == nil {
		return nil, apperror.NotFound("user not found")
	}

	return &response.UserResponse{
		ID:        strconv.FormatInt(user.ID, 10),
		Name:      user.Name,
		Email:     user.Email,
		PhoneNo:   user.PhoneNo,
		AvatarURL: user.AvatarURL,
	}, nil
}

// Update updates user's profile information.
func (s *UserService) Update(ctx context.Context, userID string, name, avatarURL string) (*response.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		s.log.Error().Err(err).Str("userID", userID).Msg("failed to fetch user for update")
		return nil, apperror.Internal(err)
	}
	if user == nil {
		return nil, apperror.NotFound("user not found")
	}

	if name != "" {
		user.Name = name
	}
	if avatarURL != "" {
		user.AvatarURL = avatarURL
	}

	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		s.log.Error().Err(err).Str("userID", userID).Msg("failed to update user")
		return nil, apperror.Internal(err)
	}

	return &response.UserResponse{
		ID:        strconv.FormatInt(user.ID, 10),
		Name:      user.Name,
		Email:     user.Email,
		PhoneNo:   user.PhoneNo,
		AvatarURL: user.AvatarURL,
	}, nil
}
