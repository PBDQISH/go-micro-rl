package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id      int    //用户编号
	Name    string `gorm:size:32;unique`   //用户名
	Pwdhash string `gorm:"size:128"`       //用户密码（加密后）
	Mobile  string `gorm:"size:11;unique"` //手机号
}

var GlobalConn *gorm.DB

func DbInit() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/go_micro_user?parseTime=True&loc=Local")
	if err == nil {
		GlobalConn = db
		GlobalConn.DB().SetMaxIdleConns(10)
		GlobalConn.DB().SetConnMaxLifetime(100)
		db.SingularTable(true)
		db.AutoMigrate(new(User))
		return db, nil
	}
	return nil, err
}
