package lexer

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid        int32   `dormCol:"uid" dormMysqlType:"int(11)" dorm:"PRIMARY;UNSIGNED;NOT-NULL;AUTO_INCREMENT;COMMENT-'自增id号'"`
	Name       string  `dormCol:"name" dormMysqlType:"varchar(64)" dorm:"NOT-NULL"`
	Gender     int8    `dormCol:"gender" dormMysqlType:"bit(1)" dorm:"NOT-NULL"`
	Password   string  `dormCol:"password" dormMysqlType:"varchar(255)" dorm:"NOT-NULL"`
	Qq         int8    `dormCol:"qq" dormMysqlType:"bit(8)" dorm:"NOT-NULL"`
	Account    string  `dormCol:"account" dormMysqlType:"char(255)" dorm:"NOT-NULL"`
	Cellphone  string  `dormCol:"cellphone" dormMysqlType:"varchar(255)" dorm:"NOT-NULL"`
	Happy      string  `dormCol:"happy" dormMysqlType:"enum('aaa','bbb','ccc','ddd')" dorm:"NOT-NULL;DEFAULT-'aaa'"`
	Cash       float64 `dormCol:"cash" dormMysqlType:"double" dorm:"NOT-NULL;DEFAULT-0"`
	CreateTime int32   `dormCol:"create_time" dormMysqlType:"timestamp" dorm:"NOT-NULL"`
	UpdateTime string  `dormCol:"update_time" dormMysqlType:"datetime" dorm:"NOT-NULL"`
}
