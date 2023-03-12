package service

import "errors"

type AuthRequest struct {
	User     string `form:"user" binding:"max=20"`
	Password string `json:"password" binding:"max=20"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.User, param.Password)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
