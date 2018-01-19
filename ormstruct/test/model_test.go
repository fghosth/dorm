package ormstruct_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/k0kubun/pp"
	"jvole.com/createProject/ormstruct"
)

func init() {

}

func Checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func TestARExec(t *testing.T) {
	hsAuthRecords := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("Exec------before") }, "Exec")
	ormstruct.AddBeforeFun(func() { fmt.Println("Exec==============before") }, "Exec")
	ormstruct.AddAfterFun(func() { fmt.Println("Exec==============after") }, "Exec")
	ormstruct.AddAfterFun(func() { fmt.Println("Exec=======dd=======after") }, "Exec")
	args := make([]interface{}, 1)
	args[0] = 100
	res, err := hsAuthRecords.Exec("update hs_auth_records set status_at=1 where id=?", args...)
	Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthRecords.GetSql())
}

func TestARDeleteBatch(t *testing.T) {
	hsAuthRecords := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("DeleteBatch------before") }, "DeleteBatch")
	ormstruct.AddBeforeFun(func() { fmt.Println("DeleteBatch==============before") }, "DeleteBatch")
	ormstruct.AddAfterFun(func() { fmt.Println("DeleteBatch==============after") }, "DeleteBatch")
	ormstruct.AddAfterFun(func() { fmt.Println("DeleteBatch=======dd=======after") }, "DeleteBatch")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 1099
	result, err := hsAuthRecords.Select("where id>? and id<?", 2, 0, args...)
	Checkerr(err)

	err = hsAuthRecords.DeleteBatch(result)
	Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}

func TestARDelete(t *testing.T) {
	hsAuthRecords := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("Delete------before") }, "Delete")
	ormstruct.AddBeforeFun(func() { fmt.Println("Delete==============before") }, "Delete")
	ormstruct.AddAfterFun(func() { fmt.Println("Delete==============after") }, "Delete")
	ormstruct.AddAfterFun(func() { fmt.Println("Delete=======dd=======after") }, "Delete")
	_, err := hsAuthRecords.FindByID(200)
	Checkerr(err)
	res, err := hsAuthRecords.Delete()
	Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthRecords.GetSql())
}

func TestARUpdateBatch(t *testing.T) {
	hsAuthRecords := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("UpdateBatch------before") }, "UpdateBatch")
	ormstruct.AddBeforeFun(func() { fmt.Println("UpdateBatch==============before") }, "UpdateBatch")
	ormstruct.AddAfterFun(func() { fmt.Println("UpdateBatch==============after") }, "UpdateBatch")
	ormstruct.AddAfterFun(func() { fmt.Println("UpdateBatch=======dd=======after") }, "UpdateBatch")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 1099
	result, err := hsAuthRecords.Select("where id>? and id<?", 100, 2, args...)
	Checkerr(err)
	for i := 0; i < len(result); i++ {
		v := result[i].(ormstruct.HsAuthRecords)
		v.Alg = "aa-" + strconv.Itoa(i)
		result[i] = v
	}

	err = hsAuthRecords.UpdateBatch(result)
	Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}

func TestARUpdate(t *testing.T) {
	hsAuthRecords := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("Update------before") }, "Update")
	ormstruct.AddBeforeFun(func() { fmt.Println("Update==============before") }, "Update")
	ormstruct.AddAfterFun(func() { fmt.Println("Update==============after") }, "Update")
	ormstruct.AddAfterFun(func() { fmt.Println("Update=======dd=======after") }, "Update")
	_, err := hsAuthRecords.FindByID(205)
	Checkerr(err)

	hsAuthRecords.AppKey = "111333"
	res, err := hsAuthRecords.Update()
	Checkerr(err)
	pp.Println(hsAuthRecords)
	pp.Println(hsAuthRecords.GetSql())
	pp.Println(res)
}

