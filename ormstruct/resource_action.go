package ormstruct

type ResourceAction struct {
	Id            string `dormCol:"id" dormCOCKDBType:"UUID" dorm:"NOT NULL DEFAULT gen_random_uuid()"`
	Url           string `dormCol:"url" dormCOCKDBType:"STRING(200)" dorm:"NOT NULL"`
	OperationType int64  `dormCol:"operation_type" dormCOCKDBType:"INT" dorm:"NOT NULL"`
	BeginTime     string `dormCol:"begin_time" dormCOCKDBType:"STRING(50)" dorm:"NOT NULL"`
	EndTime       string `dormCol:"end_time" dormCOCKDBType:"STRING(50)" dorm:"NOT NULL"`
	State         int64  `dormCol:"state" dormCOCKDBType:"INT" dorm:"NOT NULL DEFAULT 0:::INT"`
	CreateAt      string `dormCol:"create_at" dormCOCKDBType:"TIMESTAMP" dorm:"NOT NULL DEFAULT now()"`
	UpdateAt      string `dormCol:"update_at" dormCOCKDBType:"TIMESTAMP" dorm:"NOT NULL DEFAULT now()"`
	StatusAt      int8   `dormCol:"status_at" dormCOCKDBType:"BIT(5)" dorm:"NOT NULL DEFAULT 1"`
}
