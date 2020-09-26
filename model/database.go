package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	connection *gorm.DB
	err        error
)

func InitConnection() {
	connection, err = gorm.Open("mysql", "root:root@/gin_database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据连接异常", err)
	}
	//defer connection.Close()
	connection.AutoMigrate(&User{})

	connection.DB().SetMaxIdleConns(10)
	connection.DB().SetMaxOpenConns(100)
	connection.DB().SetConnMaxLifetime(10 * time.Second)
}
