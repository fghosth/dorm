package lexer_test

import (
	"fmt"
	"jvole.com/createProject/lexer"
	"os"
	"testing"
)

var cocksqlStr string
var cockDBlexer lexer.CockDBLexer

func init() {
	cockDBlexer := new(lexer.CockDBLexer)
	cocksqlStr = cockDBlexer.SqlString("../cockroach.sql")
}

func TestTde(t *testing.T) {
	str, _ := os.Getwd()
	fmt.Println(str)
}

// func TestCockCreateSqlByStructStr(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile("../ormstruct/product_information.go")
// 	slist := sl.StructStr(fileStr)
// 	for _, v := range slist {
// 		str := cockDBlexer.CreateSqlByStructStr(v)
// 		fmt.Println(str)
// 	}
//
// }

// func TestCockCreateSqlByStruct(t *testing.T) {
// 	obj := new(lexer.ProductInformation)
// 	str := cockDBlexer.CreateSqlByStruct(obj)
// 	fmt.Println(str)
// }

// func TestCockCreateStruct(t *testing.T) {
// 	// tools := new(util.Dstring)
// 	// fmt.Println(tools.CalToUnder("User"))
// 	str := cockDBlexer.CreateTableString(cocksqlStr)
// 	for _, v := range str {
// 		tname := cockDBlexer.TableName(v)
// 		field := cockDBlexer.Field(v)
// 		structStr := cockDBlexer.CreateStruct("productmodel", tname, field)
// 		fmt.Println(structStr)
// 	}
// }

// func TestCockField(t *testing.T) {
// 	str := cockDBlexer.CreateTableString(cocksqlStr)
// 	for _, v := range str {
// 		name := cockDBlexer.Field(v)
// 		fmt.Println(name)
// 	}
// }
//
//
// func TestCockTableName(t *testing.T) {
// 	str := cockDBlexer.CreateTableString(cocksqlStr)
// 	for _, v := range str {
// 		tname := cockDBlexer.TableName(v)
// 		fmt.Println("tableName:", tname)
// 	}
// }
