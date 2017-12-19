package lexer

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid        int32   `dormCol:"uid"  dorm:"UNSIGNED;NOT-NULL;AUTO_INCREMENT;COMMENT-'自增id号'"`
	Name       string  `dormCol:"name"  dorm:"NOT-NULL"`
	Gender     int8    `dormCol:"gender"  dorm:"NOT-NULL"`
	Password   string  `dormCol:"password"  dorm:"NOT-NULL"`
	Qq         int8    `dormCol:"qq"  dorm:"NOT-NULL"`
	Account    string  `dormCol:"account"  dorm:"NOT-NULL"`
	Cellphone  string  `dormCol:"cellphone"  dorm:"NOT-NULL"`
	Happy      string  `dormCol:"happy"  dorm:"NOT-NULL;DEFAULT-aaa,"`
	Cash       float64 `dormCol:"cash"  dorm:"NOT-NULL;DEFAULT-0,"`
	CreateTime int32   `dormCol:"create-time"  dorm:"NOT-NULL"`
	UpdateTime string  `dormCol:"update-time"  dorm:"NOT-NULL"`
}
