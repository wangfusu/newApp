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
	r.Use(middleware.Tracing())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json") //swagger路由 指向这个url
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //静态访问
	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("/login", v1.Login)
		//apiv1.POST("/loginOut", login)
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/get/emailCode", v1.GetEmailCode)
	}
	user := v1.User{}
	userV1 := r.Group("api/v1")
	userV1.Use(middleware.JWT())
	{
		userV1.POST("/user", user.Create)
		userV1.DELETE("/user/:id", user.Delete)
		userV1.PUT("/user/:id", user.Update)
		userV1.PATCH("/user/:id/state", user.Update)
		userV1.GET("/user", user.List)
	}
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	return r
}
