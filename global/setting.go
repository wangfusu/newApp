package global

import (
	"NewApp/pkg/logger"
	"NewApp/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS   //服务设置
	AppSetting      *setting.AppSettingS      //app设置
	DatabaseSetting *setting.DatabaseSettingS //数据库设置
	Logger          *logger.Logger            //日志设置
	JWTSetting      *setting.JWTSettingS      //校验设置
	EmailSetting    *setting.EmailSettingS    //邮件设置
)
