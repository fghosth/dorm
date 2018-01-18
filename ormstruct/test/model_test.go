package ormstruct_test

import (
	"fmt"
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
func TestARSelect(t *testing.T) {

	ar := ormstruct.NewHsAuthRecords("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
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
	ar := ormstruct.NewHsAuthRecords("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
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
	hsAuthRecords := ormstruct.NewHsAuthRecords("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
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
