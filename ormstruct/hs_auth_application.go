package ormstruct

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	sqlHsAuthApplication    string
	argsHsAuthApplication   []interface{}
	dbconnHsAuthApplication *sql.DB
	driverHsAuthApplication string
)

type HsAuthApplication struct {
	Id        int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	SecretKey string `dormCol:"secret_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	AppKey    string `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	Name      string `dormCol:"name" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	Ip        string `dormCol:"ip" dormMysqlType:"varchar(32)" dorm:"NOT NULL;DEFAULT ''"`
	Type      int8   `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
	Exp       string `dormCol:"exp" dormMysqlType:"int(11)" dorm:"NOT NULL;DEFAULT '0'"`
	CreatedAt string `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt string `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt string `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}

//返回执行语句后sql，调试用
func (hsAuthApplication HsAuthApplication) GetSql() (string, []interface{}) {
	return sqlHsAuthApplication, argsHsAuthApplication
}

//设置db
func (hsAuthApplication HsAuthApplication) SetDBConn(db, str string) {
	var err error
	driverHsAuthApplication = db
	switch db {
	case "mysql":
		dbconnHsAuthApplication, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "mariadb":
		dbconnHsAuthApplication, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "cockroachDB":
		dbconnHsAuthApplication, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "postgresql":
		dbconnHsAuthApplication, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	}
}

func NewHsAuthApplication() HsAuthApplication {
	dbconnHsAuthApplication = DB
	driverHsAuthApplication = Driver
	return HsAuthApplication{}
}

//获得args字符串(除了update)
func getHsAuthApplicationArgsStr(num int) string {
	var argsStr string
	switch driverHsAuthApplication {
	case "mysql":
		for i := 0; i < num; i++ {
			if argsStr == "" {
				argsStr = "?"
			} else {
				argsStr = argsStr + ",?"
			}
		}
	case "mariadb":
		for i := 0; i < num; i++ {
			if argsStr == "" {
				argsStr = "?"
			} else {
				argsStr = argsStr + ",?"
			}
		}
	case "cockroachDB":
		for i := 0; i < num; i++ {
			if argsStr == "" {
				argsStr = "$" + strconv.Itoa(i+1)
			} else {
				argsStr = argsStr + ",$" + strconv.Itoa(i+1)
			}
		}
	case "postgresql":
		for i := 0; i < num; i++ {
			if argsStr == "" {
				argsStr = "$" + strconv.Itoa(i+1)
			} else {
				argsStr = argsStr + ",$" + strconv.Itoa(i+1)
			}
		}
	}
	return argsStr
}

//获得args字符串(update)
func getHsAuthApplicationArgsStrUpdate() string {
	var argsStr string
	switch driverHsAuthApplication {
	case "mysql":
		argsStr = "secret_key=?,app_key=?,name=?,ip=?,type=?,exp=?,created_at=?,updated_at=?,deleted_at=?,status_at=? WHERE id=?"
	case "mariadb":
		argsStr = "secret_key=?,app_key=?,name=?,ip=?,type=?,exp=?,created_at=?,updated_at=?,deleted_at=?,status_at=? WHERE id=?"
	case "cockroachDB":
		argsStr = "secret_key=$1,app_key=$2,name=$3,ip=$4,type=$5,exp=$6,created_at=$7,updated_at=$8,deleted_at=$9,status_at=$10 WHERE id=$11"
	case "postgresql":
		argsStr = "secret_key=$1,app_key=$2,name=$3,ip=$4,type=$5,exp=$6,created_at=$7,updated_at=$8,deleted_at=$9,status_at=$10 WHERE id=$11"
	}
	return argsStr
}

func (hsAuthApplication HsAuthApplication) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select id,secret_key,app_key,name,ip,type,exp,created_at,updated_at,deleted_at,status_at from hs_auth_application " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = value
	rows, err := dbconnHsAuthApplication.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthApplication.Id
	values[1] = &hsAuthApplication.SecretKey
	values[2] = &hsAuthApplication.AppKey
	values[3] = &hsAuthApplication.Name
	values[4] = &hsAuthApplication.Ip
	values[5] = &hsAuthApplication.Type
	values[6] = &hsAuthApplication.Exp
	values[7] = &hsAuthApplication.CreatedAt
	values[8] = &hsAuthApplication.UpdatedAt
	values[9] = &hsAuthApplication.DeletedAt
	values[10] = &hsAuthApplication.StatusAt
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = hsAuthApplication
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}

func (hsAuthApplication *HsAuthApplication) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	argsStr := getHsAuthApplicationArgsStr(1)
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,secret_key,app_key,name,ip,type,exp,created_at,updated_at,deleted_at,status_at FROM hs_auth_application WHERE id = " + argsStr
	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args
	rows, err := dbconnHsAuthApplication.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsAuthApplication, err
	}

	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthApplication.Id
	values[1] = &hsAuthApplication.SecretKey
	values[2] = &hsAuthApplication.AppKey
	values[3] = &hsAuthApplication.Name
	values[4] = &hsAuthApplication.Ip
	values[5] = &hsAuthApplication.Type
	values[6] = &hsAuthApplication.Exp
	values[7] = &hsAuthApplication.CreatedAt
	values[8] = &hsAuthApplication.UpdatedAt
	values[9] = &hsAuthApplication.DeletedAt
	values[10] = &hsAuthApplication.StatusAt
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}

	return hsAuthApplication, err
}

func (hsAuthApplication HsAuthApplication) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	argsStr := getHsAuthApplicationArgsStr(10)
	sqlstr := "INSERT INTO hs_auth_application (secret_key,app_key,name,ip,type,exp,created_at,updated_at,deleted_at,status_at) VALUES (" + argsStr + ")"

	stmtIns, err := dbconnHsAuthApplication.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 10)
	args[0] = &hsAuthApplication.SecretKey
	args[1] = &hsAuthApplication.AppKey
	args[2] = &hsAuthApplication.Name
	args[3] = &hsAuthApplication.Ip
	args[4] = &hsAuthApplication.Type
	args[5] = &hsAuthApplication.Exp
	args[6] = &hsAuthApplication.CreatedAt
	args[7] = &hsAuthApplication.UpdatedAt
	args[8] = &hsAuthApplication.DeletedAt
	args[9] = &hsAuthApplication.StatusAt

	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}

func (hsAuthApplication HsAuthApplication) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	argsStr := getHsAuthApplicationArgsStr(10)
	sqlstr := "INSERT INTO hs_auth_application (secret_key,app_key,name,ip,type,exp,created_at,updated_at,deleted_at,status_at) VALUES (" + argsStr + ")"
	tx, err := dbconnHsAuthApplication.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 10)

	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args

	for _, value := range obj {
		v := value.(HsAuthApplication)
		args[0] = v.SecretKey
		args[1] = v.AppKey
		args[2] = v.Name
		args[3] = v.Ip
		args[4] = v.Type
		args[5] = v.Exp
		args[6] = v.CreatedAt
		args[7] = v.UpdatedAt
		args[8] = v.DeletedAt
		args[9] = v.StatusAt

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

func (hsAuthApplication *HsAuthApplication) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	argsStr := getHsAuthApplicationArgsStrUpdate()
	sqlstr := "UPDATE hs_auth_application SET " + argsStr
	stmtIns, err := dbconnHsAuthApplication.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 11)
	args[0] = &hsAuthApplication.SecretKey
	args[1] = &hsAuthApplication.AppKey
	args[2] = &hsAuthApplication.Name
	args[3] = &hsAuthApplication.Ip
	args[4] = &hsAuthApplication.Type
	args[5] = &hsAuthApplication.Exp
	args[6] = &hsAuthApplication.CreatedAt
	args[7] = &hsAuthApplication.UpdatedAt
	args[8] = &hsAuthApplication.DeletedAt
	args[9] = &hsAuthApplication.StatusAt
	args[10] = &hsAuthApplication.Id
	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}

func (hsAuthApplication HsAuthApplication) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	argsStr := getHsAuthApplicationArgsStrUpdate()
	sqlstr := "UPDATE hs_auth_application SET " + argsStr
	tx, err := dbconnHsAuthApplication.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 11)

	for _, value := range obj {
		v := value.(HsAuthApplication)
		args[0] = v.SecretKey
		args[1] = v.AppKey
		args[2] = v.Name
		args[3] = v.Ip
		args[4] = v.Type
		args[5] = v.Exp
		args[6] = v.CreatedAt
		args[7] = v.UpdatedAt
		args[8] = v.DeletedAt
		args[9] = v.StatusAt
		args[10] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}

func (hsAuthApplication HsAuthApplication) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
	argsStr := getHsAuthApplicationArgsStr(1)
	sqlstr := "DELETE FROM hs_auth_application WHERE id = " + argsStr
	stmt, err := dbconnHsAuthApplication.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = hsAuthApplication.Id
	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}

func (hsAuthApplication HsAuthApplication) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	argsStr := getHsAuthApplicationArgsStr(1)
	sqlstr := "DELETE FROM hs_auth_application WHERE id = " + argsStr
	tx, err := dbconnHsAuthApplication.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(HsAuthApplication)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthApplication = sqlstr
	argsHsAuthApplication = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}

func (hsAuthApplication HsAuthApplication) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := dbconnHsAuthApplication.Prepare(sql)
	Checkerr(err)
	sqlHsAuthApplication = sql
	argsHsAuthApplication = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}
