package lexer

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid        int32   `dormCol:"uid" dormMysqlType:"int(11)" dorm:"PRIMARY;UNSIGNED;NOT-NULL;AUTO_INCREMENT;COMMENT-'自增id号'"`
	Name       string  `dormCol:"name" dormMysqlType:"varchar(64)" dorm:"NOT-NULL"`
	Gender     int8    `dormCol:"gender" dormMysqlType:"bit(1)" dorm:"NOT-NULL"`
	Password   string  `dormCol:"password" dormMysqlType:"varchar(255)" dorm:"NOT-NULL"`
	Qq         int8    `dormCol:"qq" dormMysqlType:"bit(8)" dorm:"NOT-NULL"`
	Account    string  `dormCol:"account" dormMysqlType:"char(255)" dorm:"NOT-NULL"`
	Cellphone  string  `dormCol:"cellphone" dormMysqlType:"varchar(255)" dorm:"NOT-NULL"`
	Happy      string  `dormCol:"happy" dormMysqlType:"enum('aaa','bbb','ccc','ddd')" dorm:"NOT-NULL;DEFAULT-'aaa'"`
	Cash       float64 `dormCol:"cash" dormMysqlType:"double" dorm:"NOT-NULL;DEFAULT-0"`
	CreateTime int32   `dormCol:"create_time" dormMysqlType:"timestamp" dorm:"NOT-NULL"`
	UpdateTime string  `dormCol:"update_time" dormMysqlType:"datetime" dorm:"NOT-NULL"`
}

type ProductInformation struct {
	ProductId          int64   `dormCol:"product_id" dormMysqlType:"INT" dorm:"PRIMARY KEY NOT NULL"`
	ProductName        string  `dormCol:"product_name" dormMysqlType:"STRING(50)" dorm:"UNIQUE NOT NULL"`
	ProductDescription string  `dormCol:"product_description" dormMysqlType:"STRING(2000)" dorm:""`
	CategoryId         string  `dormCol:"category_id" dormMysqlType:"STRING(1)" dorm:"NOT NULL CHECK (category_id IN ('A''B''C')) "`
	WeightClass        int64   `dormCol:"weight_class" dormMysqlType:"INT" dorm:"DEFAULT 3 "`
	WarrantyPeriod     int64   `dormCol:"warranty_period" dormMysqlType:"INT" dorm:"CONSTRAINT valid_warranty CHECK (warranty_period BETWEEN 0 AND 24)"`
	SupplierId         int64   `dormCol:"supplier_id" dormMysqlType:"INT" dorm:""`
	ProductStatus      string  `dormCol:"product_status" dormMysqlType:"STRING(20)" dorm:""`
	ListPrice          float64 `dormCol:"list_price" dormMysqlType:"DECIMAL(8,2)" dorm:""`
	MinPrice           float64 `dormCol:"min_price" dormMysqlType:"DECIMAL(8,2)" dorm:""`
	CatalogUrl         string  `dormCol:"catalog_url" dormMysqlType:"STRING(50)" dorm:"UNIQUE"`
	DateAdded          string  `dormCol:"date_added" dormMysqlType:"DATE" dorm:"DEFAULT CURRENT_DATE()"`
}
