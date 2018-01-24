package model_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/k0kubun/pp"
	"jvole.com/createProject/model"
	"jvole.com/createProject/model/base"
	"jvole.com/createProject/util"
)

// var ut *util.Dstring
var hsAuthRecords = model.NewHsAuthRecordsDao()

func init() {
	ut = new(util.Dstring)
}

func TestAUExec(t *testing.T) {

	// hsAuthRecords.SetDBConn("cockroachDB", "postgresql://derek:123456@localhost:26257/auth?sslmode=disable")
	base.AddBeforeFun(func() { fmt.Println("Exec------before") }, "Exec")
	base.AddBeforeFun(func() { fmt.Println("Exec==============before") }, "Exec")
	base.AddAfterFun(func() { fmt.Println("Exec==============after") }, "Exec")
	base.AddAfterFun(func() { fmt.Println("Exec=======dd=======after") }, "Exec")
	args := make([]interface{}, 1)
	args[0] = 315583830360326145
	res, err := hsAuthRecords.Exec("update hs_auth_application set status_at=1 where id=$1", args...)
	ut.Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthRecords.GetSql())
}

func TestAUDeleteBatch(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("DeleteBatch------before") }, "DeleteBatch")
	base.AddBeforeFun(func() { fmt.Println("DeleteBatch==============before") }, "DeleteBatch")
	base.AddAfterFun(func() { fmt.Println("DeleteBatch==============after") }, "DeleteBatch")
	base.AddAfterFun(func() { fmt.Println("DeleteBatch=======dd=======after") }, "DeleteBatch")
	args := make([]interface{}, 2)
	args[0] = 10000
	args[1] = 316581278803165185
	result, err := hsAuthRecords.Select("where id>$1 and id<$2", 2, 0, args...)
	ut.Checkerr(err)

	err = hsAuthRecords.DeleteBatch(result)
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}

func TestAUDelete(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("Delete------before") }, "Delete")
	base.AddBeforeFun(func() { fmt.Println("Delete==============before") }, "Delete")
	base.AddAfterFun(func() { fmt.Println("Delete==============after") }, "Delete")
	base.AddAfterFun(func() { fmt.Println("Delete=======dd=======after") }, "Delete")
	_, err := hsAuthRecords.FindByID(315832796961341441)
	ut.Checkerr(err)
	res, err := hsAuthRecords.Delete()
	ut.Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthRecords.GetSql())
}

func TestAUUpdateBatch(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("UpdateBatch------before") }, "UpdateBatch")
	base.AddBeforeFun(func() { fmt.Println("UpdateBatch==============before") }, "UpdateBatch")
	base.AddAfterFun(func() { fmt.Println("UpdateBatch==============after") }, "UpdateBatch")
	base.AddAfterFun(func() { fmt.Println("UpdateBatch=======dd=======after") }, "UpdateBatch")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 315832796958982145
	result, err := hsAuthRecords.Select("where id>$1 and id<$2", 100, 2, args...)
	ut.Checkerr(err)
	for i := 0; i < len(result); i++ {
		v := result[i].(base.HsAuthRecords)
		v.Ip = "192.168.10." + strconv.Itoa(i)
		result[i] = v
	}

	err = hsAuthRecords.UpdateBatch(result)
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}

func TestAUUpdate(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("Update------before") }, "Update")
	base.AddBeforeFun(func() { fmt.Println("Update==============before") }, "Update")
	base.AddAfterFun(func() { fmt.Println("Update==============after") }, "Update")
	base.AddAfterFun(func() { fmt.Println("Update=======dd=======after") }, "Update")
	_, err := hsAuthRecords.FindByID(2)
	ut.Checkerr(err)
	hsAuthRecords.AppKey = "111333dd"
	hsAuthRecords.DeletedAt = "2016-11-19 02:04:25+00:00"

	// pp.Println(hsAuthRecords)
	res, err := hsAuthRecords.Update()
	pp.Println(hsAuthRecords.AppKey)
	ut.Checkerr(err)

	pp.Println(hsAuthRecords.GetSql())
	pp.Println(res)
}

// func TestAUAddBatch(t *testing.T) {
// 	// hsAuthRecords.SetDBConn("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
// 	base.AddBeforeFun(func() { fmt.Println("AddBatch----before") }, "AddBatch")
// 	base.AddBeforeFun(func() { fmt.Println("AddBatch==============before") }, "AddBatch")
// 	base.AddAfterFun(func() { fmt.Println("AddBatch======end========after") }, "AddBatch")
// 	base.AddAfterFun(func() { fmt.Println("AddBatch========endddd======after") }, "AddBatch")
// 	ha := make([]interface{}, 100)
// 	for i := 0; i < 100; i++ {
// 		ar := base.HsAuthRecords{}
// 		ar.SecretKey = "adfwerqer" + strconv.Itoa(i)
// 		ar.AppKey = "1234123" + strconv.Itoa(i)
// 		ar.Iat = "2001-01-14"
// 		ar.Ip = "192.143.11.11"
// 		ar.Exp = "2001-01-14"
// 		ar.Type = 1
// 		ar.CreatedAt = "2001-01-14"
// 		ar.UpdatedAt = "2001-01-14"
// 		ar.DeletedAt = "2001-01-14"
// 		ar.StatusAt = 1
// 		ha[i] = ar
// 	}
// 	for i := 0; i < 10; i++ {
// 		err := hsAuthRecords.AddBatch(ha)
// 		ut.Checkerr(err)
// 	}
//
// 	pp.Println(hsAuthRecords.GetSql())
// }

func TestAUSelect(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("sdfsdfsdfsdfdfs") }, "Select")
	base.AddBeforeFun(func() { fmt.Println("==============") }, "Selet")
	base.AddAfterFun(func() { fmt.Println("======end========") }, "Select")
	base.AddAfterFun(func() { fmt.Println("========endddd======") }, "Select")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 1049
	result, err := hsAuthRecords.Select("where id>$1 and id<$2", 10, 2, args...)
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}

func TestAUFindByID(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("FindByID------before") }, "FindByID")
	base.AddBeforeFun(func() { fmt.Println("FindByID==============before") }, "FindByID")
	base.AddAfterFun(func() { fmt.Println("FindByID==============after") }, "FindByID")
	base.AddAfterFun(func() { fmt.Println("FindByID=======dd=======after") }, "FindByID")

	result, err := hsAuthRecords.FindByID(2)
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}

func TestAUAdd(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("Add------before") }, "Add")
	base.AddBeforeFun(func() { fmt.Println("Add==============before") }, "Add")
	base.AddAfterFun(func() { fmt.Println("Add==============after") }, "Add")
	base.AddAfterFun(func() { fmt.Println("Add=======dd=======after") }, "Add")
	hsAuthRecords.SecretKey = "adfwerqer33"
	hsAuthRecords.AppKey = "1234123fsdfasdfw"
	hsAuthRecords.Sign = "dasfasd"
	hsAuthRecords.Token = "dddd"
	hsAuthRecords.Alg = "dddeee"
	hsAuthRecords.Iat = "sdafsadf23423e"
	hsAuthRecords.Ip = "192.143.11.11"
	hsAuthRecords.Exp = "100"
	hsAuthRecords.Type = 1
	hsAuthRecords.CreatedAt = "2001-01-14"
	hsAuthRecords.UpdatedAt = "2001-01-14"
	hsAuthRecords.DeletedAt = "2001-01-14"
	hsAuthRecords.StatusAt = 1
	result, err := hsAuthRecords.Add()
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthRecords.GetSql())
}
