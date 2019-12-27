package ormstruct

type ResourceAction struct {
	Id            string `dormCol:"id" dormCOCKDBType:"UUID" dorm:"NOT NULL DEFAULT gen_random_uuid()"`
	StatusAt      int8   `dormCol:"status_at" dormCOCKDBType:"BIT(5)" dorm:"NOT NULL DEFAULT 1"`
	Data       []byte    `dormCol:"data" dormMysqlType:"mediumblob" dorm:"NOT NULL"`
}
