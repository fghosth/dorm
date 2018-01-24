package genstruct

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type HsAuthApplication struct{
 		Id int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
 		SecretKey string  `dormCol:"secret_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
 		AppKey string  `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
 		Name string  `dormCol:"name" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
 		Ip string  `dormCol:"ip" dormMysqlType:"varchar(32)" dorm:"NOT NULL;DEFAULT ''"`
 		Type int8  `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
 		Exp int64  `dormCol:"exp" dormMysqlType:"int(11)" dorm:"NOT NULL;DEFAULT '0'"`
 		CreatedAt string  `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
 		UpdatedAt string  `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
 		DeletedAt string  `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
 		StatusAt int8  `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
   }
   