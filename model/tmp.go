package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	uid         int32   `dormCol:"uid"  dorm:"UNSIGNED;NOT_NULL;AUTO_INCREMENT;COMMENT_'自增id号'"`
	name        string  `dormCol:"name"  dorm:"NOT_NULL"`
	gender      int8    `dormCol:"gender"  dorm:"NOT_NULL"`
	password    string  `dormCol:"password"  dorm:"NOT_NULL"`
	qq          int8    `dormCol:"qq"  dorm:"NOT_NULL"`
	account     string  `dormCol:"account"  dorm:"NOT_NULL"`
	cellphone   string  `dormCol:"cellphone"  dorm:"NOT_NULL"`
	happy       string  `dormCol:"happy"  dorm:"NOT_NULL;DEFAULT_aaa,"`
	cash        float64 `dormCol:"cash"  dorm:"NOT_NULL;DEFAULT_0,"`
	create_time int32   `dormCol:"create_time"  dorm:"NOT_NULL"`
	update_time string  `dormCol:"update_time"  dorm:"NOT_NULL"`
}
