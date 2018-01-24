package genstruct

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type HsMigrations struct{
 		Id int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
 		Migration string  `dormCol:"migration" dormMysqlType:"varchar(255)" dorm:"NOT NULL"`
 		Batch int64  `dormCol:"batch" dormMysqlType:"int(11)" dorm:"NOT NULL"`
   }
   