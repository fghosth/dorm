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
	sqlHsAuthPermission          string
	argsHsAuthPermission         []interface{}
	dbconnHsAuthPermission       *sql.DB
	driverHsAuthPermission       string
	addCacheHsAuthPermission     []interface{} //添加缓存数组
	countHsAuthPermission        int           //计数 秒
	addCacheFlagHsAuthPermission = false       //缓存进程是否启动
)

type HsAuthPermission struct {
	Id        int64  `dormCol:"id" dormMysqlType:"int(10)" dorm:"PRIMARY;unsigned;NOT NULL;AUTO_INCREMENT"`
	AppKey    string `dormCol:"app_key" dormMysqlType:"varchar(128)" dorm:"NOT NULL"`
	ApiKey    string `dormCol:"api_key" dormMysqlType:"varchar(256)" dorm:"NOT NULL"`
	CreatedAt string `dormCol:"created_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt string `dormCol:"updated_at" dormMysqlType:"timestamp" dorm:"NOT NULL;DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt string `dormCol:"deleted_at" dormMysqlType:"timestamp" dorm:"DEFAULT NULL"`
	StatusAt  int8   `dormCol:"status_at" dormMysqlType:"tinyint(4)" dorm:"NOT NULL;DEFAULT '1'"`
}

//检查增加缓存
func (hsAuthPermission HsAuthPermission) checkAddCache() {
	for range time.Tick(1 * time.Second) {
		if len(addCacheHsAuthPermission) >= AddCacheLen || countHsAuthPermission >= AddCacheExp {
			err := hsAuthPermission.AddBatch(addCacheHsAuthPermission)
			if err != nil {
				fmt.Println(err)
			}
			countHsAuthPermission = 0
			addCacheHsAuthPermission = make([]interface{}, 0)
		}
		l := new(sync.RWMutex)
		l.Lock()
		countHsAuthPermission++
		l.Unlock()
	}
}

//开始添加缓存进程
func (hsAuthPermission HsAuthPermission) StartAddCache() {
	if UseAddCache {
		addCacheHsAuthPermission = make([]interface{}, 0)
		go hsAuthPermission.checkAddCache()
	}
}

//返回执行语句后sql，调试用
func (hsAuthPermission HsAuthPermission) GetSql() (string, []interface{}) {
	return sqlHsAuthPermission, argsHsAuthPermission
}

//设置db
func (hsAuthPermission HsAuthPermission) SetDBConn(db, str string) {
	var err error
	driverHsAuthPermission = db
	switch db {
	case "mysql":
		dbconnHsAuthPermission, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "mariadb":
		dbconnHsAuthPermission, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "cockroachDB":
		dbconnHsAuthPermission, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "postgresql":
		dbconnHsAuthPermission, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	}
}

func NewHsAuthPermission() HsAuthPermission {
	dbconnHsAuthPermission = DB
	driverHsAuthPermission = Driver
	hsAuthPermission := HsAuthPermission{}

	return hsAuthPermission
}

