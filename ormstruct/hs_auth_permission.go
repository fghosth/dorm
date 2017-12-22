package ormstruct

import (
	_ "github.com/go-sql-driver/mysql"
)

type HsAuthPermission struct {
	Id        int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	AppKey    string `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	ApiKey    string `dormCol:"api_key" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	CreatedAt int32  `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt int32  `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt int32  `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}
type HsAuthApplication struct {
	Id        int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	SecretKey string `dormCol:"secret_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	AppKey    string `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	Name      string `dormCol:"name" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	Ip        string `dormCol:"ip" dormMysqlType:"varchar(32)" dorm:"NOT NULL;DEFAULT ''"`
	Type      int8   `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
	Exp       int32  `dormCol:"exp" dormMysqlType:"int(11)" dorm:"NOT NULL;DEFAULT '0'"`
	CreatedAt int32  `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt int32  `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt int32  `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}
