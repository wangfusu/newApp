package main

import (
	"NewApp/global"
	"NewApp/internal/model"
	"NewApp/internal/routers"
	"NewApp/pkg/logger"
	"NewApp/pkg/setting"
	"NewApp/pkg/tracer"
	"NewApp/utils"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port         string
	runMode      string
	config       string
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	setupFlag()

	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer, err:%v", err)
	}
	utils.InitCache()
}

// @title NewAPP
// @version 1.0
// @description Go BASE APP
// @termsOfService https://github.com/wangfusu
func main() {
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.LisetenAndService err : %v", err)
		}
	}()
	// 等待中断信号
	quit := make(chan os.Signal)
	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server ...")
	// 最大事件控制，用于通知该服务端它有5秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DBEngine.AutoMigrate(model.Model{})
	if err1 := global.DBEngine.AutoMigrate(&model.User{}).Error; err1 != nil {
		c := context.Background()
		global.Logger.Error(c, err1)
	}
	return nil
}
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"newApp-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
	return nil
}
