package ormstruct

type Model interface {
	/*
			   根据条件查找结果集
			   @parm sql 除去select xxx,xxx from tablename 之后的东西
			   @parm value sql中?值 可以为空
		     @return struct 集合
		     @return error 错误
	*/
	Select(sql string, value ...string) ([]interface{}, error)
	/*
			   根据主键查找
			   @parm id 主键
		     @return struct
		     @return error 错误
	*/
	FindById(id string) (interface{}, error)
	/*
			   根据自身struct内容添加
			   @parm
		     @return 返回主键id
		     @return error 错误
	*/
	Add() (string, error)
	/*
			   批量添加
			   @parm struct数组
		     @return error 错误
	*/
	AddBatch(obj []interface{}) error
	/*
			   根据自身struct更新
			   @parm
		     @return error 错误
	*/
	Update() error
	/*
			   批量更新
			   @parm struct数组
		     @return error 错误
	*/
	UpdateBatch(obj []interface{}) error
	/*
			   根据自身struct删除
			   @parm
		     @return error 错误
	*/
	Delete() error
	/*
			   批量删除
			   @parm sql sql语句，valuesql语句中?的部分，可以为空
		     @return error 错误
	*/

	DeleteBatch(sql string, value ...string) error
	/*
			   执行sql语句 非查询的语句
			   @parm sql sql语句，valuesql语句中?的部分，可以为空
		     @return error 错误
	*/
	Exec(sql string, value ...string) error
}

type BeforeSelect []func()
type AfterSelect []func()
