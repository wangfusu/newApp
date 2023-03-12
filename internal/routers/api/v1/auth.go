package v1

import (
	"NewApp/global"
	"NewApp/internal/service"
	"NewApp/pkg/app"
	"NewApp/pkg/errcode"
	"github.com/gin-gonic/gin"
	"log"
)

// @Summary 登陆
// @Produce  json
// @Param name query string false "用户名" maxlength(20)
// @Param password query string false "密码" Emaxlength(20)
// @Success 200 {object} model.Login "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.User, param.Password)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	temp := []int{}
	log.Println(temp[1])

	response.ToResponse(gin.H{
		"user":  param.User,
		"token": token,
	})
}

// @Summary 登出
// @Produce  json
// @Param name query string false "用户名" maxlength(20)
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/loginOut [post]
func LoginOut(c *gin.Context) {
	return
}
