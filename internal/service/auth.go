package service

import (
	"NewApp/global"
	"NewApp/utils"
	"errors"
	"time"
)

type AuthRequest struct {
	User     string `form:"user" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `json:"email" binding:"required,min=10,max=30"`
	Code     string `json:"code" binding:"required,max=6"`
}
type RegisterRequest struct {
	User     string `form:"user" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `json:"email" binding:"required,min=10,max=30"`
	Code     string `json:"code" binding:"required,max=6"`
}

type EmailCodeRequest struct {
	Email string `json:"email" binding:"required,min=10,max=30"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.User, param.Password, param.Email)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}

func (svc *Service) GetEmailCode(param *EmailCodeRequest) error {
	code := utils.RandCode()
	if err := utils.MailSendCode(param.Email, code); err != nil {
		return err
	}
	global.ICache.Set(param.Email, code, 10*time.Minute)
	return nil
}
