
package model
import (
 "database/sql"
 "log"
 "strconv"
_ "github.com/go-sql-driver/mysql"
)


var sqlHsAuthApi string
var argsHsAuthApi []interface{}
var dbconnHsAuthApi *sql.DB

type HsAuthApi struct {
	Id        int32  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	ApiKey    string `dormCol:"api_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	Name      string `dormCol:"name" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	Type      int8   `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
	CreatedAt int32  `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt int32  `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt int32  `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}

//返回执行语句后sql，调试用
func (hsAuthApi HsAuthApi) GetSql() (string, []interface{}) {
	return sqlHsAuthApi, argsHsAuthApi
}

//设置db
func (hsAuthApi HsAuthApi) SetDBConn(db, str string) {
	var err error
	switch db {
	case "mysql":
		dbconnHsAuthApi, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "mariadb":
		dbconnHsAuthApi, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	case "cockroachDB":
		dbconnHsAuthApi, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("error connecting to the database: ", err)
		}
	}
}

func NewHsAuthApi() HsAuthApi {
	dbconnHsAuthApi = DB
	return HsAuthApi{}
}


func (hsAuthApi HsAuthApi) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	for i := 0; i < len(Beforefun.Select); i++ { //前置hooks
		Beforefun.Select[i]()
	}
	var err error
	if limit > MAXROWS {
		limit = MAXROWS
	}
	ar := make([]interface{}, limit) //0为可变数组长度
	// ar[0].(*HsAuthRecords)
	sqlstr := "select id,api_key,name,type,created_at,updated_at,deleted_at,status_at from hs_auth_api " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsAuthApi = sqlstr
	argsHsAuthApi = value
	rows, err := DB.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
		values[0]=&hsAuthApi.Id
		values[1]=&hsAuthApi.ApiKey
		values[2]=&hsAuthApi.Name
		values[3]=&hsAuthApi.Type
		values[4]=&hsAuthApi.CreatedAt
		values[5]=&hsAuthApi.UpdatedAt
		values[6]=&hsAuthApi.DeletedAt
		values[7]=&hsAuthApi.StatusAt
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		Checkerr(err)
		ar[num] = hsAuthApi
		num++
	}

	for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
		Afterfun.Select[i]()
	}
	return ar, err
}
	

func (hsAuthApi *HsAuthApi) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,api_key,name,type,created_at,updated_at,deleted_at,status_at FROM hs_auth_api WHERE id = ?"
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	rows, err := DB.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsAuthApi, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
		values[0]=&hsAuthApi.Id
		values[1]=&hsAuthApi.ApiKey
		values[2]=&hsAuthApi.Name
		values[3]=&hsAuthApi.Type
		values[4]=&hsAuthApi.CreatedAt
		values[5]=&hsAuthApi.UpdatedAt
		values[6]=&hsAuthApi.DeletedAt
		values[7]=&hsAuthApi.StatusAt
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return hsAuthApi, err
}
	

func (hsAuthApi HsAuthApi) Add() (int64, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	sqlstr := "INSERT INTO hs_auth_api (api_key,name,type,created_at,updated_at,deleted_at,status_at) VALUES (?,?,?,?,?,?,?)"

	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 7)
		args[0]=&hsAuthApi.ApiKey
		args[1]=&hsAuthApi.Name
		args[2]=&hsAuthApi.Type
		args[3]=&hsAuthApi.CreatedAt
		args[4]=&hsAuthApi.UpdatedAt
		args[5]=&hsAuthApi.DeletedAt
		args[6]=&hsAuthApi.StatusAt
		
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
		Afterfun.Add[i]()
	}
	return result.LastInsertId()
}
	

func (hsAuthApi HsAuthApi) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	sqlstr := "INSERT INTO hs_auth_api (api_key,name,type,created_at,updated_at,deleted_at,status_at) VALUES (?,?,?,?,?,?,?)"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 7)

	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args

	for _, value := range obj {
		v := value.(HsAuthApi)
	 		args[0]=v.ApiKey
	 		args[1]=v.Name
	 		args[2]=v.Type
	 		args[3]=v.CreatedAt
	 		args[4]=v.UpdatedAt
	 		args[5]=v.DeletedAt
	 		args[6]=v.StatusAt
	 		
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


func (hsAuthApi *HsAuthApi) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	sqlstr := "UPDATE hs_auth_api SET api_key=?,name=?,type=?,created_at=?,updated_at=?,deleted_at=?,status_at=? where id=?"
	stmtIns, err := DB.Prepare(sqlstr)
	Checkerr(err)
	defer stmtIns.Close()
	args := make([]interface{}, 8)
		args[0]=&hsAuthApi.ApiKey
		args[1]=&hsAuthApi.Name
		args[2]=&hsAuthApi.Type
		args[3]=&hsAuthApi.CreatedAt
		args[4]=&hsAuthApi.UpdatedAt
		args[5]=&hsAuthApi.DeletedAt
		args[6]=&hsAuthApi.StatusAt
		args[7]=&hsAuthApi.Id
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	result, err := stmtIns.Exec(args...)
	Checkerr(err)
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}


func (hsAuthApi HsAuthApi) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	sqlstr := "UPDATE hs_auth_api SET api_key=?,name=?,type=?,created_at=?,updated_at=?,deleted_at=?,status_at=? where id=?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 8)

	for _, value := range obj {
		v := value.(HsAuthApi)
	 		args[0]=v.ApiKey
	 		args[1]=v.Name
	 		args[2]=v.Type
	 		args[3]=v.CreatedAt
	 		args[4]=v.UpdatedAt
	 		args[5]=v.DeletedAt
	 		args[6]=v.StatusAt
	 		args[7]=v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}


func (hsAuthApi HsAuthApi) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
  sqlstr := "DELETE FROM hs_auth_api WHERE id = ?"
	stmt, err := DB.Prepare(sqlstr)
	Checkerr(err)
	args := make([]interface{}, 1)
	args[0] = hsAuthApi.Id
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	defer stmt.Close()
	result, err := stmt.Exec(args...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Delete); i++ { //后置hooks
		Afterfun.Delete[i]()
	}
	return result.RowsAffected()
}


func (hsAuthApi HsAuthApi) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	sqlstr := "DELETE FROM hs_auth_api WHERE id = ?"
	tx, err := DB.Begin()
	Checkerr(err)
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	Checkerr(err)
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(HsAuthApi)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		Checkerr(err)
	}
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	err = tx.Commit()
	Checkerr(err)
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}


func (hsAuthApi HsAuthApi) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := DB.Prepare(sql)
	Checkerr(err)

	sqlHsAuthApi = sql
	argsHsAuthApi = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	Checkerr(err)
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}
