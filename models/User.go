package models

import (
	"DataCertPlatform/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID int `form:"id"`
	phone string `phone`
	password string `form:"password"`
}
/**
 *将用户的信息保存到数据库中
 */

func (u User) AddUser (int64,error){
	//脱敏
	hashMD5 := md5.New()
	hashMD5.Write([]byte(u.password))
	pwdBytes := hashMD5.Sum(nil)
	u.password = hex.EncodeToString(pwdBytes)//把脱敏的密码MD5值重新赋值为密码进行存储
	rs,err := db_mysql.Db.Exec("insert into user(phone,password) values(?,?)",u.phone,u.password)
	//错误早发现早解决
	if err != nil{//保存数据遇到错误
		return -1,err
	}
	id, err := rs.RowsAffected()
	if err != nil{//保存数据遇到错误
		return -1,err
	}
	//id值代表的是此次数据操作影响的行数，id是一个整数int64类型
	return id,nil
}
