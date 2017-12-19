package lexer

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid        int32   `dormCol:"uid"  dorm:"UNSIGNED;NOT_NULL;AUTO_INCREMENT;COMMENT_'自增id号'"`
	Name       string  `dormCol:"name"  dorm:"NOT_NULL"`
	Gender     int8    `dormCol:"gender"  dorm:"NOT_NULL"`
	Password   string  `dormCol:"password"  dorm:"NOT_NULL"`
	Qq         int8    `dormCol:"qq"  dorm:"NOT_NULL"`
	Account    string  `dormCol:"account"  dorm:"NOT_NULL"`
	Cellphone  string  `dormCol:"cellphone"  dorm:"NOT_NULL"`
	Happy      string  `dormCol:"happy"  dorm:"NOT_NULL;DEFAULT_aaa,"`
	Cash       float64 `dormCol:"cash"  dorm:"NOT_NULL;DEFAULT_0,"`
	CreateTime int32   `dormCol:"create_time"  dorm:"NOT_NULL"`
	UpdateTime string  `dormCol:"update_time"  dorm:"NOT_NULL"`
}
