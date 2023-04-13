package services

import (
	"errors"

	"idstar.com/session8/app/dtos"
	"idstar.com/session8/app/repositories"
	"idstar.com/session8/app/tools"
)

type AuthenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
	}
}

func (u *AuthenticationService) Login(dto dtos.LoginRequest) (int, error) {
	user, err := u.userRepository.FindByUsernameOrEmail(dto.Username)

	if err != nil {
		return 500, err
	}

	if user.Id == "" {
		return 404, errors.New("user not found")
	}

	aes128 := tools.Aes128{}
	decryptedPassword, err := aes128.Decrypt(user.Password)
	if err != nil {
		return 500, err
	}

	if dto.Password != *decryptedPassword {
		return 401, errors.New("invalid password")
	}

	return 200, nil
}
