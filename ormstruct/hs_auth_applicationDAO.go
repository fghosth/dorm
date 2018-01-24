package model

import (
	"jvole.com/createProject/ormstruct/base"
)

type HsAuthApplicationDao struct {
	model base.Model
	base.HsAuthApplication
}

/*
 根据条件查找结果集
 @parm sql 除去select xxx,xxx from tablename 之后的东西
 @parm value sql中?值 可以为空
 @parm limit 显示数量
 @parm offset 数据位置0开始
 @return struct 集合
 @return error 错误
*/
func (dao HsAuthApplicationDao) Select(sql string, limit, offset int, value ...interface{}) ([]interface{}, error) {
	return dao.model.Select(sql, limit, offset, value...)
}

/*
 根据主键查找
 @parm id 主键
 @return struct
 @return error 错误
*/
func (dao *HsAuthApplicationDao) FindByID(id int64) (interface{}, error) {
	res, err := dao.model.FindByID(id)
	re := res.(*base.HsAuthApplication)
	dao.HsAuthApplication = *re
	return res, err
}

/*
 根据自身struct内容添加
 @parm
 @return 返回主键id
 @return error 错误
*/
func (dao HsAuthApplicationDao) Add() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.Add()
}

/*
 批量添加
 @parm struct数组
 @return error 错误
*/
func (dao HsAuthApplicationDao) AddBatch(obj []interface{}) error {
	return dao.model.AddBatch(obj)
}

/*
 根据自身struct更新
 @parm
 @return int64 修改记录的id
 @return error 错误
*/
func (dao HsAuthApplicationDao) Update() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.Update()
}

/*
 批量更新
 @parm struct数组
 @return error 错误
*/
func (dao HsAuthApplicationDao) UpdateBatch(obj []interface{}) error {
	return dao.model.UpdateBatch(obj)
}

/*
 根据自身struct删除
 @parm
 @return int64 影响行数
 @return error 错误
*/
func (dao HsAuthApplicationDao) Delete() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.Delete()
}

/*
 根据自身struct软删除
 @parm
 @return int64 影响行数
 @return error 错误
*/
func (dao HsAuthApplicationDao) SDelete() (int64, error) {
	b := dao.getObjWithValue(dao)
	dao.model = &b
	return dao.model.SDelete()
}

/*
 批量软删除
 @parm struct struct数组
 @return error 错误
*/

func (dao HsAuthApplicationDao) SDeleteBatch(obj []interface{}) error {
	return dao.model.SDeleteBatch(obj)
}

/*
 批量删除
 @parm struct struct数组
 @return error 错误
*/
func (dao HsAuthApplicationDao) DeleteBatch(obj []interface{}) error {
	return dao.model.DeleteBatch(obj)
}

/*
 执行sql语句 非查询的语句
 @parm sql sql语句，valuesql语句中?的部分，可以为空
 @return int64 影响的行数
 @return error 错误
*/
func (dao HsAuthApplicationDao) Exec(sql string, value ...interface{}) (int64, error) {
	return dao.model.Exec(sql, value...)
}

/*
 获取最后执行的sql语句 和参数
 @return string sql语句和参数
*/
func (dao HsAuthApplicationDao) GetSql() (string, []interface{}) {
	return dao.model.GetSql()
}

/*
 设置当前对象的链接
 @db 数据库默认值mysql 支持mysql，mariadb，cockroachDB
 @str 数据库连接 『postgresql://derek:123456@localhost:26257/auth?sslmode=disable』 【root:@tcp(localhost:3306)/praise_auth?charset=utf8】
 @return int64 影响的行数
 @return error 错误
*/
func (dao HsAuthApplicationDao) SetDBConn(db, str string) {
	dao.model.SetDBConn(db, str)
}

//获取有值的对象
func (daoo HsAuthApplicationDao) getObjWithValue(dao HsAuthApplicationDao) base.HsAuthApplication {
	hsAuthApplication := base.NewHsAuthApplication()
	hsAuthApplication.Id = dao.Id
	hsAuthApplication.SecretKey = dao.SecretKey
	hsAuthApplication.AppKey = dao.AppKey
	hsAuthApplication.Name = dao.Name
	hsAuthApplication.Ip = dao.Ip
	hsAuthApplication.Type = dao.Type
	hsAuthApplication.Exp = dao.Exp
	hsAuthApplication.CreatedAt = dao.CreatedAt
	hsAuthApplication.UpdatedAt = dao.UpdatedAt
	hsAuthApplication.DeletedAt = dao.DeletedAt
	hsAuthApplication.StatusAt = dao.StatusAt
	return hsAuthApplication
}

func NewHsAuthApplicationDao() HsAuthApplicationDao {
	ap := base.NewHsAuthApplication()
	aa := base.HsAuthApplication{}
	return HsAuthApplicationDao{&ap, aa}
}
