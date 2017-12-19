package lexer

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"

	"github.com/aymerick/raymond"
	"jvole.com/createProject/util"
)

var (
	//匹配所有mysql变量 匹配完整单词 解决 同时出现datetime 和date的问题
	sqlType = `\b(tinyint|smallint|mediumint|int|integer|bigint|date|datetime|time|bit|tinytext|mediumtext|longtext|text|tinyblob|mediumblob|longblob|blob|float|double|decimal|timestamp|year|char|varchar|varbinary|enum|set|json)\b`
	// tableName = "(?<=TABLE[\\s]{1,200}`).{1,}(?=`)"
	//获取tablename所在行(?i)忽略大小写
	tableNameLine = "(?i)(CREATE TABLE).+\\`"
	//匹配所有字段
	col = "`.+`"
	//匹配字段属性
	property = `(?i)\b(NOT NULL|(DEFAULT.+)|AUTO_INCREMENT|unsigned|zerofill|COMMENT.+'|PRIMARY.+,)`
	//找出所以cerate table代码段
	createTable = `(CREATE TABLE)[\W\w]+?\);`
	//为创造table的语句按字段分行
	colLine = `.+,`
	//找到PRIMARY KEY行
	primaryKeyLine = `(?i)(PRIMARY KEY).+`
	//找到index 行TODO
	indexLine = `(KEY).+`
)

var (
	StructToMysqlMap = map[string]string{
		"int":     "int",
		"int8":    "tinyint",
		"int16":   "smallint",
		"int32":   "int",
		"int64":   "bigint",
		"string":  "varchar",
		"float32": "float",
		"float64": "double",
		"[]byte":  "blob",
	}
	MysqlToStructMap = map[string]string{
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

type MysqlLexer struct {
}

var (
	ut  = new(util.Dstring)
	dat []byte
)

//根据struct生成数据库sql
func (ml MysqlLexer) CreateSqlByStruct(obj interface{}) string {
	var tableName, sql string
	flist := make([]string, 0)
	rtype := reflect.TypeOf(obj).Elem()
	tableName = rtype.Name()
	for k := 0; k < rtype.NumField(); k++ {
		colName := rtype.Field(k).Tag.Get("dormCol") //字段名
		if colName == "" {                           //tag中没有就用字段名
			colName = ut.CalToUnder(rtype.Field(k).Name)
		}

		colType := StructToMysqlMap[string(rtype.Field(k).Type.Kind().String())] //字段类型
		colProperty := dormToSql(rtype.Field(k).Tag.Get("dorm"))                 //字段属性
		tpms := "`" + colName + "` " + colType + " " + colProperty + ","
		flist = append(flist, tpms)
	}
	tableName, _ = ut.FUPer(tableName)
	ctx := map[string]interface{}{
		"tableName": tableName,
		"field":     flist,
	}

	sql, err := raymond.Render(MYSQL_SCRIPT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return sql
}

//处理属性转成sql语句能识别的字符串
func dormToSql(dorm string) string {
	sql := ""
	sql = strings.Replace(dorm, ";", " ", -1)
	sql = strings.Replace(sql, "-", " ", -1)
	return sql
}

//创建struct字符串
func (ml MysqlLexer) CreateStruct(packageName, tableName string, field []map[string]interface{}) string {
	var structString string
	fstr := make([]string, 0)
	for _, v := range field { //遍历字段
		tmpstr := ""
		for _, pv := range v["property"].([]string) { //dorm里面内容
			if tmpstr == "" {
				tmpstr = string(pv)
			} else {
				tmpstr = tmpstr + ";" + string(pv)
			}
		}
		str := ut.UnderToCal(v["colName"].(string)) + " " + v["goType"].(string) + "  `dormCol:\"" + v["colName"].(string) + "\" " + " dorm:\"" + tmpstr + "\"`"
		fstr = append(fstr, str)
	}
	ctx := map[string]interface{}{
		"name":        ut.FUpRLow(tableName),
		"packageName": packageName,
		"field":       fstr,
	}

	structString, err := raymond.Render(STRUCT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return structString
}

//获取字段及字段熟悉map
func (ml MysqlLexer) Field(tableStr string) []map[string]interface{} {
	field := make([]map[string]interface{}, 0)
	r := regexp.MustCompile(colLine)
	line := r.FindAllString(tableStr, -1)

	for _, v := range line {
		coln := getColnameByLine(v) //获取字段名

		if coln != "" { //字段名不能为空 不是字段
			colmap := make(map[string]interface{})
			colmap["colName"] = coln
			colmap["sqltype"], colmap["goType"] = getColTypeByLine(v)
			colmap["property"] = getColptyByLine(v)
			field = append(field, colmap)

		}
	}
	// fmt.Println(field)
	return field
}

//根据行获取熟悉数组
func getColptyByLine(str string) []string {
	pty := ""
	r := regexp.MustCompile(property)
	ptylist := r.FindAllString(str, -1)

	for _, v := range ptylist {
		if pty == "" {
			pty = strings.Replace(v, " ", "-", -1)
		} else {
			pty = pty + "||" + strings.Replace(v, " ", "_", -1)

		}
	}
	return strings.Split(pty, "||")
}

//根据行取出字段类型
func getColTypeByLine(str string) (sqltype, gotype string) {
	r := regexp.MustCompile(sqlType)
	sqltype = r.FindString(str)
	gotype = MysqlToStructMap[sqltype]
	return
}

//根据行取出字段名 如果不是字段行返回""
func getColnameByLine(str string) string {
	colname := ""

	isCol, err := regexp.MatchString(sqlType, str)
	if err != nil {
		fmt.Println(err)
	}
	if !isCol {
		return colname
	}

	r := regexp.MustCompile(col)

	line := r.FindString(str)

	colname = ut.PixContent(line, "`")
	return colname
}

//获取某一个table的primarykey
func (ml MysqlLexer) Primarykey(tableStr string) string {
	r := regexp.MustCompile(primaryKeyLine)
	pline := r.FindString(tableStr)
	pk := ut.PixContent(pline, "`")
	return pk
}

//获取sql脚本
func (ml MysqlLexer) SqlString(file string) string {
	var err error
	dat, err = ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	return string(dat)
}

//获取createTable的字符串数组
func (ml MysqlLexer) CreateTableString(sqlStr string) []string {

	r := regexp.MustCompile(createTable)
	str := r.FindAllString(sqlStr, -1)
	return str
}

func (ml MysqlLexer) TableName(tableStr string) string {

	r := regexp.MustCompile(tableNameLine)
	tname := r.FindString(tableStr)
	tname = ut.PixContent(tname, "`")
	return tname
}
