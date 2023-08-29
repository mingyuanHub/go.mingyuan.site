package setting

import (
	"gopkg.in/ini.v1"
	"strings"
)

type appConfig struct {
	RunMode   string
	Host      string
	LocalHost string
}

type serverConfig struct {
	ServerPort   int
	ReadTimeout  int
	WriteTimeout int
}

type mysqlConfig struct {
	MysqlHost     string
	MysqlPort     int
	MysqlUser     string
	MysqlPassword string
	MysqlDb       string
}

var (
	AppConfig    *appConfig
	ServerConfig *serverConfig
	MysqlConfig  *mysqlConfig
)

func Init() error {
	cfg, err := ini.Load("./internal/conf/app.ini")
	if err != nil {
		return err
	}

	AppConfig = &appConfig{
		RunMode: cfg.Section("").Key("RUN_MODE").String(),
	}

	ServerConfig = &serverConfig{
		ServerPort:   cfg.Section("server").Key("HTTP_PORT").MustInt(9080),
		ReadTimeout:  cfg.Section("server").Key("READ_TIMEOUT").MustInt(60),
		WriteTimeout: cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60),
	}

	MysqlConfig = &mysqlConfig{
		MysqlHost:     cfg.Section("mysql").Key("MYSQL_HOST").String(),
		MysqlPort:     cfg.Section("mysql").Key("MYSQL_PORT").MustInt(3306),
		MysqlUser:     cfg.Section("mysql").Key("MYSQL_USER").String(),
		MysqlPassword: cfg.Section("mysql").Key("MYSQL_PASSWORD").String(),
		MysqlDb:       cfg.Section("mysql").Key("MYSQL_DB").String(),
	}

	AppConfig.Host = cfg.Section("host").Key("HOST_" + strings.ToUpper(AppConfig.RunMode)).String()

	//AppConfig.LocalHost = fmt.Sprintf("http://127.0.0.1:%d", ServerConfig.ServerPort)

	AppConfig.LocalHost = "http://www.mingyuan.site"

	return nil
}
