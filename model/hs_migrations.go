package model

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	sqlHsMigrations    string
	argsHsMigrations   []interface{}
	dbconnHsMigrations *sql.DB
	driverHsMigrations string
)

type HsMigrations struct {
	Id        int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	Migration string `dormCol:"migration" dormMysqlType:"varchar(255)" dorm:"NOT NULL"`
	Batch     int32  `dormCol:"batch" dormMysqlType:"int(11)" dorm:"NOT NULL"`
}

//返回执行语句后sql，调试用
func (hsMigrations HsMigrations) GetSql() (string, []interface{}) {
	return sqlHsMigrations, argsHsMigrations
}

//设置db
func (hsMigrations HsMigrations) SetDBConn(db, str string) {
	var err error
	driverHsMigrations = db
	switch db {
	case "mysql":
		dbconnHsMigrations, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "mariadb":
		dbconnHsMigrations, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "cockroachDB":
		dbconnHsMigrations, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "postgresql":
		dbconnHsMigrations, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	}
}

func NewHsMigrations() HsMigrations {
	dbconnHsMigrations = DB
	driverHsMigrations = Driver
	return HsMigrations{}
}

//获得args字符串(除了update)
func getHsMigrationsArgsStr(num int) string {
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
func getHsMigrationsArgsStrUpdate() string {
	var argsStr string
	switch driverHsAuthApplication {
	case "mysql":
		argsStr = "migration=?,batch=? WHERE id=?"
	case "mariadb":
		argsStr = "migration=?,batch=? WHERE id=?"
	case "cockroachDB":
		argsStr = "migration=$1,batch=$2 WHERE id=$3"
	case "postgresql":
		argsStr = "migration=$1,batch=$2 WHERE id=$3"
	}
	return argsStr
}

func (hsMigrations HsMigrations) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select id,migration,batch from hs_migrations " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsMigrations = sqlstr
	argsHsMigrations = value
	rows, err := dbconnHsMigrations.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsMigrations.Id
	values[1] = &hsMigrations.Migration
	values[2] = &hsMigrations.Batch
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = hsMigrations
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}

func (hsMigrations *HsMigrations) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	argsStr := getHsMigrationsArgsStr(1)
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,migration,batch FROM hs_migrations WHERE id = " + argsStr
	sqlHsMigrations = sqlstr
	argsHsMigrations = args
	rows, err := dbconnHsMigrations.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsMigrations, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsMigrations.Id
	values[1] = &hsMigrations.Migration
	values[2] = &hsMigrations.Batch
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return hsMigrations, err
}

func (hsMigrations HsMigrations) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	argsStr := getHsMigrationsArgsStr(2)
	sqlstr := "INSERT INTO hs_migrations (migration,batch) VALUES (" + argsStr + ")"

	stmtIns, err := dbconnHsMigrations.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 2)
	args[0] = &hsMigrations.Migration
	args[1] = &hsMigrations.Batch

	sqlHsMigrations = sqlstr
	argsHsMigrations = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}

func (hsMigrations HsMigrations) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	argsStr := getHsMigrationsArgsStr(2)
	sqlstr := "INSERT INTO hs_migrations (migration,batch) VALUES (" + argsStr + ")"
	tx, err := dbconnHsMigrations.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 2)

	sqlHsMigrations = sqlstr
	argsHsMigrations = args

	for _, value := range obj {
		v := value.(HsMigrations)
		args[0] = v.Migration
		args[1] = v.Batch

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

func (hsMigrations *HsMigrations) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	argsStr := getHsMigrationsArgsStrUpdate()
	sqlstr := "UPDATE hs_migrations SET " + argsStr
	stmtIns, err := dbconnHsMigrations.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 3)
	args[0] = &hsMigrations.Migration
	args[1] = &hsMigrations.Batch
	args[2] = &hsMigrations.Id
	sqlHsMigrations = sqlstr
	argsHsMigrations = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}

func (hsMigrations HsMigrations) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	argsStr := getHsMigrationsArgsStrUpdate()
	sqlstr := "UPDATE hs_migrations SET " + argsStr
	tx, err := dbconnHsMigrations.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 3)

	for _, value := range obj {
		v := value.(HsMigrations)
		args[0] = v.Migration
		args[1] = v.Batch
		args[2] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsMigrations = sqlstr
	argsHsMigrations = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}

func (hsMigrations HsMigrations) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
	argsStr := getHsMigrationsArgsStr(1)
	sqlstr := "DELETE FROM hs_migrations WHERE id = " + argsStr
	stmt, err := dbconnHsMigrations.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = hsMigrations.Id
	sqlHsMigrations = sqlstr
	argsHsMigrations = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}

func (hsMigrations HsMigrations) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	argsStr := getHsMigrationsArgsStr(1)
	sqlstr := "DELETE FROM hs_migrations WHERE id = " + argsStr
	tx, err := dbconnHsMigrations.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(HsMigrations)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsMigrations = sqlstr
	argsHsMigrations = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}

func (hsMigrations HsMigrations) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := dbconnHsMigrations.Prepare(sql)
	Checkerr(err)

	sqlHsMigrations = sql
	argsHsMigrations = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}
