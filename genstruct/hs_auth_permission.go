package genstruct

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type HsAuthPermission struct{
 		Id int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
 		AppKey string  `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
 		ApiKey string  `dormCol:"api_key" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
 		CreatedAt string  `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
 		UpdatedAt string  `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
 		DeletedAt string  `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
 		StatusAt int8  `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
   }
   