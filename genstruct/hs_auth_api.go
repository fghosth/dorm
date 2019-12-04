package genstruct

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type HsAuthApi struct {
	Id        int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	ApiKey    string `dormCol:"api_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	Name      string `dormCol:"name" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	Type      int8   `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
	CreatedAt string `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt string `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt string `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}
