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
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
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

func (svc *Service) CountUser(param *CountUserRequest) (int, error) {
	return svc.dao.CountUser(param.Name, param.State)
}

func (svc *Service) GetUserList(param *UserListRequest, pager *app.Pager) ([]*model.User, error) {
	return svc.dao.GetUserList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateUser(param *UpdateUserRequest) error {
	return svc.dao.UpdateUser(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteUser(param *DeleteUserRequest) error {
	return svc.dao.DeleteUser(param.ID)
}
