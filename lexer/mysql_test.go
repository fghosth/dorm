package lexer_test

import (
	"fmt"
	"testing"

	"jvole.com/createProject/lexer"
)

var sqlStr string
var mysqlLexer lexer.MysqlLexer

func init() {
	mysqlLexer := new(lexer.MysqlLexer)
	sqlStr = mysqlLexer.SqlString("../orm.sql")
}

func TestCreateSqlByStruct(t *testing.T) {
	user := new(lexer.User)
	str := mysqlLexer.CreateSqlByStruct(user)
	fmt.Println(str)
}

//
// func TestCreateStruct(t *testing.T) {
// 	// tools := new(util.Dstring)
// 	// fmt.Println(tools.CalToUnder("User"))
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		tname := mysqlLexer.TableName(v)
// 		field := mysqlLexer.Field(v)
// 		structStr := mysqlLexer.CreateStruct("usermodel", tname, field)
// 		fmt.Println(structStr)
// 	}
// }

// func TestField(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		name := mysqlLexer.Field(v)
// 		fmt.Println(name)
// 	}
// }

// func TestPKName(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		name := mysqlLexer.Primarykey(v)
// 		fmt.Println("pkName:", name)
// 	}
// }
//
// func TestTableName(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		tname := mysqlLexer.TableName(v)
// 		fmt.Println("tableName:", tname)
// 	}
// }
//
// func TestCreateTableString(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	pp.Println(str)
// }
