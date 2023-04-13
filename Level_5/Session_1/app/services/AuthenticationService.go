package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"idstar.com/session8/app/dtos"
	"idstar.com/session8/app/repositories"
	"idstar.com/session8/app/tools"
)

const key = "abcdefghij1234567890"

type AuthenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
	}
}

func (u *AuthenticationService) Login(dto dtos.LoginRequest) (int, *dtos.Token, error) {
	user, err := u.userRepository.FindByUsernameOrEmail(dto.Username)

	if err != nil {
		return 500, nil, err
	}

	if user.Id == "" {
		return 404, nil, errors.New("user not found")
	}

	aes128 := tools.Aes128{}
	decryptedPassword, err := aes128.Decrypt(user.Password)
	if err != nil {
		return 500, nil, err
	}

	if dto.Password != *decryptedPassword {
		return 401, nil, errors.New("invalid password")
	}

	token, err := GenerateToken(user.Username)
	if err != nil {
		return 500, nil, err
	}

	return 200, token, nil
}

func GenerateToken(username string) (*dtos.Token, error) {
	iat := jwt.NewNumericDate(time.Now())
	exp := jwt.NewNumericDate(time.Now().Add(time.Millisecond * 3600000))

	claims := &dtos.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "issuer",
			Audience:  []string{"audience01", "audience02"},
			IssuedAt:  iat,
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	return &dtos.Token{
		Value:     tokenString,
		IssuedOn:  time.Unix(iat.Unix(), 0),
		ExpiresOn: time.Unix(exp.Unix(), 0),
	}, nil
}