// func TestARAddBatch(t *testing.T) {
// 	hsAuthRecords := ormstruct.NewHsAuthRecords()
// 	ormstruct.AddBeforeFun(func() { fmt.Println("AddBatch----before") }, "AddBatch")
// 	ormstruct.AddBeforeFun(func() { fmt.Println("AddBatch==============before") }, "AddBatch")
// 	ormstruct.AddAfterFun(func() { fmt.Println("AddBatch======end========after") }, "AddBatch")
// 	ormstruct.AddAfterFun(func() { fmt.Println("AddBatch========endddd======after") }, "AddBatch")
// 	ha := make([]interface{}, 100)
// 	for i := 0; i < 100; i++ {
// 		ar := ormstruct.HsAuthRecords{}
// 		ar.SecretKey = "adfwerqer" + strconv.Itoa(i)
// 		ar.AppKey = "1234123fsdfasdf" + strconv.Itoa(i)
// 		ar.Sign = "sdafsadf23423" + strconv.Itoa(i)
// 		ar.Token = "sdfasdfwer" + strconv.Itoa(i)
// 		ar.Alg = "sdf" + strconv.Itoa(i)
// 		ar.Ip = "192.143.11.11"
// 		ar.Exp = "2001-01-14"
// 		ar.Iat = "2001-01-14"
// 		ar.Type = 1
// 		ar.CreatedAt = "2001-01-14"
// 		ar.UpdatedAt = "2001-01-14"
// 		ar.DeletedAt = "2001-01-14"
// 		ar.StatusAt = 1
// 		ha[i] = ar
// 	}
// 	err := hsAuthRecords.AddBatch(ha)
// 	Checkerr(err)
//
// 	pp.Println(hsAuthRecords.GetSql())
// }

func TestARSelect(t *testing.T) {
	ar := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("sdfsdfsdfsdfdfs") }, "Select")
	ormstruct.AddBeforeFun(func() { fmt.Println("==============") }, "Selet")
	ormstruct.AddAfterFun(func() { fmt.Println("======end========") }, "Select")
	ormstruct.AddAfterFun(func() { fmt.Println("========endddd======") }, "Select")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 1099
	result, err := ar.Select("where id>? and id<?", 100, 2, args...)
	Checkerr(err)
	pp.Println(result)
	pp.Println(ar.GetSql())
}

func TestARFindByID(t *testing.T) {
	ar := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("FindByID------before") }, "FindByID")
	ormstruct.AddBeforeFun(func() { fmt.Println("FindByID==============before") }, "FindByID")
	ormstruct.AddAfterFun(func() { fmt.Println("FindByID==============after") }, "FindByID")
	ormstruct.AddAfterFun(func() { fmt.Println("FindByID=======dd=======after") }, "FindByID")

	result, err := ar.FindByID(20)
	Checkerr(err)
	pp.Println(result)
	pp.Println(ar.GetSql())
}

func TestARAdd(t *testing.T) {
	hsAuthRecords := ormstruct.NewHsAuthRecords()
	ormstruct.AddBeforeFun(func() { fmt.Println("Add------before") }, "Add")
	ormstruct.AddBeforeFun(func() { fmt.Println("Add==============before") }, "Add")
	ormstruct.AddAfterFun(func() { fmt.Println("Add==============after") }, "Add")
	ormstruct.AddAfterFun(func() { fmt.Println("Add=======dd=======after") }, "Add")
	hsAuthRecords.SecretKey = "adfwerqer"
	hsAuthRecords.AppKey = "1234123fsdfasdf"
	hsAuthRecords.Sign = "sdafsadf23423"
	hsAuthRecords.Token = "sdfasdfwer"
	hsAuthRecords.Alg = "sdf"
	hsAuthRecords.Ip = "192.143.11.11"
	hsAuthRecords.Exp = "2001-01-14"
	hsAuthRecords.Iat = "2001-01-14"
	hsAuthRecords.Type = 1
	hsAuthRecords.CreatedAt = "2001-01-14"
	hsAuthRecords.UpdatedAt = "2001-01-14"
	hsAuthRecords.DeletedAt = "2001-01-14"
	hsAuthRecords.StatusAt = 1
	result, err := hsAuthRecords.Add()
	Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}
