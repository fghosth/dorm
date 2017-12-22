package ormstruct

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type HsMigrations struct{
 		Id int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
 		Migration string  `dormCol:"migration" dormMysqlType:"varchar(255)" dorm:"NOT NULL"`
 		Batch int32  `dormCol:"batch" dormMysqlType:"int(11)" dorm:"NOT NULL"`
   }
   