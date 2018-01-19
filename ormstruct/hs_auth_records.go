package ormstruct

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var sqlHsAuthRecords string
var argsHsAuthRecords []interface{}
var dbconnHsAuthRecords *sql.DB

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

//返回执行语句后sql，调试用
func (hsAuthRecords HsAuthRecords) GetSql() (string, []interface{}) {
	return sqlHsAuthRecords, argsHsAuthRecords
}

//设置db
func (hsAuthRecords HsAuthRecords) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconnHsAuthRecords, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconnHsAuthRecords, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconnHsAuthRecords, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func NewHsAuthRecords() HsAuthRecords {
	dbconnHsAuthRecords = DB
	return HsAuthRecords{}
}

func (hsAuthRecords HsAuthRecords) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select id,secret_key,app_key,sign,token,alg,ip,exp,iat,type,created_at,updated_at,deleted_at,status_at from hs_auth_records " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = value
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
		ar[num] = hsAuthRecords
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}

func (hsAuthRecords *HsAuthRecords) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,secret_key,app_key,sign,token,alg,ip,exp,iat,type,created_at,updated_at,deleted_at,status_at FROM hs_auth_records WHERE id = ?"
	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args
	rows, err := DB.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsAuthRecords, err
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
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return hsAuthRecords, err
}

func (hsAuthRecords HsAuthRecords) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO hs_auth_records (secret_key,app_key,sign,token,alg,ip,exp,iat,type,created_at,updated_at,deleted_at,status_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 13)
	args[0] = &hsAuthRecords.SecretKey
	args[1] = &hsAuthRecords.AppKey
	args[2] = &hsAuthRecords.Sign
	args[3] = &hsAuthRecords.Token
	args[4] = &hsAuthRecords.Alg
	args[5] = &hsAuthRecords.Ip
	args[6] = &hsAuthRecords.Exp
	args[7] = &hsAuthRecords.Iat
	args[8] = &hsAuthRecords.Type
	args[9] = &hsAuthRecords.CreatedAt
	args[10] = &hsAuthRecords.UpdatedAt
	args[11] = &hsAuthRecords.DeletedAt
	args[12] = &hsAuthRecords.StatusAt

	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}

func (hsAuthRecords HsAuthRecords) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	sqlstr := "INSERT INTO hs_auth_records (secret_key,app_key,sign,token,alg,ip,exp,iat,type,created_at,updated_at,deleted_at,status_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 13)

	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args

	for _, value := range obj {
		v := value.(HsAuthRecords)
		args[0] = v.SecretKey
		args[1] = v.AppKey
		args[2] = v.Sign
		args[3] = v.Token
		args[4] = v.Alg
		args[5] = v.Ip
		args[6] = v.Exp
		args[7] = v.Iat
		args[8] = v.Type
		args[9] = v.CreatedAt
		args[10] = v.UpdatedAt
		args[11] = v.DeletedAt
		args[12] = v.StatusAt

		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.AddBatch); i++ { //后置hooks
		Afterfun.AddBatch[i]()
	}

	return err
}

func (hsAuthRecords *HsAuthRecords) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE hs_auth_records SET secret_key=?,app_key=?,sign=?,token=?,alg=?,ip=?,exp=?,iat=?,type=?,created_at=?,updated_at=?,deleted_at=?,status_at=? where id=?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 14)
	args[0] = &hsAuthRecords.SecretKey
	args[1] = &hsAuthRecords.AppKey
	args[2] = &hsAuthRecords.Sign
	args[3] = &hsAuthRecords.Token
	args[4] = &hsAuthRecords.Alg
	args[5] = &hsAuthRecords.Ip
	args[6] = &hsAuthRecords.Exp
	args[7] = &hsAuthRecords.Iat
	args[8] = &hsAuthRecords.Type
	args[9] = &hsAuthRecords.CreatedAt
	args[10] = &hsAuthRecords.UpdatedAt
	args[11] = &hsAuthRecords.DeletedAt
	args[12] = &hsAuthRecords.StatusAt
	args[13] = &hsAuthRecords.Id
	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}

func (hsAuthRecords HsAuthRecords) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE hs_auth_records SET secret_key=?,app_key=?,sign=?,token=?,alg=?,ip=?,exp=?,iat=?,type=?,created_at=?,updated_at=?,deleted_at=?,status_at=? where id=?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 14)

	for _, value := range obj {
		v := value.(HsAuthRecords)
		args[0] = v.SecretKey
		args[1] = v.AppKey
		args[2] = v.Sign
		args[3] = v.Token
		args[4] = v.Alg
		args[5] = v.Ip
		args[6] = v.Exp
		args[7] = v.Iat
		args[8] = v.Type
		args[9] = v.CreatedAt
		args[10] = v.UpdatedAt
		args[11] = v.DeletedAt
		args[12] = v.StatusAt
		args[13] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}

func (hsAuthRecords HsAuthRecords) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
	sqlstr := "DELETE FROM hs_auth_records WHERE id = ?"
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = hsAuthRecords.Id
	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}

func (hsAuthRecords HsAuthRecords) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM hs_auth_records WHERE id = ?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(HsAuthRecords)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthRecords = sqlstr
	argsHsAuthRecords = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}

func (hsAuthRecords HsAuthRecords) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlHsAuthRecords = sql
	argsHsAuthRecords = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}
