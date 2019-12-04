package genstruct

import "time"

type Bom struct{
 		Id int  `dormCol:"id" dormMysqlType:"bigint(20)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
 		ReportId int  `dormCol:"report_id" dormMysqlType:"bigint(20)" dorm:"unsigned;NOT NULL"`
 		Mid int  `dormCol:"mid" dormMysqlType:"bigint(20)" dorm:"unsigned;NOT NULL"`
 		Qty float64  `dormCol:"qty" dormMysqlType:"double" dorm:"NOT NULL"`
 		CreateTime time.Time  `dormCol:"create_time" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT current_timestamp() ON UPDATE current_timestamp()"`
 		ModifyTime time.Time  `dormCol:"modify_time" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT '0000-00-00 00:00:00'"`
 		IsDeleted int8  `dormCol:"is_deleted" dormMysqlType:"tinyint(1)" dorm:"NOT NULL;DEFAULT 0"`
   }
   