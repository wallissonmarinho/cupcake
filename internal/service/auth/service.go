package auth

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/wallissonmarinho/cupcake/internal/domain"
	"github.com/wallissonmarinho/cupcake/internal/repository/profile"
	"github.com/wallissonmarinho/cupcake/internal/repository/user"
	"github.com/wallissonmarinho/cupcake/internal/util"
)

type AuthService interface {
	Register(ctx context.Context, Request domain.UserProfileRequest) error
	Authenticate(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	userRepository    user.UserRepository
	profileRepository profile.ProfileRepository
	logger            *log.Logger
}

func NewAuthService(logger *log.Logger, db *sqlx.DB) AuthService {
	return &authService{
		userRepository:    user.NewUserRepository(db),
		profileRepository: profile.NewProfileRepository(db),
		logger:            logger,
	}
}

func (s *authService) Register(ctx context.Context, req domain.UserProfileRequest) (err error) {
	req.User.ID = uuid.New().String()
	req.ID = uuid.New().String()
	req.UserID = req.User.ID

	hashedPassword, err := util.HashPassword(req.User.Password)
	if err != nil {
		s.logger.Printf("error: failed to hash password: %v", err)
		return err
	}
	req.User.Password = hashedPassword

	err = s.userRepository.CreateUser(req.User)
	if err != nil {
		s.logger.Printf("error: failed to create user: %v", err)
		return err
	}

	err = s.profileRepository.CreateProfile(req)
	if err != nil {
		s.logger.Printf("error: failed to create profile: %v", err)
		return err
	}

	s.logger.Printf("info: user and profile created successfully")
	return nil
}

func (s *authService) Authenticate(ctx context.Context, email, password string) (string, error) {
	// TODO: implement authentication logic
	return "", nil
}
