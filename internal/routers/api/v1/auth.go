package v1

import (
	"NewApp/global"
	"NewApp/internal/service"
	"NewApp/pkg/app"
	"NewApp/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// @Summary 登陆
// @Produce  json
// @Param user query string false "用户名" maxlength(20)
// @Param password query string false "密码" Emaxlength(20)
// @Param eamil query string false "邮箱" Emaxlength(20)
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

	response.ToResponse(gin.H{
		"user":  param.User,
		"token": token,
	})
}

// @Summary 登出
// @Produce  json
// @Param user query string false "用户名" maxlength(20)
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/loginOut [post]
func LoginOut(c *gin.Context) {
	return
}

// @Summary 注册
// @Produce  json
// @Param user query string false "用户名" maxlength(20)
// @Param password query string false "密码" Emaxlength(20)
// @Param eamil query string false "邮箱" Emaxlength(20)
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/register [post]
func Register(c *gin.Context) {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	//检查email是否存在
	email := &service.CheckEmail{Email: param.Email}
	if err := svc.CheckEmail(email); err != nil {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorCheckEmailFail.Msg(),
			"code": errcode.ErrorCheckEmailFail.Code(),
		})
		return
	}
	//检查用户是否存在
	user := &service.CheckUser{User: param.User}
	if err := svc.CheckUser(user); err != nil {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorCheckUserFail.Msg(),
			"code": errcode.ErrorCheckUserFail.Code(),
		})
		return
	}
	//获取邮箱验证码
	iCode, ok := global.ICache.Get(param.Email)
	if !ok {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorEmailCodeNonValid.Msg(),
			"code": errcode.ErrorEmailCodeNonValid.Code(),
		})
		return
	}
	//判断邮箱验证码
	if param.Code != iCode.(string) {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorEmailCodeFail.Msg(),
			"code": errcode.ErrorEmailCodeFail.Code(),
		})
		return
	}
	//验证码校验通过，删除邮箱验证码
	global.ICache.Delete(param.Email)
	createUser := &service.CreateUserRequest{
		User:     param.User,
		Password: param.Password,
		Email:    param.Email,
		ParentID: param.ParentID,
	}
	//创建用户
	if err := svc.CreateUser(createUser); err != nil {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorCreateUserFail.Msg(),
			"code": errcode.ErrorCreateUserFail.Code(),
		})
		return
	}
	//返回成功
	response.ToResponse(gin.H{
		"data": gin.H{
			"user": param.User,
		},
		"msg":  "success",
		"code": 0,
	})
	return
}

// @Summary 获取邮箱验证码
// @Produce  json
// @Param eamil query string false "邮箱" Emaxlength(20)
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/get/emailCode [post]
func GetEmailCode(c *gin.Context) {
	param := &service.EmailCodeRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	email := &service.CheckEmail{Email: param.Email}
	if err := svc.CheckEmail(email); err != nil {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorCheckEmailFail.Msg(),
			"code": errcode.ErrorCheckEmailFail.Code(),
		})
	}
	if err := svc.GetEmailCode(param); err != nil {
		response.ToResponse(gin.H{
			"data": gin.H{},
			"msg":  errcode.ErrorGetEmailCodeFail.Msg(),
			"code": errcode.ErrorGetEmailCodeFail.Code(),
		})
	}
	response.ToResponse(gin.H{
		"data": gin.H{},
		"msg":  "success",
		"code": 0,
	})
	return
}
