package model_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/k0kubun/pp"
	"jvole.com/createProject/model"
	"jvole.com/createProject/model/base"
	"jvole.com/createProject/util"
)

var ut *util.Dstring
var hsAuthApplication = model.NewHsAuthApplicationDao()

func init() {
	ut = new(util.Dstring)
}
func TestAASDEL(t *testing.T) {
	// hsAuthApplication.SetDBConn("cockroachDB", "postgresql://derek:123456@localhost:26257/auth?sslmode=disable")
	base.AddBeforeFun(func() { fmt.Println("SDEL------before") }, "SDEL")
	base.AddBeforeFun(func() { fmt.Println("SDEL==============before") }, "SDEL")
	base.AddAfterFun(func() { fmt.Println("SDEL==============after") }, "SDEL")
	base.AddAfterFun(func() { fmt.Println("SDEL=======dd=======after") }, "SDEL")
	_, err := hsAuthApplication.FindByID(316654330975027201)
	pp.Println(hsAuthApplication)
	ut.Checkerr(err)
	res, err := hsAuthApplication.SDelete()
	ut.Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthApplication.GetSql())
}

func TestAASDeleteBatch(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("DeleteBatch------before") }, "DeleteBatch")
	base.AddBeforeFun(func() { fmt.Println("DeleteBatch==============before") }, "DeleteBatch")
	base.AddAfterFun(func() { fmt.Println("DeleteBatch==============after") }, "DeleteBatch")
	base.AddAfterFun(func() { fmt.Println("DeleteBatch=======dd=======after") }, "DeleteBatch")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 915832796958982145
	result, err := hsAuthApplication.Select("and id>$1 and id<$2", 2, 0, args...)
	ut.Checkerr(err)

	err = hsAuthApplication.SDeleteBatch(result)
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthApplication.GetSql())
}

func TestAAExec(t *testing.T) {

	// hsAuthApplication.SetDBConn("cockroachDB", "postgresql://derek:123456@localhost:26257/auth?sslmode=disable")
	base.AddBeforeFun(func() { fmt.Println("Exec------before") }, "Exec")
	base.AddBeforeFun(func() { fmt.Println("Exec==============before") }, "Exec")
	base.AddAfterFun(func() { fmt.Println("Exec==============after") }, "Exec")
	base.AddAfterFun(func() { fmt.Println("Exec=======dd=======after") }, "Exec")
	args := make([]interface{}, 1)
	args[0] = 316622689471660033
	res, err := hsAuthApplication.Exec("update hs_auth_application set status_at=1 where id=$1", args...)
	ut.Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthApplication.GetSql())
}

func TestAAdeleteBatch(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("DeleteBatch------before") }, "DeleteBatch")
	base.AddBeforeFun(func() { fmt.Println("DeleteBatch==============before") }, "DeleteBatch")
	base.AddAfterFun(func() { fmt.Println("DeleteBatch==============after") }, "DeleteBatch")
	base.AddAfterFun(func() { fmt.Println("DeleteBatch=======dd=======after") }, "DeleteBatch")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 915832796958982145
	result, err := hsAuthApplication.Select("and id>$1 and id<$2", 2, 0, args...)
	ut.Checkerr(err)

	err = hsAuthApplication.DeleteBatch(result)
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthApplication.GetSql())
}

func TestAADelete(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("Delete------before") }, "Delete")
	base.AddBeforeFun(func() { fmt.Println("Delete==============before") }, "Delete")
	base.AddAfterFun(func() { fmt.Println("Delete==============after") }, "Delete")
	base.AddAfterFun(func() { fmt.Println("Delete=======dd=======after") }, "Delete")
	_, err := hsAuthApplication.FindByID(316622689496596481)
	ut.Checkerr(err)
	res, err := hsAuthApplication.Delete()
	ut.Checkerr(err)
	pp.Println(res)
	pp.Println(hsAuthApplication.GetSql())
}

func TestAAupdateBatch(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("UpdateBatch------before") }, "UpdateBatch")
	base.AddBeforeFun(func() { fmt.Println("UpdateBatch==============before") }, "UpdateBatch")
	base.AddAfterFun(func() { fmt.Println("UpdateBatch==============after") }, "UpdateBatch")
	base.AddAfterFun(func() { fmt.Println("UpdateBatch=======dd=======after") }, "UpdateBatch")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 915832796958982145
	result, err := hsAuthApplication.Select("and id>$1 and id<$2", 100, 2, args...)
	ut.Checkerr(err)
	for i := 0; i < len(result); i++ {
		v := result[i].(base.HsAuthApplication)
		v.Ip = "19.168.10." + strconv.Itoa(i)
		result[i] = v
	}
	// pp.Println(result)
	err = hsAuthApplication.UpdateBatch(result)
	ut.Checkerr(err)

	pp.Println(hsAuthApplication.GetSql())
}

