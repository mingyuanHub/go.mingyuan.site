package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mingyuanHub/mingyuan.site/pkg/setting"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() error {
	var err error
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlConfig.MysqlUser, setting.MysqlConfig.MysqlPassword, setting.MysqlConfig.MysqlHost, setting.MysqlConfig.MysqlPort, setting.MysqlConfig.MysqlDb)

	db, err = gorm.Open("mysql", connArgs)
	if err != nil {
		return err
	}
	//设置最大空闲连接
	db.DB().SetMaxIdleConns(10)
	//设置最大连接数
	db.DB().SetMaxOpenConns(100)
	//设置连接超时时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
	// 全局禁用表名复数
	db.SingularTable(true)
	//sql日志
	db.LogMode(true)

	//ping
	err = db.DB().Ping()
	if err != nil {
		return err
	}

	return nil
}