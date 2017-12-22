package lexer

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"

	"github.com/aymerick/raymond"
)

var (
	//匹配所有mysql变量 匹配完整单词 解决 同时出现datetime 和date的问题
	cockDBsqlType = `(?i)\b(INT|INTEGER|INT8|INT64|BIGINT|INT4|INT2|SMALLINT|BIT|UUID|SERIAL|SMALLSERIAL|BIGSERIAL|DECIMAL|DEC|NUMERIC|FLOAT|DOUBLE PRECISION|FLOAT8|BOOLEAN|BOOL|DATE|TIMESTAMPTZ|TIMESTAMP|INTERVAL|STRING|CHARACTER|CHAR|VARCHAR|TEXT|COLLATE|BYTES|BYTEA|BLOB|ARRAY)\b\(.+?\)|\b(INT|INTEGER|INT8|INT64|BIGINT|INT4|INT2|SMALLINT|BIT|UUID|SERIAL|SMALLSERIAL|BIGSERIAL|DECIMAL|DEC|NUMERIC|FLOAT|DOUBLE PRECISION|FLOAT8|BOOLEAN|BOOL|DATE|TIMESTAMPTZ|TIMESTAMP|INTERVAL|STRING|CHARACTER|CHAR|VARCHAR|TEXT|COLLATE|BYTES|BYTEA|BLOB|ARRAY)\b`
	// tableName = "(?<=TABLE[\\s]{1,200}`).{1,}(?=`)"
	//获取tablename所在行(?i)忽略大小写
	cockDBtableNameLine = "(?i)CREATE TABLE.+\\("

	//匹配用户自定义的表名，字段名等-----以小写字幕开头，包含大小写-_和数字的字符串
	cockDBName = "[a-z][a-zA-Z\\d_-]+"
	//选中非字段行
	cockDBNotField = `^\b(FAMILY|NOT NULL|DEFAULT|PRIMARY KEY|UNIQUE|CHECK|CONSTRAINT|INDEX|IF NOT EXISTS|SHOW|FROM|CREATE TABLE|DELETE|EXPLAIN|IMPORT|INSERT|SELECT|SHOW TRACE|TRUNCATE|UPDATE|UPSERT)`
	//保留关键字
	cockDBKeyword = `(?i)\b(FAMILY|NOT NULL|DEFAULT|PRIMARY KEY|UNIQUE|CHECK|CONSTRAINT|INDEX|IF NOT EXISTS|SHOW|FROM|CREATE TABLE|DELETE|EXPLAIN|IMPORT|INSERT|SELECT|SHOW TRACE|TRUNCATE|UPDATE|UPSERT)`
	//找出所以cerate table代码段
	cockDBcreateTable = `(?i)(CREATE TABLE)[\W\w]+?;`
	//为创造table的语句按字段分行
	cockDBcolLine = `[a-zA-Z].+,\n`
	//找到index 行TODO
	cockDBindexLine = `(?i)INDEX.+`
)

var (
	StructToCockDBMap = map[string]string{
		"int":     "INT",
		"int8":    "BIT",
		"int16":   "INT2",
		"int32":   "INT",
		"int64":   "BIGINT",
		"string":  "STRING",
		"float32": "FLOAT",
		"float64": "FLOAT",
		"[]byte":  "BYTES",
		"bool":    "BOOL",
	}
	CockdbToStructMap = map[string]string{
		"INT":              "int64",
		"INTEGER":          "int64",
		"INT8":             "int64",
		"INT64":            "int64",
		"BIGINT":           "int64",
		"INT4":             "int32",
		"INT2":             "int16",
		"BIT":              "int8",
		"UUID":             "string",
		"SERIAL":           "int64",
		"SMALLSERIAL":      "int64",
		"BIGSERIAL":        "int64",
		"DECIMAL":          "float64",
		"DEC":              "float64",
		"NUMERIC":          "float64",
		"FLOAT":            "float64",
		"REAL":             "float64",
		"FLOAT4":           "float64",
		"DOUBLE PRECISION": "float64",
		"FLOAT8":           "float64",
		"BOOLEAN":          "bool",
		"BOOL":             "bool",
		"DATE":             "string",
		"TIMESTAMP":        "int",
		"TIMESTAMPTZ":      "int",
		"INTERVAL":         "int",
		"TEXT":             "string",
		"VARCHAR":          "string",
		"CHAR":             "string",
		"CHARACTER":        "string",
		"STRING":           "string",
		"COLLATE":          "string",
		"BLOB":             "[]byte",
		"BYTEA":            "[]byte",
		"BYTES":            "[]byte",
		"ARRAY":            "string",
	}
)

type CockDBLexer struct {
}

