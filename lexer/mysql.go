package lexer

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"jvole.com/createProject/util"
)

var (
	sqlType = `\b(tinyint|smallint|mediumint|int|integer|bigint|date|datetime|time|bit|tinytext|mediumtext|longtext|text|tinyblob|mediumblob|longblob|blob|float|double|decimal|timestamp|year|char|varchar|varbinary|enum|set|json)`
	// tableName = "(?<=TABLE[\\s]{1,200}`).{1,}(?=`)"
	tableNameLine = "(CREATE TABLE).+\\`"
	col           = "(?<=`).{1,}(?=`)"
	property      = `\b(NOT NULL|(DEFAULT ['a-zA-Z]+)|AUTO_INCREMENT|unsigned|zerofill|COMMENT ['a-zA-Z]+|PRIMARY.+,)`
)

var (
	TypeMap = map[string]string{
		"tinyint":    "int8",
		"smallint":   "int16",
		"mediumint":  "int32",
		"int":        "int32",
		"integer":    "int32",
		"bigint":     "int64",
		"date":       "string",
		"datetime":   "string",
		"time":       "string",
		"bit":        "int8",
		"bool":       "int16",
		"tinytext":   "string",
		"mediumtext": "string",
		"longtext":   "string",
		"text":       "string",
		"tinyblob":   "[]byte",
		"mediumblob": "[]byte",
		"longblob":   "[]byte",
		"blob":       "[]byte",
		"float":      "float32",
		"double":     "float64",
		"decimal":    "float64",
		"timestamp":  "int32",
		"year":       "string",
		"char":       "string",
		"varchar":    "string",
		"varbinary":  "[]byte",
		"enum":       "string",
		"set":        "string",
		"json":       "string",
	}
)

type MysqlLexer struct{}

var (
	ut = new(util.Dstring)
)

func (ml MysqlLexer) TableName() string {
	dat, err := ioutil.ReadFile("../orm.sql")

	if err != nil {
		fmt.Println(err)
	}

	r := regexp.MustCompile(tableNameLine)
	tname := r.FindString(string(dat))
	tname = ut.PixContent(tname, "`")
	fmt.Println("=======================", tname)
	return tname
}
