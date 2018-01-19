
package model
import (
 "database/sql"
 "log"
 "strconv"
_ "github.com/go-sql-driver/mysql"
)


var sqlModel string
var argsModel []interface{}
var dbconnModel *sql.DB

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
	Select(sql string, limit, offset int, value ...interface{}

//返回执行语句后sql，调试用
func (model Model) GetSql() (string, []interface{}) {
	return sqlModel, argsModel
}

//设置db
func (model Model) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconnModel, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconnModel, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconnModel, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func NewModel() Model {
	dbconnModel = DB
	return Model{}
}


func (model Model) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select select from model " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlModel = sqlstr
	argsModel = value
	rows, err := DB.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
		values[0]=&model.Select(sql
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = model
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}
	

func (model *Model) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT select FROM model WHERE select = ?"
	sqlModel = sqlstr
	argsModel = args
	rows, err := DB.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return model, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
		values[0]=&model.Select(sql
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return model, err
}
	

func (model Model) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO model () VALUES ()"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 0)
		
	sqlModel = sqlstr
	argsModel = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}
	

func (model Model) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	sqlstr := "INSERT INTO model () VALUES ()"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 0)

	sqlModel = sqlstr
	argsModel = args

	for _, value := range obj {
		v := value.(Model)
	 		
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


func (model *Model) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE model SET  where select=?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 1)
		args[0]=&model.Select(sql
	sqlModel = sqlstr
	argsModel = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}


func (model Model) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE model SET  where select=?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(Model)
	 		args[0]=v.Select(sql
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlModel = sqlstr
	argsModel = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}


func (model Model) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
  sqlstr := "DELETE FROM model WHERE select = ?"
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = model.Select(sql
	sqlModel = sqlstr
	argsModel = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}


func (model Model) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM model WHERE select = ?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(Model)
		args[0] = v.Select(sql
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlModel = sqlstr
	argsModel = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}


func (model Model) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlModel = sql
	argsModel = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}

