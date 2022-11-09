package dbmodel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB // 全局操作数据库的对象

func DataBase(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println(err)
		panic("database connect ERROR")
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       // 表名字不加s
	db.DB().SetMaxIdleConns(20)  // 连接池
	db.DB().SetMaxOpenConns(100) // 最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
