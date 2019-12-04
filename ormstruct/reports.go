package genstruct

import "time"

type Reports struct{
 		Id int  `dormCol:"id" dormMysqlType:"bigint(20)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
 		MpId int  `dormCol:"mp_ID" dormMysqlType:"bigint(20)" dorm:"unsigned;NOT NULL"`
 		Category int8  `dormCol:"category" dormMysqlType:"tinyint(4)" dorm:"NOT NULL"`
 		Cid int  `dormCol:"cid" dormMysqlType:"bigint(20)" dorm:"unsigned;NOT NULL"`
 		Scost int  `dormCol:"scost" dormMysqlType:"bigint(20)" dorm:"unsigned;NOT NULL"`
 		Stime time.Time  `dormCol:"stime" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT current_timestamp() ON UPDATE current_timestamp()"`
 		Etime time.Time  `dormCol:"etime" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT '0000-00-00 00:00:00'"`
 		IsDeleted int8  `dormCol:"is_deleted" dormMysqlType:"tinyint(1)" dorm:"NOT NULL;DEFAULT 0"`
 		CreateTime time.Time  `dormCol:"create_time" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT '0000-00-00 00:00:00'"`
 		ModifyTime time.Time  `dormCol:"modify_time" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT '0000-00-00 00:00:00'"`
 		Ecost int  `dormCol:"ecost" dormMysqlType:"bigint(20)" dorm:"unsigned;NOT NULL"`
 		SQty float64  `dormCol:"s_QTY" dormMysqlType:"double" dorm:"NOT NULL"`
 		EQty float64  `dormCol:"e_QTY" dormMysqlType:"double" dorm:"NOT NULL"`
 		ReciptQty float64  `dormCol:"recipt_QTY" dormMysqlType:"double" dorm:"NOT NULL"`
 		IncreaseQty float64  `dormCol:"increase_QTY" dormMysqlType:"double" dorm:"NOT NULL"`
 		DecreaseQty float64  `dormCol:"decrease_QTY" dormMysqlType:"double" dorm:"NOT NULL"`
 		Bom string  `dormCol:"bom longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'bom 物料详情' CHECK (json_valid(bom" dormMysqlType:"longtext" dorm:"NOT NULL"`
 		CheckIncrease float64  `dormCol:"check_increase" dormMysqlType:"double" dorm:"NOT NULL"`
 		CheckDecrease float64  `dormCol:"check_decrease" dormMysqlType:"double" dorm:"NOT NULL"`
 		ScrapQty float64  `dormCol:"scrap_qty" dormMysqlType:"double" dorm:"NOT NULL"`
 		SoldQty float64  `dormCol:"sold_qty" dormMysqlType:"double" dorm:"NOT NULL"`
   }
   