
package model
import (
 "database/sql"
 "log"
 "strconv"
_ "github.com/go-sql-driver/mysql"
)


var sqlHsAuthPermission string
var argsHsAuthPermission []interface{}
var dbconnHsAuthPermission *sql.DB

type HsAuthPermission struct {
	Id        int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	AppKey    string `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	ApiKey    string `dormCol:"api_key" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	CreatedAt int32  `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt int32  `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt int32  `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}

//返回执行语句后sql，调试用
func (hsAuthPermission HsAuthPermission) GetSql() (string, []interface{}) {
	return sqlHsAuthPermission, argsHsAuthPermission
}

//设置db
func (hsAuthPermission HsAuthPermission) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconnHsAuthPermission, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconnHsAuthPermission, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconnHsAuthPermission, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func NewHsAuthPermission() HsAuthPermission {
	dbconnHsAuthPermission = DB
	return HsAuthPermission{}
}


func (hsAuthPermission HsAuthPermission) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select id,app_key,api_key,created_at,updated_at,deleted_at,status_at from hs_auth_permission " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = value
	rows, err := DB.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
		values[0]=&hsAuthPermission.Id
		values[1]=&hsAuthPermission.AppKey
		values[2]=&hsAuthPermission.ApiKey
		values[3]=&hsAuthPermission.CreatedAt
		values[4]=&hsAuthPermission.UpdatedAt
		values[5]=&hsAuthPermission.DeletedAt
		values[6]=&hsAuthPermission.StatusAt
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = hsAuthPermission
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}
	

func (hsAuthPermission *HsAuthPermission) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,app_key,api_key,created_at,updated_at,deleted_at,status_at FROM hs_auth_permission WHERE id = ?"
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	rows, err := DB.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsAuthPermission, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
		values[0]=&hsAuthPermission.Id
		values[1]=&hsAuthPermission.AppKey
		values[2]=&hsAuthPermission.ApiKey
		values[3]=&hsAuthPermission.CreatedAt
		values[4]=&hsAuthPermission.UpdatedAt
		values[5]=&hsAuthPermission.DeletedAt
		values[6]=&hsAuthPermission.StatusAt
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return hsAuthPermission, err
}
	

func (hsAuthPermission HsAuthPermission) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO hs_auth_permission (app_key,api_key,created_at,updated_at,deleted_at,status_at) VALUES (?,?,?,?,?,?)"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 6)
		args[0]=&hsAuthPermission.AppKey
		args[1]=&hsAuthPermission.ApiKey
		args[2]=&hsAuthPermission.CreatedAt
		args[3]=&hsAuthPermission.UpdatedAt
		args[4]=&hsAuthPermission.DeletedAt
		args[5]=&hsAuthPermission.StatusAt
		
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}
	

func (hsAuthPermission HsAuthPermission) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	sqlstr := "INSERT INTO hs_auth_permission (app_key,api_key,created_at,updated_at,deleted_at,status_at) VALUES (?,?,?,?,?,?)"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 6)

	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args

	for _, value := range obj {
		v := value.(HsAuthPermission)
	 		args[0]=v.AppKey
	 		args[1]=v.ApiKey
	 		args[2]=v.CreatedAt
	 		args[3]=v.UpdatedAt
	 		args[4]=v.DeletedAt
	 		args[5]=v.StatusAt
	 		
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


func (hsAuthPermission *HsAuthPermission) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE hs_auth_permission SET app_key=?,api_key=?,created_at=?,updated_at=?,deleted_at=?,status_at=? where id=?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 7)
		args[0]=&hsAuthPermission.AppKey
		args[1]=&hsAuthPermission.ApiKey
		args[2]=&hsAuthPermission.CreatedAt
		args[3]=&hsAuthPermission.UpdatedAt
		args[4]=&hsAuthPermission.DeletedAt
		args[5]=&hsAuthPermission.StatusAt
		args[6]=&hsAuthPermission.Id
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}


func (hsAuthPermission HsAuthPermission) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE hs_auth_permission SET app_key=?,api_key=?,created_at=?,updated_at=?,deleted_at=?,status_at=? where id=?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 7)

	for _, value := range obj {
		v := value.(HsAuthPermission)
	 		args[0]=v.AppKey
	 		args[1]=v.ApiKey
	 		args[2]=v.CreatedAt
	 		args[3]=v.UpdatedAt
	 		args[4]=v.DeletedAt
	 		args[5]=v.StatusAt
	 		args[6]=v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}


func (hsAuthPermission HsAuthPermission) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
  sqlstr := "DELETE FROM hs_auth_permission WHERE id = ?"
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = hsAuthPermission.Id
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}


func (hsAuthPermission HsAuthPermission) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM hs_auth_permission WHERE id = ?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(HsAuthPermission)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}


func (hsAuthPermission HsAuthPermission) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlHsAuthPermission = sql
	argsHsAuthPermission = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}

