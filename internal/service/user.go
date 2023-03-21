package service

import (
	"NewApp/internal/model"
	"NewApp/pkg/app"
)

type CountUserRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UserListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateUserRequest struct {
	User     string `form:"user" binding:"required,min=2,max=100"`
	Password string `json:"password" binding:"required,ming=8,max=20"`
	Email    string `json:"email" binding:"required,min=10,max=30"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateUserRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteUserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type CheckEmail struct {
	Email string `form:"email" json:"email" binding:"required"`
}
type CheckUser struct {
	User string `json:"user"`
}

func (svc *Service) CountUser(param *CountUserRequest) (int, error) {
	return svc.dao.CountUser(param.Name, param.State)
}

func (svc *Service) GetUserList(param *UserListRequest, pager *app.Pager) ([]*model.User, error) {
	return svc.dao.GetUserList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.User, param.Password, param.Email)
}

func (svc *Service) UpdateUser(param *UpdateUserRequest) error {
	return svc.dao.UpdateUser(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteUser(param *DeleteUserRequest) error {
	return svc.dao.DeleteUser(param.ID)
}

func (svc *Service) CheckEmail(param *CheckEmail) error {
	return svc.dao.CheckEmail(param.Email)
}

func (svc *Service) CheckUser(param *CheckUser) error {
	return svc.dao.CheckUser(param.User)
}
