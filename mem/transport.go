package ormstruct

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type Address struct{
 		Aid int32  `dormCol:"aid" dormMysqlType:"int(11)" dorm:"PRIMARY;UNSIGNED;NOT-NULL;AUTO_INCREMENT"`
 		City string  `dormCol:"city" dormMysqlType:"varchar(255)" dorm:"NOT-NULL"`
 		Privent string  `dormCol:"privent" dormMysqlType:"varchar(255)" dorm:""`
 		Address string  `dormCol:"address" dormMysqlType:"varchar(255)" dorm:"NOT-NULL"`
 		Uid int32  `dormCol:"uid" dormMysqlType:"int(11)" dorm:"UNSIGNED;NOT-NULL"`
   }
   