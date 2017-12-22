package ormstruct

   import (
   	"database/sql"
   	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   )
   type ProductInformation struct{
 		ProductId int64  `dormCol:"product_id" dormCOCKDBType:"INT" dorm:"NOT NULL"`
 		ProductName string  `dormCol:"product_name" dormCOCKDBType:"STRING(50)" dorm:"NOT NULL"`
 		ProductDescription string  `dormCol:"product_description" dormCOCKDBType:"STRING(2000)" dorm:"NULL"`
 		CategoryId string  `dormCol:"category_id" dormCOCKDBType:"STRING(1)" dorm:"NOT NULL"`
 		WeightClass int64  `dormCol:"weight_class" dormCOCKDBType:"INT" dorm:"NULL"`
 		WarrantyPeriod int64  `dormCol:"warranty_period" dormCOCKDBType:"INT" dorm:"NULL"`
 		SupplierId int64  `dormCol:"supplier_id" dormCOCKDBType:"INT" dorm:"NULL"`
 		ProductStatus string  `dormCol:"product_status" dormCOCKDBType:"STRING(20)" dorm:"NULL"`
 		ListPrice float64  `dormCol:"list_price" dormCOCKDBType:"DECIMAL(8,2)" dorm:"NULL"`
 		MinPrice float64  `dormCol:"min_price" dormCOCKDBType:"DECIMAL(8,2)" dorm:"NULL"`
 		CatalogUrl string  `dormCol:"catalog_url" dormCOCKDBType:"STRING(50)" dorm:"NULL"`
 		DateAdded string  `dormCol:"date_added" dormCOCKDBType:"DATE" dorm:"NULL DEFAULT '2017-12-20':::DATE"`
   }
   