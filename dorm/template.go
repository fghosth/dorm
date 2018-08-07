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
	sqlstr := "select {{{fields}}} from {{{tableName}}} where "+ SDELFLAG + "=0 " + sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	sql{{{obj}}} = sqlstr
	args{{{obj}}} = value

	//设置缓存
	var ckey string
	if CacheUsed() {
		ckey = UT.Md5Str(sql{{{obj}}} + UT.JoinInterface(args{{{obj}}}, "-"))
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
		if err != nil {
			return ar, err
		}
		ar[num] = {{{objvar}}}
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
	`
	FindByID_TPL = `
func ({{{objvar}}} *{{{obj}}}) FindByID(id {{{keytype}}}) (interface{}, error) {
	for i := 0; i < len(Beforefun.FindByID); i++ { //前置hooks
		Beforefun.FindByID[i]()
	}
	argsStr := get{{{obj}}}ArgsStr(1)
	args := make([]interface{}, 1)
	args[0] = id
	sqlstr := "SELECT {{{fields}}} FROM {{{tableName}}} WHERE " + SDELFLAG + "=0 and  {{{sqlField}}} = " + argsStr
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args

	//设置缓存
	var ckey string
	if CacheUsed() {
		ckey = UT.Md5Str(sql{{{obj}}} + UT.JoinInterface(args{{{obj}}}, "-"))
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

	//设置缓存
	if CacheUsed() {
		err = SetCache(ckey, {{{objvar}}})
	}
	for i := 0; i < len(Afterfun.FindByID); i++ { //后置hooks
		Afterfun.FindByID[i]()
	}
	return {{{objvar}}}, err
}
	`

	Add_TPL = `
func ({{{objvar}}} {{{obj}}}) Add() ({{{keytype}}}, error) {
	for i := 0; i < len(Beforefun.Add); i++ { //前置hooks
		Beforefun.Add[i]()
	}
	argsStr := get{{{obj}}}ArgsStr({{{len}}})
	sqlstr := "INSERT INTO {{{tableName}}} ({{{fields}}}) VALUES (" + argsStr + ")"

	stmtIns, err := dbconn{{{obj}}}.Prepare(sqlstr)
	if err != nil {
		{{{retstr}}}
	}
	defer stmtIns.Close()
	args := make([]interface{}, {{{len}}})
	{{#each field}}
		{{{this}}}
	{{/each}}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args


	if UseAddCache {
		if !addCacheFlag{{{obj}}}  {
			{{{objvar}}}.StartAddCache()
			addCacheFlag{{{obj}}}  = true
		}
		l := new(sync.RWMutex)
		l.Lock()
		addCache{{{obj}}} = append(addCache{{{obj}}}, {{{objvar}}})
		defer l.Unlock()
		{{{retstr}}}
	} else {
		result, err := stmtIns.Exec(args...)
		if err != nil {
			{{{retstr}}}
		}
		for i := 0; i < len(Afterfun.Add); i++ { //后置hooks
			Afterfun.Add[i]()
		}
		_, e := result.LastInsertId()
		if err == nil && e != nil {
			{{{retstr}}}
		}
		{{{retIDstr}}}
	}

}
	`
	AddBatch_TPL = `
func ({{{objvar}}} {{{obj}}}) AddBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.AddBatch); i++ { //前置hooks
		Beforefun.AddBatch[i]()
	}
	argsStr := get{{{obj}}}ArgsStr({{{len}}})
	sqlstr := "INSERT INTO {{{tableName}}} ({{{fields}}}) VALUES (" + argsStr + ")"
	tx, err := dbconn{{{obj}}}.Begin()
	if err != nil {
		return  err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return  err
	}
	args := make([]interface{}, {{{len}}})

	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args

	for _, value := range obj {
		v := value.({{{obj}}})
		{{#each field}}
	 		{{{this}}}
		{{/each}}
		_, err = stmt.Exec(args...)
		if err != nil {
			return  err
		}
	}
	err = tx.Commit()
	if err != nil {
		return  err
	}
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
	argsStr := get{{{obj}}}ArgsStrUpdate()
	sqlstr := "UPDATE {{{tableName}}} SET " + argsStr
	stmtIns, err := dbconn{{{obj}}}.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()
	args := make([]interface{}, {{{len}}})
	{{#each field}}
		{{{this}}}
	{{/each}}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	result, err := stmtIns.Exec(args...)
	if err != nil {
		return 0, err
	}
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
	argsStr := get{{{obj}}}ArgsStrUpdate()
	sqlstr := "UPDATE {{{tableName}}} SET " + argsStr
	tx, err := dbconn{{{obj}}}.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil {
		return err
	}
	args := make([]interface{}, {{{len}}})

	for _, value := range obj {
		v := value.({{{obj}}})
		{{#each field}}
	 		{{{this}}}
		{{/each}}
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	err = tx.Commit()
	if err != nil {
		return err
	}
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
	argsStr := get{{{obj}}}ArgsStr(1)
  sqlstr := "DELETE FROM {{{tableName}}} WHERE {{{sqlField}}} = " + argsStr
	stmt, err := dbconn{{{obj}}}.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
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
	if result == nil {
		return 0, nil
	} else {
		return result.RowsAffected()
	}
}
`
	DeleteBatch_TPL = `
func ({{{objvar}}} {{{obj}}}) DeleteBatch(obj []interface{}) error {
	for i := 0; i < len(Beforefun.DeleteBatch); i++ { //前置hooks
		Beforefun.DeleteBatch[i]()
	}
	argsStr := get{{{obj}}}ArgsStr(1)
	sqlstr := "DELETE FROM {{{tableName}}} WHERE {{{sqlField}}} = " + argsStr
	tx, err := dbconn{{{obj}}}.Begin()
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
		v := value.({{{obj}}})
		args[0] = v.{{{structField}}}
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	sql{{{obj}}} = sqlstr
	args{{{obj}}} = args
	err = tx.Commit()
	if err != nil {
		return err
	}
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
	if err != nil {
		return 0, err
	}

	sql{{{obj}}} = sql
	args{{{obj}}} = value
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
`
	GetArgsStrFun_TPL = `
	//获得args字符串(除了update)
	func get{{{obj}}}ArgsStr(num int) string {
		var argsStr string
		switch driver{{{obj}}} {
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
	func get{{{obj}}}ArgsStrUpdate() string {
		var argsStr string
		switch driver{{{obj}}} {
		case "mysql":
			argsStr = "{{{mysqlField}}}"
		case "mariadb":
			argsStr = "{{{mariadbField}}}"
		case "cockroachDB":
			argsStr = "{{{cockroachDBField}}}"
		case "postgresql":
			argsStr = "{{{postgresqlField}}}"
		}
		return argsStr
	}

`
	SDEL_TPL = `
func ({{{objvar}}} {{{obj}}}) SDelete() (int64, error) {
	{{{objvar}}}.StatusAt = 1
	return {{{objvar}}}.Update()
}

func ({{{objvar}}} {{{obj}}}) SDeleteBatch(obj []interface{}) error {
	for i := 0; i < len(obj); i++ {
		o := obj[i].({{{obj}}})
		o.StatusAt = 1
		obj[i] = o
	}
	return {{{objvar}}}.UpdateBatch(obj)
}
`
	Field_TPL = `
var (
	sql{{{obj}}} string
	args{{{obj}}} []interface{}
	dbconn{{{obj}}} *sql.DB
	driver{{{obj}}} string
	addCache{{{obj}}} []interface{} //添加缓存数组
	count{{{obj}}} int              //计数 秒
	addCacheFlag{{{obj}}}      = false //缓存进程是否启动
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
	//检查增加缓存
	func ({{{objvar}}} {{{obj}}}) checkAddCache() {
		for range time.Tick(1 * time.Second) {
			if len(addCache{{{obj}}}) >= AddCacheLen || count{{{obj}}} >= AddCacheExp {
				err := {{{objvar}}}.AddBatch(addCache{{{obj}}})
				if err != nil {
					fmt.Println(err)
				}
				count{{{obj}}} = 0
				addCache{{{obj}}} = make([]interface{}, 0)
			}
			l := new(sync.RWMutex)
			l.Lock()
			count{{{obj}}}++
			l.Unlock()
		}
	}



	//开始添加缓存进程
	func ({{{objvar}}} {{{obj}}}) StartAddCache()  {
		if UseAddCache {
			addCache{{{obj}}} = make([]interface{}, 0)
			go {{{objvar}}}.checkAddCache()
		}
	}

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
			log.Fatal("数据库连接错误: ", err)
		}
	case "mariadb":
		dbconn{{{obj}}}, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "cockroachDB":
		dbconn{{{obj}}}, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "postgresql":
		dbconn{{{obj}}}, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	}
}

func New{{{obj}}}() {{{obj}}} {
	dbconn{{{obj}}} = DB
	driver{{{obj}}} = Driver
	{{{objvar}}} := {{{obj}}}{}

	return {{{objvar}}}
}
`
	DAO_TPL = `
package {{{pkname}}}

import (
 "{{{modelImport}}}"
)

type {{{obj}}}Dao struct {
	model base.Model
	base.{{{obj}}}
}
/*
			 根据条件查找结果集
			 @parm sql 除去select where 1=1  xxx,xxx from tablename 之后的东西 如果要加where先加『and』eg【and username = "derek"】
			 @parm value sql中?值 可以为空
			 @parm limit 显示数量
			 @parm offset 数据位置0开始
			 @return struct 集合
			 @return error 错误
*/
func (dao {{{obj}}}Dao) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	return dao.model.Select(sql, limit, offset, value...)
}
/*
			 根据主键查找
			 @parm id 主键
			 @return struct
			 @return error 错误
*/
func (dao *{{{obj}}}Dao) FindByID(id {{{keytype}}}) (interface{}, error) {
	res, err := dao.model.FindByID(id)
	re := res.(*base.{{{obj}}})
	dao.{{{obj}}} = *re
	return res, err
}
/*
			 根据自身struct内容添加
			 @parm
			 @return 返回主键id
			 @return error 错误
*/
func (dao {{{obj}}}Dao) Add() ({{{keytype}}}, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.Add()
}
/*
			 批量添加
			 @parm struct数组
			 @return error 错误
*/
func (dao {{{obj}}}Dao) AddBatch(obj []interface{}) error {
	return dao.model.AddBatch(obj)
}
/*
			 根据自身struct更新
			 @parm
			 @return  修改记录的id
			 @return error 错误
*/
func (dao {{{obj}}}Dao) Update() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.Update()
}
/*
			 批量更新
			 @parm struct数组
			 @return error 错误
*/
func (dao {{{obj}}}Dao) UpdateBatch(obj []interface{}) error {
	return dao.model.UpdateBatch(obj)
}
/*
			 根据自身struct删除
			 @parm
			 @return  影响行数
			 @return error 错误
*/
func (dao {{{obj}}}Dao) Delete() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.Delete()
}
/*
			 批量删除
			 @parm struct struct数组
			 @return error 错误
*/
func (dao {{{obj}}}Dao) DeleteBatch(obj []interface{}) error {
	return dao.model.DeleteBatch(obj)
}
/*
 根据自身struct软删除
 @parm
 @return  影响数据id
 @return error 错误
*/
func (dao {{{obj}}}Dao) SDelete() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.SDelete()
}

/*
 批量软删除
 @parm struct struct数组
 @return error 错误
*/

func (dao {{{obj}}}Dao) SDeleteBatch(obj []interface{}) error {
	return dao.model.SDeleteBatch(obj)
}

/*
			 执行sql语句 非查询的语句
			 @parm sql sql语句，valuesql语句中?的部分，可以为空
			 @return int64 影响的行数
			 @return error 错误
*/
func (dao {{{obj}}}Dao) Exec(sql string, value ...interface{}) (int64, error) {
	return dao.model.Exec(sql, value...)
}
/*
			 获取最后执行的sql语句 和参数
			 @return string sql语句和参数
*/
func (dao {{{obj}}}Dao) GetSql() (string, []interface{}) {
	return dao.model.GetSql()
}
/*
			 设置当前对象的链接
			 @db 数据库默认值mysql 支持mysql，mariadb，cockroachDB
			 @str 数据库连接 『postgresql://derek:123456@localhost:26257/auth?sslmode=disable』 【root:@tcp(localhost:3306)/praise_auth?charset=utf8】
*/
func (dao {{{obj}}}Dao) SetDBConn(db, str string) {
	dao.model.SetDBConn(db, str)
}

//获取有值的对象
func (daoo {{{obj}}}Dao) getObjWithValue(dao {{{obj}}}Dao) base.{{{obj}}} {
	{{{objvar}}} := base.New{{{obj}}}()
	{{#each field}}
	{{{this}}}
	{{/each}}
	return {{{objvar}}}
}

func New{{{obj}}}Dao() {{{obj}}}Dao {
	ap := base.New{{{obj}}}()
	aa := base.{{{obj}}}{}
	return {{{obj}}}Dao{&ap, aa}
}
`
	MODEL_TPL = `
package {{{pkname}}}

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/fghosth/algo-go/gcache"
)

const (
	LIMIT   = 500  //默认查询条数限制
	OFFSET  = 0    //默认位置
	MAXROWS = 1000 //最多查出多少条,-1为不限制
	SDELFLAG = "status_at" //数据库必出有这个字段才有用，软删除字段 0为未删除，1为删除
	UNDEL    = 0           //为删除
	DELED    = 1           //删除
)


var (
	DB        *sql.DB //数据库连接
	Beforefun Before
	Afterfun  After
	Driver    string
	cacheLen  = 3000  //缓存条数
	cacheTime = 30   //缓存时间 秒
	useCache  = true  //是否使用缓存
	cacheType = "ARC" //缓存类型:LRU,LFU,ARC
	cache     gcache.Cache
	UseAddCache = false           //是否使用插入缓存 如每2秒写一次数据库，或者超过300条写一次数据库
	AddCacheLen = 300            //插入缓存数量
	AddCacheExp = 3              //插入缓存过期时间 秒
	UT        = Dstring{} //工具类
)

//设置缓存类型
func SetCacheType(ctype string, clen int) {
	switch ctype {
	case "LRU":
		cache = gcache.New(clen).LRU().Build()
	case "LFU":
		cache = gcache.New(clen).LFU().Build()
	case "ARC":
		cache = gcache.New(clen).ARC().Build()
	}
}

//设置缓存时间 秒 默认30
func SetCacheTime(t int) {
	cacheTime = t
}

//获取缓存时间
func GetCacheTime() int {
	return cacheTime
}

//设置缓存容量 个数默认2000
func SetCacheLen(l int) {
	cacheLen = l
}

//获取缓存容量
func GetCacheLen() int {
	return cacheLen
}

//获取已缓存的数量
func GetCacheUsedLen() int {
	return cache.Len()
}


//设置是否开启缓存true为开启，false关闭
func UseCache(uc bool) {
	useCache = uc
}

//获取缓存是否开启
func CacheUsed() bool {
	return useCache
}

//获取缓存命中率
func GetCacheRate() float64 {
	return cache.HitRate()
}

//设置缓存
func SetCache(k, v interface{}) error {
	return cache.SetWithExpire(k, v, time.Second*time.Duration(cacheTime))
}

//获取缓存
func GetCache(k interface{}) (interface{}, error) {
	return cache.Get(k)
}

func init() {
	SetCacheType(cacheType, cacheLen)
	// SetConn("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
	SetConn("cockroachDB", "postgresql://root@alcockroach1:26257/uuabc?sslmode=disable")
}

/*
模型的基本方法接口
*/
type Model interface {
	/*
			   根据条件查找结果集
			   @parm sql 除去select where 1=1  xxx,xxx from tablename 之后的东西 如果要加where先加『and』eg【and username = "derek"】
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
	FindByID(id {{{keytype}}}) (interface{}, error)
	/*
			   根据自身struct内容添加
			   @parm
		     @return 返回主键id
		     @return error 错误
	*/
	Add() ({{{keytype}}}, error)
	/*
			   批量添加
			   @parm struct数组
		     @return error 错误
	*/
	AddBatch(obj []interface{}) error
	/*
			   根据自身struct更新
			   @parm
		     @return  修改记录的id
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
		     @return  影响行数
		     @return error 错误
	*/
	Delete() (int64, error)
	/*
			   根据自身struct软删除
			   @parm
		     @return  影响行数
		     @return error 错误
	*/
	SDelete() (int64, error)
	/*
			   批量软删除
			   @parm struct struct数组
		     @return error 错误
	*/

	SDeleteBatch(obj []interface{}) error
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
	/*
			   获取最后执行的sql语句 和参数
		     @return string sql语句和参数
	*/
	GetSql() (string, []interface{})
	/*
			   设置当前对象的链接
			   @db 数据库默认值mysql 支持mysql，mariadb，cockroachDB
			   @str 数据库连接 『postgresql://derek:123456@localhost:26257/auth?sslmode=disable』 【root:@tcp(localhost:3306)/praise_auth?charset=utf8】
	*/
	SetDBConn(db, str string)
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
			log.Fatal("数据库连接错误: ", err)
		}
	case "mariadb":
		DB, err = sql.Open("mysql", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "cockroachDB":
		DB, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
		}
	case "postgresql":
		DB, err = sql.Open("postgres", str)
		if err != nil {
			log.Fatal("数据库连接错误: ", err)
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
	UTIL_TMP = `
package base

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type Dstring struct {
}

//对象数组转化为义某分隔符合并的字符串
func (ds *Dstring) JoinInterface(obj []interface{}, seq string) string {
	var str, tmp string
	for i := 0; i < len(obj); i++ {
		b, ok := obj[i].(int)
		if ok {
			tmp = strconv.Itoa(b)
		}
		c, ok := obj[i].(string)
		if ok {
			tmp = c
		}
		if str == "" {
			str = tmp
		} else {
			str = str + seq + tmp
		}
	}
	return str
}

//md5加密
func (ds *Dstring) Md5Str(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
`
)
