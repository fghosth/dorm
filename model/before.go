
package model
import (
 "database/sql"
 "log"
 "strconv"
_ "github.com/go-sql-driver/mysql"
)


var sqlBefore string
var argsBefore []interface{}
var dbconnBefore *sql.DB

type Before struct {
	Select      []func()
	Update      []func()
	FindByID    []func()
	Add         []func()
	AddBatch    []func()
	UpdateBatch []func()
	Delete      []func()
	DeleteBatch []func()
	Exec        []func()
}

//返回执行语句后sql，调试用
func (before Before) GetSql() (string, []interface{}) {
	return sqlBefore, argsBefore
}

//设置db
func (before Before) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconnBefore, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconnBefore, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconnBefore, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func NewBefore() Before {
	dbconnBefore = DB
	return Before{}
}


func (before Before) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select  from before " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlBefore = sqlstr
	argsBefore = value
	rows, err := DB.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = before
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}
	

func (before *Before) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT  FROM before WHERE  = ?"
	sqlBefore = sqlstr
	argsBefore = args
	rows, err := DB.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return before, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return before, err
}
	

func (before Before) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO before () VALUES ()"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, -1)
	sqlBefore = sqlstr
	argsBefore = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}
	

func (before Before) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	sqlstr := "INSERT INTO before () VALUES ()"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, -1)

	sqlBefore = sqlstr
	argsBefore = args

	for _, value := range obj {
		v := value.(Before)
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


func (before *Before) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE before SET  where =?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 0)
	sqlBefore = sqlstr
	argsBefore = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}


func (before Before) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE before SET  where =?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 0)

	for _, value := range obj {
		v := value.(Before)
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlBefore = sqlstr
	argsBefore = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}


func (before Before) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
  sqlstr := "DELETE FROM before WHERE  = ?"
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = before.
	sqlBefore = sqlstr
	argsBefore = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}


func (before Before) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM before WHERE  = ?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(Before)
		args[0] = v.
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlBefore = sqlstr
	argsBefore = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}


func (before Before) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlBefore = sql
	argsBefore = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}

