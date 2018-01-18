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

	sqlPrint = sqlstr
	argsPrint = value
	rows, err := DB.Query(sqlstr, value...)
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
func ({{{objvar}}} {{{obj}}}) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT {{{fields}}} FROM {{{tableName}}} WHERE {{{sqlField}}} = ?"
	sqlPrint = sqlstr
	argsPrint = args
	rows, err := DB.Query(sqlstr, args...)
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
	sqlstr := "INSERT INTO {{{tableName}}} ({{{fields}}}) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 13)
	{{#each field}}
		{{{this}}}
	{{/each}}
	sqlPrint = sqlstr
	argsPrint = args
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
	sqlstr := "INSERT INTO {{{tableName}}} ({{{fields}}}) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 13)

	sqlPrint = sqlstr
	argsPrint = args

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
func ({{{objvar}}} {{{obj}}}) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE {{{tableName}}} SET {{{fields}}} where {{{sqlField}}}=?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 14)
	{{#each field}}
		{{{this}}}
	{{/each}}
	sqlPrint = sqlstr
	argsPrint = args
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
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 14)

	for _, value := range obj {
		v := value.({{{obj}}})
		{{#each field}}
	 		{{{this}}}
		{{/each}}
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlPrint = sqlstr
	argsPrint = args
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
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = {{{objvar}}}.{{{structField}}}
	sqlPrint = sqlstr
	argsPrint = args
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
	tx, err := DB.Begin()
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
	sqlPrint = sqlstr
	argsPrint = args
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

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlPrint = sql
	argsPrint = value
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
var sqlPrint string
var argsPrint []interface{}
var dbconn *sql.DB
`
	Header_TPL = `
package {{{pkname}}}
import (
	{{#each field}}
 {{{this}}}
	{{/each}}
	_ "{{{sqlDriver}}}"
)
`
	Function_TPL = `
//返回执行语句后sql，调试用
func ({{{objvar}}} {{{obj}}}) GetSql() (string, []interface{}) {
	return sqlPrint, argsPrint
}

//设置db
func ({{{objvar}}} {{{obj}}}) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconn, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconn, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconn, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func New{{{obj}}}() HsAuthRecords {
	dbconn = DB
	return HsAuthRecords{}
}
`
)
