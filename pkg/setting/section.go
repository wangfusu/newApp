package setting

import "time"

type ServerSettingS struct {
	RunMode        string        //运行模式
	HttpPort       string        //运行端口
	ReadTimeout    time.Duration //读超时 时间
	WriteTimeout   time.Duration //写超时时间
	ContextTimeout time.Duration //上下文超时时间
}
type AppSettingS struct {
	DefaultPageSize      int      //默认分页大小
	MaxPageSize          int      //最大分页大小
	LogSavePath          string   //日志保存路径
	LogFileName          string   //日志文件名称
	LogFileExt           string   //日志文件后缀名
	UploadSavePath       string   //上传文件的最终保存目录
	UploadServerUrl      string   //上传文件后的用于展示的文件服务地址。
	UploadImageMaxSize   int      //上传文件所允许的最大空间大小（MB）
	UploadImageAllowExts []string //上传文件所允许的文件后缀。
}
type DatabaseSettingS struct {
	DBType       string //数据库类型
	UserName     string //账号
	Password     string //密码
	Host         string //数据库地址
	DBName       string //数据库名
	TablePrefix  string //
	Charset      string //数据库文本格式
	ParseTime    bool   //
	MaxIdleConns int    //
	MaxOpenConns int    //
}

type JWTSettingS struct {
	Secret string        //密钥
	Issuer string        //名称
	Expire time.Duration //鉴权时间
}

type EmailSettingS struct {
	Host     string   //ip地址
	Port     int      //端口
	UserName string   //账号
	Password string   //密码（一般指代邮箱的对接的密钥）
	IsSSL    bool     //是否ssl
	From     string   //from名称
	To       []string //接收邮件用户
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadALLSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
