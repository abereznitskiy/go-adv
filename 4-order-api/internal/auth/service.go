package auth

import (
	"fmt"
	"go-adv/4-order-api/configs"
	"go-adv/4-order-api/internal/user"
	"go-adv/4-order-api/pkg/code"
	"go-adv/4-order-api/pkg/crypto"
	"go-adv/4-order-api/pkg/jwt"
	"go-adv/4-order-api/pkg/models"

	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository *user.UserRepository
	Config         *configs.Config
}

func NewAuthService(authRepository *user.UserRepository, config *configs.Config) *AuthService {
	return &AuthService{
		UserRepository: authRepository,
		Config:         config,
	}
}

func (service *AuthService) Login(phoneNumber string) (string, error) {
	existedUser, err := service.UserRepository.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	sessionId, err := crypto.GenerateSessionId()
	if err != nil {
		return "", err
	}

	if existedUser != nil {
		updatedUser, err := service.UserRepository.Update(&models.User{
			Model:       gorm.Model{ID: existedUser.ID},
			PhoneNumber: existedUser.PhoneNumber,
			SessionId:   sessionId,
			Code:        code.GenerateCode(),
		})
		if err != nil {
			return "", err
		}

		return updatedUser.SessionId, nil
	}

	createdUser, err := service.UserRepository.Create(&models.User{
		PhoneNumber: phoneNumber,
		SessionId:   sessionId,
		Code:        code.GenerateCode(),
	})
	if err != nil {
		return "", err
	}

	return createdUser.SessionId, nil
}

func (service *AuthService) Verify(sessionId, code string) (string, error) {
	existedUser, err := service.UserRepository.GetBySessionId(sessionId)
	if err != nil {
		return "", err
	}

	if code == existedUser.Code {
		token, err := jwt.NewJWT(service.Config.Db.Secret).Create(jwt.JWTData{
			PhoneNumber: existedUser.PhoneNumber,
			SessionId:   existedUser.SessionId,
			Code:        existedUser.Code,
			Id:          fmt.Sprint(existedUser.ID),
		})
		if err != nil {
			return "", err
		}
		return token, nil
	}

	return "", nil
}
