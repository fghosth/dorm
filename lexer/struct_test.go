package lexer_test

import (
	"fmt"
	"testing"

	"jvole.com/createProject/lexer"
)

var (
	sl = new(lexer.StructLexer)
)

// func TestStructName(t *testing.T) {
// 	fileStr := sl.GetStructFile("../ormstruct/hs_auth_permission.go")
// 	slist := sl.StructStr(fileStr)
// 	for _, v := range slist {
// 		str := sl.StructName(v)
// 		fmt.Println(str)
// 	}
// }

func TestStructField(t *testing.T) {
	fileStr := sl.GetStructFile("../ormstruct/hs_auth_permission.go")
	slist := sl.StructStr(fileStr)
	for _, v := range slist {
		str := sl.FieldName(v)
		fmt.Println(str)
	}
}

// func TestStructTagLex(t *testing.T) {
// 	fileStr := sl.GetStructFile("../ormstruct/hs_auth_permission.go")
// 	slist := sl.StructStr(fileStr)
// 	for _, v := range slist {
// 		str := sl.FieldName(v)
// 		for _, tv := range str {
// 			tag := sl.Taglex(tv["tag"])
// 			pp.Println(tag)
// 		}
// 	}
// }
