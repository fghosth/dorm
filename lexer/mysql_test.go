package lexer_test

import (
	"github.com/k0kubun/pp"
	"jvole.com/createProject/lexer"
	"testing"
)

var sqlStr string
var mysqlLexer lexer.MysqlLexer

func init() {
	mysqlLexer := new(lexer.MysqlLexer)
	sqlStr = mysqlLexer.SqlString("../reports.sql")
	//sqlStr = mysqlLexer.SqlString("../auth.sql")
}

// func TestCreateCockInsertSqlFromMysql(t *testing.T) {
// 	str := mysqlLexer.InsertStr(sqlStr)
// 	name := mysqlLexer.CreateCockInsertSqlFromMysql(str)
// 	fmt.Println(name)
//
// }

// func TestCreateCockroachSqlFromMysql(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		name := mysqlLexer.CreateCockroachSqlFromMysql(v)
// 		fmt.Println(name)
// 	}
// }

// func TestInsertStr(t *testing.T) {
// 	str := mysqlLexer.InsertStr(sqlStr)
// 	pp.Println(str)
// }

// func TestCreateSqlByStructStr(t *testing.T) {
// 	sl := new(lexer.StructLexer)
// 	fileStr := sl.GetStructFile("../ormstruct/hs_auth_permission.go")
// 	slist := sl.StructStr(fileStr)
// 	for _, v := range slist {
// 		str := mysqlLexer.CreateSqlByStructStr(v)
// 		fmt.Println(str)
// 	}
//
// }

// func TestCreateSqlByStruct(t *testing.T) {
// 	user := new(lexer.User)
// 	str := mysqlLexer.CreateSqlByStruct(user)
// 	fmt.Println(str)
// }

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
//
func TestField(t *testing.T) {
	str := mysqlLexer.CreateTableString(sqlStr)
	for _, v := range str {
		name := mysqlLexer.Field(v)
		pp.Println(name)
	}
}

// func TestUQName(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		name := mysqlLexer.Uniquekey(v)
// 		fmt.Println("UQName:", name)
// 	}
// }

// func TestKeyIndex(t *testing.T) {
// 	str := mysqlLexer.CreateTableString(sqlStr)
// 	for _, v := range str {
// 		name := mysqlLexer.IndexKey(v)
// 		fmt.Println("IndexName:", name)
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
