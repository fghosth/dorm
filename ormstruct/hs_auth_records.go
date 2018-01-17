package ormstruct

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type HsAuthRecords struct {
	Id        int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	SecretKey string `dormCol:"secret_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	AppKey    string `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	Sign      string `dormCol:"sign" dormMysqlType:"varchar(128)" dorm:"NOT NULL;DEFAULT ''"`
	Token     string `dormCol:"token" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	Alg       string `dormCol:"alg" dormMysqlType:"varchar(64)" dorm:"NOT NULL"`
	Ip        string `dormCol:"ip" dormMysqlType:"varchar(32)" dorm:"NOT NULL;DEFAULT ''"`
	Exp       string `dormCol:"exp" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	Iat       string `dormCol:"iat" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	Type      int8   `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
	CreatedAt string `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt string `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt string `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}

func (hsAuthRecords HsAuthRecords) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var ar []interface{}
	var err error
	ar = make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select id,secret_key,app_key,sign,token,alg,ip,exp,iat,type,created_at,updated_at,deleted_at,status_at from hs_auth_records " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	rows, err := DB.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthRecords.Id
	values[1] = &hsAuthRecords.SecretKey
	values[2] = &hsAuthRecords.AppKey
	values[3] = &hsAuthRecords.Sign
	values[4] = &hsAuthRecords.Token
	values[5] = &hsAuthRecords.Alg
	values[6] = &hsAuthRecords.Ip
	values[7] = &hsAuthRecords.Exp
	values[8] = &hsAuthRecords.Iat
	values[9] = &hsAuthRecords.Type
	values[10] = &hsAuthRecords.CreatedAt
	values[11] = &hsAuthRecords.UpdatedAt
	values[12] = &hsAuthRecords.DeletedAt
	values[13] = &hsAuthRecords.StatusAt
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar = append(ar, hsAuthRecords)
		num++
	}
	for i := 0; i < len(Afterfun.Select); i++ { //前置hooks
		Afterfun.Select[i]()
	}
	return ar, nil
}
