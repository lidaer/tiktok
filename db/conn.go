package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	// 连接数据库
	DB, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

// DBConn : 返回数据库连接对象
func DBConn() *gorm.DB {
	return DB
}
