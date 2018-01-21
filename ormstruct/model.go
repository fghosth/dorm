package ormstruct

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
	DB        *sql.DB //数据库连接
	Beforefun Before
	Afterfun  After
	Driver    string
)

func init() {
	// SetConn("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
	SetConn("cockroachDB", "postgresql://derek:123456@localhost:26257/auth?sslmode=disable")
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
