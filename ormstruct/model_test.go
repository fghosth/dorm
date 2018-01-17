package ormstruct

import (
	"fmt"
	"testing"

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
	ormstruct.SetConn("mysql", "root:@tcp(localhost:3306)/praise_auth?charset=utf8")
	ar := new(ormstruct.HsAuthRecords)
	ormstruct.AddBeforeFun(func() { fmt.Println("sdfsdfsdfsdfdfs") }, "Select")
	ormstruct.AddBeforeFun(func() { fmt.Println("==============") }, "Selet")
	ormstruct.AddAfterFun(func() { fmt.Println("======end========") }, "Select")
	ormstruct.AddAfterFun(func() { fmt.Println("========endddd======") }, "Select")
	args := make([]interface{}, 1)
	args[0] = 10
	_, err := ar.Select("where id>? ", 10, 2, args...)
	Checkerr(err)
	// pp.Println(result)
}
