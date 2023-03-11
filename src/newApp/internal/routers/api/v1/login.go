package v1

import (
	"NewApp/global"
	"NewApp/internal/model"
	"NewApp/pkg/app"
	"NewApp/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Login struct {
}

func NewLogin() Login {
	return Login{}
}

// @Summary 登陆
// @Produce  json
// @Param name query string false "用户名" maxlength(20)
// @Param password query string false "密码" Emaxlength(20)
// @Success 200 {object} model.Login "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/login [post]
func (l *Login) Login(c *gin.Context) {
	login := model.Login{}
	global.Logger.Info(login)
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 登出
// @Produce  json
// @Param name query string false "用户名" maxlength(20)
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/loginOut [post]
func (l Login) LoginOut(c *gin.Context) {
	return
}
