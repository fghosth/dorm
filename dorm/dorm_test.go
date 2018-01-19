package dorm_test

import (
	"fmt"
	"testing"

	"jvole.com/createProject/dorm"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

var ut = new(util.Dstring)
var file = "../ormstruct/hs_auth_application.go"

func init() {

}

// func TestCreateModel(t *testing.T) {
//
// 	Str := dorm.CreateModel("ormstruct")
// 	fmt.Println(Str)
// }

func TestCreateDorm(t *testing.T) {
	sl := new(lexer.StructLexer)
	fileStr := sl.GetStructFile(file)
	arrStruct := sl.StructStr(fileStr)
	for _, v := range arrStruct {
		Str := dorm.CreateDorm("ormstruct", "mysql", v)
		fmt.Println(Str)
	}

}

func TestCreateField(t *testing.T) {
	sl := new(lexer.StructLexer)
	fileStr := sl.GetStructFile(file)
	arrStruct := sl.StructStr(fileStr)
	for _, v := range arrStruct {
		Str := dorm.CreateField(v)
		fmt.Println(Str)
	}
}

// func TestCreateHeader(t *testing.T) {
//
// 	Str := dorm.CreateHeader("dorm", "cockroachDB")
// 	// Str := dorm.CreateHeader("dorm", "mysql")
// 	fmt.Println(Str)
//
// }

// func TestCreateFunction(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateFunction(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateExec(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateExec(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateDeleteBatch(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateDeleteBatch(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateDelete(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateDelete(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateUpdateBatch(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateUpdateBatch(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateUpdate(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateUpdate(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateAddBatch(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateAddBatch(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateAdd(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateAdd(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateFindByID(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateFindByID(v)
// 		fmt.Println(Str)
// 	}
// }

// func TestCreateSelect(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile(file)
// 	arrStruct := sl.StructStr(fileStr)
// 	for _, v := range arrStruct {
// 		Str := dorm.CreateSelect(v)
// 		fmt.Println(Str)
// 	}
// }