//获得args字符串(除了update)
func getHsAuthPermissionArgsStr(num int) string {
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
func getHsAuthPermissionArgsStrUpdate() string {
	var argsStr string
	switch driverHsAuthApplication {
	case "mysql":
		argsStr = "app_key=?,api_key=?,created_at=?,updated_at=?,deleted_at=?,status_at=? WHERE " + SDELFLAG + "=0 and id=?"
	case "mariadb":
		argsStr = "app_key=?,api_key=?,created_at=?,updated_at=?,deleted_at=?,status_at=? WHERE " + SDELFLAG + "=0 and id=?"
	case "cockroachDB":
		argsStr = "app_key=$1,api_key=$2,created_at=$3,updated_at=$4,deleted_at=$5,status_at=$6 WHERE " + SDELFLAG + "=0 and id=$7"
	case "postgresql":
		argsStr = "app_key=$1,api_key=$2,created_at=$3,updated_at=$4,deleted_at=$5,status_at=$6 WHERE " + SDELFLAG + "=0 and id=$7"
	}
	return argsStr
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
	sqlstr := "select id,app_key,api_key,created_at,updated_at,deleted_at,status_at from hs_auth_permission where " + SDELFLAG + "=0 " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = value

	//设置缓存
	var ckey string
	if CacheUsed() {
		ckey = UT.Md5Str(sqlHsAuthPermission + UT.JoinInterface(argsHsAuthPermission, "-"))
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

	rows, err := dbconnHsAuthPermission.Query(sqlstr, value...)
	defer rows.Close()
	if err != nil {
		return ar, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthPermission.Id
	values[1] = &hsAuthPermission.AppKey
	values[2] = &hsAuthPermission.ApiKey
	values[3] = &hsAuthPermission.CreatedAt
	values[4] = &hsAuthPermission.UpdatedAt
	values[5] = &hsAuthPermission.DeletedAt
	values[6] = &hsAuthPermission.StatusAt
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		err := rows.Scan(values...)
		if err != nil {
			return ar, err
		}
		ar[num] = hsAuthPermission
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

func (hsAuthPermission *HsAuthPermission) FindByID(id int64) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	argsStr := getHsAuthPermissionArgsStr(1)
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT id,app_key,api_key,created_at,updated_at,deleted_at,status_at FROM hs_auth_permission WHERE " + SDELFLAG + "=0 and  id = " + argsStr
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args

	//设置缓存
	var ckey string
	if CacheUsed() {
		ckey = UT.Md5Str(sqlHsAuthPermission + UT.JoinInterface(argsHsAuthPermission, "-"))
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

	rows, err := dbconnHsAuthPermission.Query(sqlstr, args...)
	defer rows.Close()
	if err != nil {
		return hsAuthPermission, err
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &hsAuthPermission.Id
	values[1] = &hsAuthPermission.AppKey
	values[2] = &hsAuthPermission.ApiKey
	values[3] = &hsAuthPermission.CreatedAt
	values[4] = &hsAuthPermission.UpdatedAt
	values[5] = &hsAuthPermission.DeletedAt
	values[6] = &hsAuthPermission.StatusAt
	if rows.Next() {
		err = rows.Scan(values...)
		Checkerr(err)
	}

	//设置缓存
	if CacheUsed() {
		err = SetCache(ckey, hsAuthPermission)
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
	argsStr := getHsAuthPermissionArgsStr(6)
	sqlstr := "INSERT INTO hs_auth_permission (app_key,api_key,created_at,updated_at,deleted_at,status_at) VALUES (" + argsStr + ")"

	stmtIns, err := dbconnHsAuthPermission.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	args := make([]interface{}, 6)
	args[0] = &hsAuthPermission.AppKey
	args[1] = &hsAuthPermission.ApiKey
	args[2] = &hsAuthPermission.CreatedAt
	args[3] = &hsAuthPermission.UpdatedAt
	args[4] = &hsAuthPermission.DeletedAt
	args[5] = &hsAuthPermission.StatusAt

	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args

	if UseAddCache {
		if !addCacheFlagHsAuthPermission {
			hsAuthPermission.StartAddCache()
			addCacheFlagHsAuthPermission = true
		}
		l := new(sync.RWMutex)
		l.Lock()
		addCacheHsAuthPermission = append(addCacheHsAuthPermission, hsAuthPermission)
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

func (hsAuthPermission HsAuthPermission) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	argsStr := getHsAuthPermissionArgsStr(6)
	sqlstr := "INSERT INTO hs_auth_permission (app_key,api_key,created_at,updated_at,deleted_at,status_at) VALUES (" + argsStr + ")"
	tx, err := dbconnHsAuthPermission.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return err
	}
	args := make([]interface{}, 6)

	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args

	for _, value := range obj {
		v := value.(HsAuthPermission)
		args[0] = v.AppKey
		args[1] = v.ApiKey
		args[2] = v.CreatedAt
		args[3] = v.UpdatedAt
		args[4] = v.DeletedAt
		args[5] = v.StatusAt

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

func (hsAuthPermission *HsAuthPermission) Update() (int64, error) {
	for i := 0; i < len(Beforefun.Update); i++ { //前置hooks
		Beforefun.Update[i]()
	}
	argsStr := getHsAuthPermissionArgsStrUpdate()
	sqlstr := "UPDATE hs_auth_permission SET " + argsStr
	stmtIns, err := dbconnHsAuthPermission.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	args := make([]interface{}, 7)
	args[0] = hsAuthPermission.AppKey
	args[1] = hsAuthPermission.ApiKey
	args[2] = hsAuthPermission.CreatedAt
	args[3] = hsAuthPermission.UpdatedAt
	args[4] = hsAuthPermission.DeletedAt
	args[5] = hsAuthPermission.StatusAt
	args[6] = &hsAuthPermission.Id
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	result, err := stmtIns.Exec(args...)
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(Afterfun.Update); i++ { //后置hooks
		Afterfun.Update[i]()
	}
	return result.RowsAffected()
}

func (hsAuthPermission HsAuthPermission) UpdateBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.UpdateBatch); i++ { //前置hooks
		Beforefun.UpdateBatch[i]()
	}
	argsStr := getHsAuthPermissionArgsStrUpdate()
	sqlstr := "UPDATE hs_auth_permission SET " + argsStr
	tx, err := dbconnHsAuthPermission.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return err
	}
	args := make([]interface{}, 7)

	for _, value := range obj {
		v := value.(HsAuthPermission)
		args[0] = v.AppKey
		args[1] = v.ApiKey
		args[2] = v.CreatedAt
		args[3] = v.UpdatedAt
		args[4] = v.DeletedAt
		args[5] = v.StatusAt
		args[6] = v.Id
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	err = tx.Commit()
	if err != nil {
		return err
	}
	for i := 0; i < len(Afterfun.UpdateBatch); i++ { //后置hooks
		Afterfun.UpdateBatch[i]()
	}

	return err
}

func (hsAuthPermission HsAuthPermission) SDelete() (int64, error) {
	hsAuthPermission.StatusAt = 1
	return hsAuthPermission.Update()
}

func (hsAuthPermission HsAuthPermission) SDeleteBatch(obj []interface{}) error {
	for i := 0; i < len(obj); i++ {
		o := obj[i].(HsAuthPermission)
		o.StatusAt = 1
		obj[i] = o
	}
	return hsAuthPermission.UpdateBatch(obj)
}

func (hsAuthPermission HsAuthPermission) Delete() (int64, error) {
	for i := 0; i < len(Beforefun.Delete); i++ { //前置hooks
		Beforefun.Delete[i]()
	}
	argsStr := getHsAuthPermissionArgsStr(1)
	sqlstr := "DELETE FROM hs_auth_permission WHERE id = " + argsStr
	stmt, err := dbconnHsAuthPermission.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
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
	argsStr := getHsAuthPermissionArgsStr(1)
	sqlstr := "DELETE FROM hs_auth_permission WHERE id = " + argsStr
	tx, err := dbconnHsAuthPermission.Begin()
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
		v := value.(HsAuthPermission)
		args[0] = v.Id
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	sqlHsAuthPermission = sqlstr
	argsHsAuthPermission = args
	err = tx.Commit()
	if err != nil {
		return err
	}
	for i := 0; i < len(Afterfun.DeleteBatch); i++ { //后置hooks
		Afterfun.DeleteBatch[i]()
	}
	return err
}

func (hsAuthPermission HsAuthPermission) Exec(sql string, value ...interface{}) (int64, error) {
	for i := 0; i < len(Beforefun.Exec); i++ { //前置hooks
		Beforefun.Exec[i]()
	}

	stmt, err := dbconnHsAuthPermission.Prepare(sql)
	if err != nil {
		return 0, err
	}

	sqlHsAuthPermission = sql
	argsHsAuthPermission = value
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
