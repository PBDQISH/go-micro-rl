package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func RegisterUser(mobile, pwd string) error {
	var user User
	user.Name = mobile
	md5 := md5.New()
	md5.Write([]byte(pwd))
	PwDHASH := hex.EncodeToString(md5.Sum(nil))
	user.Pwdhash = PwDHASH
	user.Mobile = mobile

	return GlobalConn.Create(&user).Error
}

func Login(mobile, pwd string) (string, error) {
	var user User
	md5 := md5.New()
	md5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(md5.Sum(nil))
	err := GlobalConn.Where("mobile=?", mobile).Select("name,id").Where("pwdhash=?", pwd_hash).Find(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return user.Name, err
}
