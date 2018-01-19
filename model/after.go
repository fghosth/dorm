
package model
import (
 "database/sql"
 "log"
 "strconv"
_ "github.com/go-sql-driver/mysql"
)


var sqlAfter string
var argsAfter []interface{}
var dbconnAfter *sql.DB

type After struct {
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
func (after After) GetSql() (string, []interface{}) {
	return sqlAfter, argsAfter
}

//设置db
func (after After) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconnAfter, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconnAfter, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconnAfter, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func NewAfter() After {
	dbconnAfter = DB
	return After{}
}


func (after After) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select  from after " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlAfter = sqlstr
	argsAfter = value
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
		ar[num] = after
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}
	

func (after *After) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT  FROM after WHERE  = ?"
	sqlAfter = sqlstr
	argsAfter = args
	rows, err := DB.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return after, err
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
	return after, err
}
	

func (after After) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO after () VALUES ()"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, -1)
	sqlAfter = sqlstr
	argsAfter = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}
	

func (after After) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	sqlstr := "INSERT INTO after () VALUES ()"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, -1)

	sqlAfter = sqlstr
	argsAfter = args

	for _, value := range obj {
		v := value.(After)
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


func (after *After) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE after SET  where =?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 0)
	sqlAfter = sqlstr
	argsAfter = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}


func (after After) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE after SET  where =?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 0)

	for _, value := range obj {
		v := value.(After)
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlAfter = sqlstr
	argsAfter = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}


func (after After) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
  sqlstr := "DELETE FROM after WHERE  = ?"
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = after.
	sqlAfter = sqlstr
	argsAfter = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}


func (after After) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM after WHERE  = ?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(After)
		args[0] = v.
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlAfter = sqlstr
	argsAfter = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}


func (after After) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlAfter = sql
	argsAfter = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}

