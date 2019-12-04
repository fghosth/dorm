package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	sqlHsAuthApi          string
	argsHsAuthApi         []interface{}
	dbconnHsAuthApi       *sql.DB
	driverHsAuthApi       string
	addCacheHsAuthApi     []interface{} //添加缓存数组
	countHsAuthApi        int           //计数 秒
	addCacheFlagHsAuthApi = false       //缓存进程是否启动
)

type HsAuthApi struct {
	Id        int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	ApiKey    string `dormCol:"api_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	Name      string `dormCol:"name" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	Type      int8   `dormCol:"type" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '0'"`
	CreatedAt string `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt string `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt string `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}

//检查增加缓存
func (hsAuthApi HsAuthApi) checkAddCache() {
	for range time.Tick(1 * time.Second) {
		if len(addCacheHsAuthApi) >= AddCacheLen || countHsAuthApi >= AddCacheExp {
			err := hsAuthApi.AddBatch(addCacheHsAuthApi)
			if err != nil {
				fmt.Println(err)
			}
			countHsAuthApi = 0
			addCacheHsAuthApi = make([]interface{}, 0)
		}
		l := new(sync.RWMutex)
		l.Lock()
		countHsAuthApi++
		l.Unlock()
	}
}

//开始添加缓存进程
func (hsAuthApi HsAuthApi) StartAddCache() {
	if UseAddCache {
		addCacheHsAuthApi = make([]interface{}, 0)
		go hsAuthApi.checkAddCache()
	}
}

//返回执行语句后sql，调试用
func (hsAuthApi HsAuthApi) GetSql() (string, []interface{}) {
	return sqlHsAuthApi, argsHsAuthApi
}

//设置db
func (hsAuthApi HsAuthApi) SetDBConn(db, str string) {
	var err error
	driverHsAuthApi = db
	switch db {
	case "mysql":
		dbconnHsAuthApi, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "mariadb":
		dbconnHsAuthApi, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "cockroachDB":
		dbconnHsAuthApi, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "postgresql":
		dbconnHsAuthApi, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	}
}

func NewHsAuthApi() HsAuthApi {
	dbconnHsAuthApi = DB
	driverHsAuthApi = Driver
	hsAuthApi := HsAuthApi{}

	return hsAuthApi
}

//获得args字符串(除了update)
func getHsAuthApiArgsStr(num int) string {
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
func getHsAuthApiArgsStrUpdate() string {
	var argsStr string
	switch driverHsAuthApplication {
	case "mysql":
		argsStr = "api_key=?,name=?,type=?,created_at=?,updated_at=?,deleted_at=?,status_at=? WHERE " + SDELFLAG + "=0 and id=?"
	case "mariadb":
		argsStr = "api_key=?,name=?,type=?,created_at=?,updated_at=?,deleted_at=?,status_at=? WHERE " + SDELFLAG + "=0 and id=?"
	case "cockroachDB":
		argsStr = "api_key=$1,name=$2,type=$3,created_at=$4,updated_at=$5,deleted_at=$6,status_at=$7 WHERE " + SDELFLAG + "=0 and id=$8"
	case "postgresql":
		argsStr = "api_key=$1,name=$2,type=$3,created_at=$4,updated_at=$5,deleted_at=$6,status_at=$7 WHERE " + SDELFLAG + "=0 and id=$8"
	}
	return argsStr
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
	sqlstr := "select id,api_key,name,type,created_at,updated_at,deleted_at,status_at from hs_auth_api where " + SDELFLAG + "=0 " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsAuthApi = sqlstr
	argsHsAuthApi = value

	//设置缓存
	var ckey string
	if CacheUsed() {
		ckey = UT.Md5Str(sqlHsAuthApi + UT.JoinInterface(argsHsAuthApi, "-"))
		cv, err := GetCache(ckey)
		Checkerr(err)
		if err == nil { //命中缓存
			for i := 0; i < len(Afterfun.Select); i++ { //后置hooks
				Afterfun.Select[i]()
			}
			res, ok := cv.([]interface{})
			if ok {
				return res, nil
			}
		}
	}

	rows, err := dbconnHsAuthApi.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthApi.Id
	values[1] = &hsAuthApi.ApiKey
	values[2] = &hsAuthApi.Name
	values[3] = &hsAuthApi.Type
	values[4] = &hsAuthApi.CreatedAt
	values[5] = &hsAuthApi.UpdatedAt
	values[6] = &hsAuthApi.DeletedAt
	values[7] = &hsAuthApi.StatusAt
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		if err != nil {
			return ar, err
		}
		ar[num] = hsAuthApi
		num++
	}
	//设置缓存
	if CacheUsed() {
		err = SetCache(ckey, ar)
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
	argsStr := getHsAuthApiArgsStr(1)
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,api_key,name,type,created_at,updated_at,deleted_at,status_at FROM hs_auth_api WHERE " + SDELFLAG + "=0 and  id = " + argsStr
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args

	//设置缓存
	var ckey string
	if CacheUsed() {
		ckey = UT.Md5Str(sqlHsAuthApi + UT.JoinInterface(argsHsAuthApi, "-"))
		cv, err := GetCache(ckey)
		if err == nil { //命中缓存
			for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
				Afterfun.FindByID[i]()
			}
			res, ok := cv.([]interface{})
			if ok {
				return res, nil
			}
		}
	}

	rows, err := dbconnHsAuthApi.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsAuthApi, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthApi.Id
	values[1] = &hsAuthApi.ApiKey
	values[2] = &hsAuthApi.Name
	values[3] = &hsAuthApi.Type
	values[4] = &hsAuthApi.CreatedAt
	values[5] = &hsAuthApi.UpdatedAt
	values[6] = &hsAuthApi.DeletedAt
	values[7] = &hsAuthApi.StatusAt
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}

	//设置缓存
	if CacheUsed() {
		err = SetCache(ckey, hsAuthApi)
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
	argsStr := getHsAuthApiArgsStr(7)
	sqlstr := "INSERT INTO hs_auth_api (api_key,name,type,created_at,updated_at,deleted_at,status_at) VALUES (" + argsStr + ")"

	stmtIns, err := dbconnHsAuthApi.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	args := make([]interface{}, 7)
	args[0] = &hsAuthApi.ApiKey
	args[1] = &hsAuthApi.Name
	args[2] = &hsAuthApi.Type
	args[3] = &hsAuthApi.CreatedAt
	args[4] = &hsAuthApi.UpdatedAt
	args[5] = &hsAuthApi.DeletedAt
	args[6] = &hsAuthApi.StatusAt

	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args

	if UseAddCache {
		if !addCacheFlagHsAuthApi {
			hsAuthApi.StartAddCache()
			addCacheFlagHsAuthApi = true
		}
		l := new(sync.RWMutex)
		l.Lock()
		addCacheHsAuthApi = append(addCacheHsAuthApi, hsAuthApi)
		defer l.Unlock()
		return 0, err
	} else {
		result, err := stmtIns.Exec(args...)
		if err != nil {
			return 0, err
		}
		for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
			Afterfun.Add[i]()
		}
		_, e := result.LastInsertId()
		if err == nil && e != nil {
			return 0, nil
		}
		return result.LastInsertId()
	}

}

func (hsAuthApi HsAuthApi) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	argsStr := getHsAuthApiArgsStr(7)
	sqlstr := "INSERT INTO hs_auth_api (api_key,name,type,created_at,updated_at,deleted_at,status_at) VALUES (" + argsStr + ")"
	tx, err := dbconnHsAuthApi.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return err
	}
	args := make([]interface{}, 7)

	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args

	for _, value := range obj {
		v := value.(HsAuthApi)
		args[0] = v.ApiKey
		args[1] = v.Name
		args[2] = v.Type
		args[3] = v.CreatedAt
		args[4] = v.UpdatedAt
		args[5] = v.DeletedAt
		args[6] = v.StatusAt

		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	for i := 0; i < len(Afterfun.AddBatch); i++ { //后置hooks
		Afterfun.AddBatch[i]()
	}

	return err
}

func (hsAuthApi *HsAuthApi) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	argsStr := getHsAuthApiArgsStrUpdate()
	sqlstr := "UPDATE hs_auth_api SET " + argsStr
	stmtIns, err := dbconnHsAuthApi.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	args := make([]interface{}, 8)
	args[0] = hsAuthApi.ApiKey
	args[1] = hsAuthApi.Name
	args[2] = hsAuthApi.Type
	args[3] = hsAuthApi.CreatedAt
	args[4] = hsAuthApi.UpdatedAt
	args[5] = hsAuthApi.DeletedAt
	args[6] = hsAuthApi.StatusAt
	args[7] = &hsAuthApi.Id
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	result, err := stmtIns.Exec(args...)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}

func (hsAuthApi HsAuthApi) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	argsStr := getHsAuthApiArgsStrUpdate()
	sqlstr := "UPDATE hs_auth_api SET " + argsStr
	tx, err := dbconnHsAuthApi.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return err
	}
	args := make([]interface{}, 8)

	for _, value := range obj {
		v := value.(HsAuthApi)
		args[0] = v.ApiKey
		args[1] = v.Name
		args[2] = v.Type
		args[3] = v.CreatedAt
		args[4] = v.UpdatedAt
		args[5] = v.DeletedAt
		args[6] = v.StatusAt
		args[7] = v.Id
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	err = tx.Commit()
	if err != nil {
		return err
	}
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}

func (hsAuthApi HsAuthApi) SDelete() (int64, error) {
	hsAuthApi.StatusAt = 1
	return hsAuthApi.Update()
}

func (hsAuthApi HsAuthApi) SDeleteBatch(obj []interface{}) error {
	for i := 0; i < len(obj); i++ {
		o := obj[i].(HsAuthApi)
		o.StatusAt = 1
		obj[i] = o
	}
	return hsAuthApi.UpdateBatch(obj)
}

func (hsAuthApi HsAuthApi) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
	argsStr := getHsAuthApiArgsStr(1)
	sqlstr := "DELETE FROM hs_auth_api WHERE id = " + argsStr
	stmt, err := dbconnHsAuthApi.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
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
	argsStr := getHsAuthApiArgsStr(1)
	sqlstr := "DELETE FROM hs_auth_api WHERE id = " + argsStr
	tx, err := dbconnHsAuthApi.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return err
	}
	args := make([]interface{}, 1)

	for _, value := range obj {
		v := value.(HsAuthApi)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	sqlHsAuthApi = sqlstr
	argsHsAuthApi = args
	err = tx.Commit()
	if err != nil {
		return err
	}
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}

func (hsAuthApi HsAuthApi) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := dbconnHsAuthApi.Prepare(sql)
	if err != nil {
		return 0, err
	}

	sqlHsAuthApi = sql
	argsHsAuthApi = value
	defer stmt.Close()
	result, err := stmt.Exec(value...)

	if err != nil {
		return 0, err
	}
	for i := 0; i < len(Afterfun.Exec); i++ { //后置hooks
		Afterfun.Exec[i]()
	}
	return result.RowsAffected()
}