func TestAAUpdate(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("Update------before") }, "Update")
	base.AddBeforeFun(func() { fmt.Println("Update==============before") }, "Update")
	base.AddAfterFun(func() { fmt.Println("Update==============after") }, "Update")
	base.AddAfterFun(func() { fmt.Println("Update=======dd=======after") }, "Update")
	_, err := hsAuthApplication.FindByID(3166226894716600333)
	ut.Checkerr(err)
	hsAuthApplication.AppKey = "11ffffff111333dd33333"
	hsAuthApplication.DeletedAt = "2016-11-19 02:04:25+00:00"

	// pp.Println(hsAuthApplication)
	res, err := hsAuthApplication.Update()
	pp.Println(hsAuthApplication)
	ut.Checkerr(err)

	pp.Println(hsAuthApplication.GetSql())
	pp.Println(res)
}

func TestAAaddBatch(t *testing.T) {
	// hsAuthApplication.SetDBConn("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
	base.AddBeforeFun(func() { fmt.Println("AddBatch----before") }, "AddBatch")
	base.AddBeforeFun(func() { fmt.Println("AddBatch==============before") }, "AddBatch")
	base.AddAfterFun(func() { fmt.Println("AddBatch======end========after") }, "AddBatch")
	base.AddAfterFun(func() { fmt.Println("AddBatch========endddd======after") }, "AddBatch")
	ha := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		ar := base.HsAuthApplication{}
		ar.SecretKey = "111adefwe2rqereee" + strconv.Itoa(i)
		ar.AppKey = "11123e412e2ee3" + strconv.Itoa(i)
		ar.Name = "sed" + strconv.Itoa(i)
		ar.Ip = "192.143.11.11"
		ar.Exp = 3
		ar.Type = 1
		ar.CreatedAt = "2001-01-14"
		ar.UpdatedAt = "2001-01-14"
		ar.DeletedAt = "2001-01-14"
		ar.StatusAt = 1
		ha[i] = ar
	}

	err := hsAuthApplication.AddBatch(ha)
	ut.Checkerr(err)

	pp.Println(hsAuthApplication.GetSql())
}

func TestAASelect(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("sdfsdfsdfsdfdfs") }, "Select")
	base.AddBeforeFun(func() { fmt.Println("==============") }, "Selet")
	base.AddAfterFun(func() { fmt.Println("======end========") }, "Select")
	base.AddAfterFun(func() { fmt.Println("========endddd======") }, "Select")
	args := make([]interface{}, 2)
	args[0] = 10
	args[1] = 916583981147095041
	// base.UseCache(false)
	// base.SetCacheTime(2)
	t1 := time.Now() // get current time
	_, err := hsAuthApplication.Select("and id>$1 and id<$2", 2000, 2, args...)
	elapsed := time.Since(t1)
	fmt.Println("App elapsed1: ", elapsed)
	ut.Checkerr(err)
	// time.Sleep(time.Second * 1)
	//是否使用缓存
	t2 := time.Now() // get current time
	_, err = hsAuthApplication.Select("and id>$1 and id<$2", 2000, 2, args...)
	elapsed2 := time.Since(t2)
	fmt.Println("App elapsed2: ", elapsed2)
	ut.Checkerr(err)
	// pp.Println(result)
	pp.Println(hsAuthApplication.GetSql())
}

func TestAAFindByID(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("FindByID------before") }, "FindByID")
	base.AddBeforeFun(func() { fmt.Println("FindByID==============before") }, "FindByID")
	base.AddAfterFun(func() { fmt.Println("FindByID==============after") }, "FindByID")
	base.AddAfterFun(func() { fmt.Println("FindByID=======dd=======after") }, "FindByID")
	base.UseCache(false)

	t1 := time.Now() // get current time
	_, err := hsAuthApplication.FindByID(316583980943048705)
	elapsed := time.Since(t1)
	fmt.Println("App elapsed1: ", elapsed)
	ut.Checkerr(err)
	//是否使用缓存
	t2 := time.Now() // get current time
	_, err = hsAuthApplication.FindByID(316654331186282497)
	elapsed2 := time.Since(t2)
	fmt.Println("App elapsed2: ", elapsed2)

	pp.Println(base.GetCacheRate())
	pp.Println(hsAuthApplication.GetSql())
}

func TestAAAdd(t *testing.T) {
	base.AddBeforeFun(func() { fmt.Println("Add------before") }, "Add")
	base.AddBeforeFun(func() { fmt.Println("Add==============before") }, "Add")
	base.AddAfterFun(func() { fmt.Println("Add==============after") }, "Add")
	base.AddAfterFun(func() { fmt.Println("Add=======dd=======after") }, "Add")
	hsAuthApplication.SecretKey = "w3erqeer332"
	hsAuthApplication.AppKey = "41323fsedfasdfw1"
	hsAuthApplication.Name = "ssa3df2e3423e1"
	hsAuthApplication.Ip = "192.143.11.112"
	hsAuthApplication.Exp = 100
	hsAuthApplication.Type = 1
	hsAuthApplication.CreatedAt = "2001-01-14"
	hsAuthApplication.UpdatedAt = "2001-01-14"
	hsAuthApplication.DeletedAt = "2001-01-14"
	hsAuthApplication.StatusAt = 1
	result, err := hsAuthApplication.Add()
	ut.Checkerr(err)
	pp.Println(result)
	pp.Println(hsAuthApplication.GetSql())
}
