package dorm

const (
	Select_TPL = `
func ({{{objvar}}} {{{obj}}}) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select {{{fields}}} from {{{tableName}}} " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sql{{{obj}}} = sqlstr
	args{{{obj}}} = value
	rows, err := dbconn{{{obj}}}.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	{{#each field}}
		{{{this}}}
	{{/each}}
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = {{{objvar}}}
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}
	`
	FindByID_TPL = `
func ({{{objvar}}} *{{{obj}}}) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT {{{fields}}} FROM {{{tableName}}} WHERE {{{sqlField}}} = ?"
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	rows, err := dbconn{{{obj}}}.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return {{{objvar}}}, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	{{#each field}}
		{{{this}}}
	{{/each}}
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return {{{objvar}}}, err
}
	`

	Add_TPL = `
func ({{{objvar}}} {{{obj}}}) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO {{{tableName}}} ({{{fields}}}) VALUES ({{{parms}}})"

	stmtIns, err := dbconn{{{obj}}}.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, {{{len}}})
	{{#each field}}
		{{{this}}}
	{{/each}}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}
	`
	AddBatch_TPL = `
func ({{{objvar}}} {{{obj}}}) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	var argsStr string
	switch driverHsAuthApplication {
	case "mysql":
		argsStr = "{{{mysqlparms}}}"
	case "mariadb":
		argsStr = "{{{mysqlparms}}}"
	case "cockroachDB":
		argsStr = "{{{cockroachparms}}}"
	}
	sqlstr := "INSERT INTO {{{tableName}}} ({{{fields}}}) VALUES ("+argsStr+")"
	tx, err := dbconn{{{obj}}}.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, {{{len}}})

	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args

	for _, value := range obj {
		v := value.({{{obj}}})
		{{#each field}}
	 		{{{this}}}
		{{/each}}
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
`
	Update_TPL = `
func ({{{objvar}}} *{{{obj}}}) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE {{{tableName}}} SET {{{fields}}} where {{{sqlField}}}=?"
	stmtIns, err := dbconn{{{obj}}}.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, {{{len}}})
	{{#each field}}
		{{{this}}}
	{{/each}}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}
`
	UpdateBatch_TPL = `
func ({{{objvar}}} {{{obj}}}) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE {{{tableName}}} SET {{{fields}}} where {{{sqlField}}}=?"
	tx, err := dbconn{{{obj}}}.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, {{{len}}})

	for _, value := range obj {
		v := value.({{{obj}}})
		{{#each field}}
	 		{{{this}}}
		{{/each}}
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}
`
	Delete_TPL = `
func ({{{objvar}}} {{{obj}}}) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
  sqlstr := "DELETE FROM {{{tableName}}} WHERE {{{sqlField}}} = ?"
	stmt, err := dbconn{{{obj}}}.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = {{{objvar}}}.{{{structField}}}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}
`
	DeleteBatch_TPL = `
func ({{{objvar}}} {{{obj}}}) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM {{{tableName}}} WHERE {{{sqlField}}} = ?"
	tx, err := dbconn{{{obj}}}.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.({{{obj}}})
		args[0] = v.{{{structField}}}
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}
`
	Exec_TPL = `
func ({{{objvar}}} {{{obj}}}) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := dbconn{{{obj}}}.Prepare(sql)
	Checkerr(err)

	sql{{{obj}}} = sql
	args{{{obj}}} = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}
`
	Field_TPL = `
var (
	sql{{{obj}}} string
	args{{{obj}}} []interface{}
	dbconn{{{obj}}} *sql.DB
	driver{{{obj}}} string
)
`
	Header_TPL = `
package {{{pkname}}}
import (
	{{#each field}}
 {{{this}}}
	{{/each}}
_ "github.com/go-sql-driver/mysql"
_ "github.com/lib/pq"
)
`
	Function_TPL = `
//返回执行语句后sql，调试用
func ({{{objvar}}} {{{obj}}}) GetSql() (string, []interface{}) {
	return sql{{{obj}}}, args{{{obj}}}
}

//设置db
func ({{{objvar}}} {{{obj}}}) SetDBConn(db, str string) {
	var err error
	driver{{{obj}}} = db
	switch db {
	case "mysql":
		dbconn{{{obj}}}, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconn{{{obj}}}, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconn{{{obj}}}, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func New{{{obj}}}() {{{obj}}} {
	dbconn{{{obj}}} = DB
	driver{{{obj}}} = Driver
	return {{{obj}}}{}
}
`

	MODEL_TPL = `
package {{{pkname}}}

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	LIMIT   = 500  //默认查询条数限制
	OFFSET  = 0    //默认位置
	MAXROWS = 1000 //最多查出多少条,-1为不限制
)


var (
	DB *sql.DB //数据库连接
	Beforefun Before
	Afterfun After
	Driver string
)

func init() {
	SetConn("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
}

/*
模型的基本方法接口
*/
type Model interface {
	/*
			   根据条件查找结果集
			   @parm sql 除去select xxx,xxx from tablename 之后的东西
			   @parm value sql中?值 可以为空
			   @parm limit 显示数量
			   @parm offset 数据位置0开始
		     @return struct 集合
		     @return error 错误
	*/
	Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error)
	/*
			   根据主键查找
			   @parm id 主键
		     @return struct
		     @return error 错误
	*/
	FindByID(id int64) (interface{}, error)
	/*
			   根据自身struct内容添加
			   @parm
		     @return 返回主键id
		     @return error 错误
	*/
	Add() (int64, error)
	/*
			   批量添加
			   @parm struct数组
		     @return error 错误
	*/
	AddBatch(obj []interface{}) error
	/*
			   根据自身struct更新
			   @parm
		     @return int64 修改记录的id
		     @return error 错误
	*/
	Update() (int64, error)
	/*
			   批量更新
			   @parm struct数组
		     @return error 错误
	*/
	UpdateBatch(obj []interface{}) error
	/*
			   根据自身struct删除
			   @parm
		     @return int64 影响行数
		     @return error 错误
	*/
	Delete() (int64, error)
	/*
			   批量删除
			   @parm struct struct数组
		     @return error 错误
	*/

	DeleteBatch(obj []interface{}) error
	/*
			   执行sql语句 非查询的语句
			   @parm sql sql语句，valuesql语句中?的部分，可以为空
		     @return int64 影响的行数
		     @return error 错误
	*/
	Exec(sql string, value ...interface{}) (int64, error)
}

/*
获取不同类型数据库连接，支持mysql，mariadb，cockroachDB
*/
func SetConn(db, str string) {
	var err error
	Driver = db
	switch db {
	case "mysql":
		DB, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		DB, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		DB, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}

}
func Checkerr(err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//hooks前置方法
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

func AddBeforeFun(f func(), w string) bool {
	success := true
	switch w {
	case "Select":
		Beforefun.Select = append(Beforefun.Select, f)
	case "Update":
		Beforefun.Update = append(Beforefun.Update, f)
	case "FindByID":
		Beforefun.FindByID = append(Beforefun.FindByID, f)
	case "Add":
		Beforefun.Add = append(Beforefun.Add, f)
	case "AddBatch":
		Beforefun.AddBatch = append(Beforefun.AddBatch, f)
	case "UpdateBatch":
		Beforefun.UpdateBatch = append(Beforefun.UpdateBatch, f)
	case "Delete":
		Beforefun.Delete = append(Beforefun.Delete, f)
	case "DeleteBatch":
		Beforefun.DeleteBatch = append(Beforefun.DeleteBatch, f)
	case "Exec":
		Beforefun.Exec = append(Beforefun.Exec, f)
	}

	return success
}

//hooks后置方法
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

func AddAfterFun(f func(), w string) bool {
	success := true
	switch w {
	case "Select":
		Afterfun.Select = append(Afterfun.Select, f)
	case "Update":
		Afterfun.Update = append(Afterfun.Update, f)
	case "FindByID":
		Afterfun.FindByID = append(Afterfun.FindByID, f)
	case "Add":
		Afterfun.Add = append(Afterfun.Add, f)
	case "AddBatch":
		Afterfun.AddBatch = append(Afterfun.AddBatch, f)
	case "UpdateBatch":
		Afterfun.UpdateBatch = append(Afterfun.UpdateBatch, f)
	case "Delete":
		Afterfun.Delete = append(Afterfun.Delete, f)
	case "DeleteBatch":
		Afterfun.DeleteBatch = append(Afterfun.DeleteBatch, f)
	case "Exec":
		Afterfun.Exec = append(Afterfun.Exec, f)
	}

	return success
}

`
)
