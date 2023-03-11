package v1

import (
	"NewApp/global"
	"NewApp/internal/service"
	"NewApp/pkg/app"
	"NewApp/pkg/convert"
	"NewApp/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type User struct {
	service service.Service
}

func NewUser() User {
	//var ctx *http.Request
	//svc := service.New(ctx.Context())
	return User{
		//service: svc,
	}
}

// @Users 用户
// @Summary 获取多个用户
// @Produce  json
// @Param usern query string false "用户名" maxlength(20)
// @Param password query string false "密码" maxlength(20)
// @Param name query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Users [get]
func (t User) List(c *gin.Context) {
	param := service.UserListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountUser(&service.CountUserRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountUserFail)
		return
	}

	Users, err := svc.GetUserList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}

	response.ToResponseList(Users, totalRows)
	return
}

// @Users 用户
// @Summary 新增用户
// @Produce  json
// @Param user body string true "用户" minlength(2) maxlength(20)
// @Param password body string false "密码" minlength(8) maxlength(20)
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Users [post]
func (t User) Create(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(gin.H{})
	return

}

// @Users 用户
// @Summary 更新用户
// @Produce  json
// @Param id path int true "用户 ID"
// @Param name body string false "用户名" minlength(2) maxlength(20)
// @Param password body string false "密码" minlength(8) maxlength(20)
// @Success 200 {array} model.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Users/{id} [put]
func (t User) Update(c *gin.Context) {
	param := service.UpdateUserRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Users 用户
// @Summary 删除标签
// @Produce  json
// @Param id path int true "用户 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Users/{id} [delete]
func (t User) Delete(c *gin.Context) {
	param := service.DeleteUserRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteUserFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