//根据struct字符串生成数据库sql
func (ml CockDBLexer) CreateSqlByStructStr(strStruct string) string {
	sl := new(StructLexer)
	var tableName, sql string
	flist := make([]string, 0)
	str := sl.FieldName(strStruct)
	tableName = sl.StructName(strStruct) //表明
	for k, v := range str {              //遍历struct字段
		tag := sl.Taglex(v["tag"])
		colName := tag["dormCol"] //字段名
		if colName == "" {        //tag中没有就用字段名
			colName = ut.CalToUnder(v["field"])
		}
		colType := tag["dormCOCKDBType"]
		if colType == "" { //如果 没有tag则使用默认匹配map
			colType = StructToMysqlMap[v["type"]] //字段类型
		}
		colProperty := tag["dorm"] //字段属性
		tpms := colName + " " + colType + " " + colProperty

		if k < len(str)-1 { //最后一句不加逗号，
			tpms = tpms + ","
		}

		flist = append(flist, tpms)

	}
	tableName = ut.CalToUnder(tableName)
	//根据模板生成
	ctx := map[string]interface{}{
		"tableName": tableName,
		"field":     flist,
	}

	sql, err := raymond.Render(COCKROACH_SCRIPT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return sql
}

//根据struct-go反射生成数据库sql
func (ml CockDBLexer) CreateSqlByStruct(obj interface{}) string {
	var tableName, sql string
	flist := make([]string, 0)
	rtype := reflect.TypeOf(obj).Elem()
	tableName = rtype.Name()                //表明
	for k := 0; k < rtype.NumField(); k++ { //遍历struct字段
		colName := rtype.Field(k).Tag.Get("dormCol") //字段名
		if colName == "" {                           //tag中没有就用字段名
			colName = ut.CalToUnder(rtype.Field(k).Name)
		}
		colType := rtype.Field(k).Tag.Get("dormCOCKDBType")
		if colType == "" { //如果 没有tag则使用默认匹配map
			colType = StructToMysqlMap[string(rtype.Field(k).Type.Kind().String())] //字段类型
		}
		colProperty := rtype.Field(k).Tag.Get("dorm") //字段属性
		tpms := colName + " " + colType + " " + colProperty

		if k < rtype.NumField()-1 { //最后一句不加逗号，
			tpms = tpms + ","
		}

		flist = append(flist, tpms)

	}
	tableName = ut.CalToUnder(tableName)
	//根据模板生成
	ctx := map[string]interface{}{
		"tableName": tableName,
		"field":     flist,
	}

	sql, err := raymond.Render(COCKROACH_SCRIPT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return sql
}

//创建struct字符串
func (ml CockDBLexer) CreateStruct(packageName, tableName string, field []map[string]interface{}) string {
	var structString string
	fstr := make([]string, 0)
	for _, v := range field { //遍历字段

		str := ut.UnderToCal(v["colName"].(string)) + " " + v["goType"].(string) + "  `dormCol:\"" + v["colName"].(string) + "\" " + "dormCOCKDBType:\"" + v["sqltype"].(string) + "\"" + " dorm:\"" + v["property"].(string) + "\"`"
		fstr = append(fstr, str)
	}
	ctx := map[string]interface{}{
		"name":        ut.UnderToCal(tableName),
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
func (ml CockDBLexer) Field(tableStr string) []map[string]interface{} {
	field := make([]map[string]interface{}, 0)
	r := regexp.MustCompile(cockDBcolLine)
	line := r.FindAllString(tableStr, -1)
	for _, v := range line {
		coln := getCockColnameByLine(v) //获取字段名

		if coln != "" { //字段名不能为空 不是字段
			colmap := make(map[string]interface{})
			colmap["colName"] = coln
			colmap["sqltype"], colmap["goType"] = getCockColTypeByLine(v)
			colmap["property"] = getCockColptyByLine(v)
			field = append(field, colmap)

		}
	}
	// fmt.Println(field)
	return field
}

//根据行获取熟悉数组
func getCockColptyByLine(str string) string {
	pty := ""
	r := regexp.MustCompile(cockDBsqlType)
	str = strings.Replace(str, ",", "", -1)
	tmp := r.FindString(str)
	if tmp != "" {
		ptyList := strings.Split(str, " "+tmp)
		if len(ptyList) > 1 {
			pty = ptyList[1]
		}
	}
	// pp.Println(str, "+++++", tmp, "====", pty)
	pty = strings.Replace(pty, "\n", "", -1)

	return strings.TrimLeft(pty, " ")
}

//根据行取出字段类型
func getCockColTypeByLine(str string) (sqltype, gotype string) {

	r := regexp.MustCompile(cockDBsqlType)
	sqltype = r.FindString(str)
	r = regexp.MustCompile(`\(.+\)`)
	tmap := r.ReplaceAllString(sqltype, "")
	gotype = CockdbToStructMap[tmap]
	// pp.Println(gotype)
	return
}

//根据行取出字段名 如果不是字段行返回""
func getCockColnameByLine(str string) string {
	colname := ""

	isCol, err := regexp.MatchString(cockDBsqlType, str)     //是否包含字段类型
	notField, err := regexp.MatchString(cockDBNotField, str) //是否包含保留关键字开头
	if err != nil {
		fmt.Println(err)
	}
	if !isCol || notField {
		return colname
	}

	r := regexp.MustCompile(cockDBName)

	colname = r.FindString(str)

	return colname
}

//获得tableName
func (ml CockDBLexer) TableName(tableStr string) string {

	r := regexp.MustCompile(cockDBtableNameLine)
	tname := r.FindString(tableStr)
	r = regexp.MustCompile(cockDBName)
	tname = r.FindString(tname)
	return tname
}

//获取createTable的字符串数组
func (ml CockDBLexer) CreateTableString(sqlStr string) []string {

	r := regexp.MustCompile(cockDBcreateTable)
	str := r.FindAllString(sqlStr, -1)
	return str
}

//获取sql脚本
func (ml CockDBLexer) SqlString(file string) string {
	var err error
	dat, err = ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(dat)
}
