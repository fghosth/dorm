package lexer_test

import (
	"fmt"
	"testing"

	"jvole.com/createProject/lexer"
)

var cocksqlStr string
var cockDBlexer lexer.CockDBLexer

func init() {
	cockDBlexer := new(lexer.CockDBLexer)
	cocksqlStr = cockDBlexer.SqlString("../cockroach.sql")
}

func TestCreateSqlByStruct(t *testing.T) {
	obj := new(lexer.ProductInformation)
	str := cockDBlexer.CreateSqlByStruct(obj)
	fmt.Println(str)
}

// func TestCreateStruct(t *testing.T) {
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

// func TestField(t *testing.T) {
// 	str := cockDBlexer.CreateTableString(cocksqlStr)
// 	for _, v := range str {
// 		name := cockDBlexer.Field(v)
// 		fmt.Println(name)
// 	}
// }
//
//
// func TestTableName(t *testing.T) {
// 	str := cockDBlexer.CreateTableString(cocksqlStr)
// 	for _, v := range str {
// 		tname := cockDBlexer.TableName(v)
// 		fmt.Println("tableName:", tname)
// 	}
// }
