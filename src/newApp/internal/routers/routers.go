package routers

import (
	_ "NewApp/docs"
	"NewApp/global"
	"NewApp/internal/middleware"
	"NewApp/internal/routers/api"
	v1 "NewApp/internal/routers/api/v1"
	"NewApp/pkg/limiter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.ServerSetting.ContextTimeout * time.Second))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json") //swagger路由 指向这个url
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //静态访问
	login := v1.Login{}
	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("/login", login.Login)
		apiv1.POST("/loginOut", login.LoginOut)
		r.POST("/auth", v1.GetAuth)
		//apiv1.POST("/register", login.Register)
	}
	user := v1.User{}
	userv1 := r.Group("api/v1")
	userv1.Use(middleware.JWT())
	{
		userv1.POST("/user", user.Create)
		userv1.DELETE("/user/:id", user.Delete)
		userv1.PUT("/user/:id", user.Update)
		userv1.PATCH("/user/:id/state", user.Update)
		userv1.GET("/user", user.List)
	}
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	return r
}
