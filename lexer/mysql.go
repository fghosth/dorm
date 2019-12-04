package lexer

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/aymerick/raymond"
	"jvole.com/createProject/util"
)

var (
	//匹配所有mysql变量 匹配完整单词 解决 同时出现datetime 和date的问题
	sqlType = `(?i)\b(bool|tinyint|smallint|mediumint|int|integer|bigint|date|datetime|time|bit|tinytext|mediumtext|longtext|text|tinyblob|mediumblob|longblob|blob|float|double|decimal|timestamp|year|char|varchar|varbinary|enum|set|json)\b\(.+?\)|\b(bool|tinyint|smallint|mediumint|int|integer|bigint|date|datetime|time|bit|tinytext|mediumtext|longtext|text|tinyblob|mediumblob|longblob|blob|float|double|decimal|timestamp|year|char|varchar|varbinary|enum|set|json)\b`
	// tableName = "(?<=TABLE[\\s]{1,200}`).{1,}(?=`)"
	//获取tablename所在行(?i)忽略大小写
	tableNameLine = "(?i)(CREATE TABLE).+\\`"
	//匹配所有字段
	col = "`.+?`"
	//匹配字段属性
	property = `(?i)\b(NOT NULL|(DEFAULT.+)|AUTO_INCREMENT|unsigned|zerofill|PRIMARY.+,)`
	//找出所以cerate table代码段
	createTable = `(?i)(CREATE TABLE)[\W\w]+?;`
	//为创造table的语句按字段分行
	colLine = `.+,`
	//找到PRIMARY KEY行
	primaryKeyLine = `(?i)(PRIMARY KEY).+`
	//找到UNIQUE KEY的行
	uniqueKeyLine = `(?i)(UNIQUE KEY).+`
	//找到index 行TODO
	indexLine = `(?i)(KEY).+`
	//在(`开始`)结束之间的内容
	sContent = "\\(\\`[\\W\\w]+?\\`\\)"
	//找到所有索引行KEY
	keyIndex = `([^a-zA-Z\n]+[ ]+KEY).+`
	//找到comment
	comment = `(?i)\b(COMMENT.+')`
	//查到所有insert语句
	insertSql = `(?i)(insert into)[\w\W]+?;`
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
		"bool":    "bit",
	}
	MysqlToStructMap = map[string]string{
		"tinyint":    "int8",
		"smallint":   "int",
		"mediumint":  "int",
		"int":        "int",
		"integer":    "int",
		"bigint":     "int",
		"date":       "string",
		"datetime":   "string",
		"time":       "string",
		"bit":        "int8",
		"bool":       "bool",
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
		"timestamp":  "time.Time",
		"year":       "string",
		"char":       "string",
		"varchar":    "string",
		"varbinary":  "[]byte",
		"enum":       "string",
		"set":        "string",
		"json":       "string",
	}
	MysqlToCockroach = map[string]string{
		"tinyint":    "SMALLINT",
		"smallint":   "SMALLINT",
		"mediumint":  "INT4",
		"int":        "INT",
		"integer":    "INT",
		"bigint":     "INT",
		"date":       "DATE",
		"datetime":   "TIMESTAMP",
		"time":       "STRING",
		"bit":        "SMALLINT",
		"bool":       "BOOL",
		"tinytext":   "STRING",
		"mediumtext": "STRING",
		"longtext":   "STRING",
		"text":       "STRING",
		"tinyblob":   "BYTES",
		"mediumblob": "BYTES",
		"longblob":   "BYTES",
		"blob":       "BYTES",
		"float":      "FLOAT",
		"double":     "REAL",
		"decimal":    "DECIMAL",
		"timestamp":  "TIMESTAMP",
		"year":       "STRING",
		"char":       "STRING",
		"varchar":    "STRING",
		"varbinary":  "BYTES",
		"enum":       "STRING",
		"set":        "STRING",
		"json":       "STRING",
	}
)

type MysqlLexer struct {
}

var (
	ut        = new(util.Dstring)
	dat       []byte //加载的文件
	splitFlag = " "  //在tag中标识类似commit-'用户id'的分隔符
)

func (ml MysqlLexer) CreateCockInsertSqlFromMysql(insertStr []string) string {

	//根据模板生成
	ctx := map[string]interface{}{
		"field": insertStr,
	}

	cocksql, err := raymond.Render(COCKROACH_INSERT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return cocksql
}

func (ml MysqlLexer) InsertStr(sqlStr string) []string {
	sqlStr = strings.Replace(sqlStr, "`", `"`, -1)
	r := regexp.MustCompile(insertSql)
	str := r.FindAllString(sqlStr, -1)
	return str
}

//根据mysql脚本生成cockroachDB脚本(创建表),支持primaryKey Index Unique DEFAULT值（其中index不支持联合索引--会拆分为多个索引）
func (ml MysqlLexer) CreateCockroachSqlFromMysql(tableStr string) string {
	tableName := ml.TableName(tableStr)

	field := ml.Field(tableStr)
	flist := make([]string, len(field))

	pk := ml.Primarykey(tableStr)
	indexk := ml.IndexKey(tableStr)
	uq := ml.Uniquekey(tableStr)
	for i, v := range field {
		if i == 0 && ut.ContainStrBysplit(pk, v["colName"].(string), ",") {

			flist[i] = v["colName"].(string) + " SERIAL" + getCockProperty(v["cockDBType"].(string), v["property"].([]string))

		} else {
			flist[i] = v["colName"].(string) + " " + v["cockDBType"].(string) + getCockProperty(v["cockDBType"].(string), v["property"].([]string))
		}
		if i < len(field)-1 {
			flist[i] = flist[i] + ","
		} else {
			if pk != "" || indexk != "" || uq != "" {
				flist[i] = flist[i] + ","
			}
		}
	}
	if pk != "" { //PRIMARY KEY
		pk = "PRIMARY KEY (" + pk + ")"
		if uq != "" || indexk != "" {
			pk = pk + ","
		}
	}
	if uq != "" { //UNIQUE
		uq = "UNIQUE (" + uq + ")"
		if indexk != "" {
			uq = uq + ","
		}
	}
	var ik []string
	if indexk != "" { //INDEX
		arrik := strings.Split(indexk, ",")
		ik = make([]string, len(arrik))
		for i, v := range arrik {
			ik[i] = "INDEX " + tableName + "_index_" + strconv.Itoa(i) + " (" + v + ")"
			if i < len(arrik)-1 {
				ik[i] = ik[i] + ","
			}
		}

	}
	//根据模板生成
	ctx := map[string]interface{}{
		"tableName": tableName,
		"field":     flist,
		"pk":        pk,
		"ik":        ik,
		"uq":        uq,
	}

	cocksql, err := raymond.Render(COCKROACH_SCRIPT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return cocksql
}

func getCockProperty(ftype string, dormStr []string) (pstr string) {
	for _, v := range dormStr {
		//处理comment
		r := regexp.MustCompile(comment)
		commstr := r.FindString(v)
		r = regexp.MustCompile(comment)
		v := r.ReplaceAllString(v, "")
		v = strings.TrimSpace(v)
		if v == "NOT NULL" {
			pstr = pstr + " NOT NULL"
		}
		if commstr != "" {
			// pp.Println(commstr)
		}
		if strings.Contains(v, "DEFAULT") {
			var value string //default 的默认值
			arrvalue := strings.Split(v, " ")
			if len(arrvalue) > 1 {
				value = strings.Replace(arrvalue[1], "'", "", -1) //去除'
			}
			switch value {
			case "CURRENT_TIMESTAMP":
				pstr = pstr + " DEFAULT current_timestamp()"
			case "NULL":
				pstr = pstr + " DEFAULT NULL"
			default:
				pstr = pstr + " DEFAULT '" + value + "':::" + ftype
			}
		}

	}
	return
}

//根据struct-字符串成数据库sql
func (ml MysqlLexer) CreateSqlByStructStr(strStruct string) string {
	sl := new(StructLexer)
	var tableName, sql, primaryKey string
	flist := make([]string, 0)

	tableName = sl.StructName(strStruct) //表明
	str := sl.FieldName(strStruct)
	for k, v := range str { //遍历struct字段
		tag := sl.Taglex(v["tag"])
		primaryPOS := strings.Index(tag["dorm"], "PRIMARY;")
		if primaryPOS != -1 {
			primaryKey = "PRIMARY KEY (`" + tag["dormCol"] + "`)" //primaryKey
		}
		colName := tag["dormCol"] //字段名
		if colName == "" {        //tag中没有就用字段名
			colName = ut.CalToUnder(v["field"])
		}
		colType := tag["dormMysqlType"]
		if colType == "" { //如果 没有tag则使用默认匹配map
			colType = StructToMysqlMap[v["type"]] //字段类型
		}
		colProperty := dormToSql(tag["dorm"]) //字段属性
		tpms := "`" + colName + "` " + colType + " " + colProperty

		if k < len(str)-1 || primaryKey != "" { //最后一句不加逗号，
			tpms = tpms + ","
		}

		flist = append(flist, tpms)

	}
	tableName = ut.CalToUnder(tableName)
	//根据模板生成
	ctx := map[string]interface{}{
		"tableName":  tableName,
		"field":      flist,
		"primaryKey": primaryKey,
	}

	sql, err := raymond.Render(MYSQL_SCRIPT_TMP, ctx)
	if err != nil {
		fmt.Println(err)
	}
	return sql
}

//根据struct-go反射生成数据库sql
func (ml MysqlLexer) CreateSqlByStruct(obj interface{}) string {
	var tableName, sql, primaryKey string
	flist := make([]string, 0)
	rtype := reflect.TypeOf(obj).Elem()
	tableName = rtype.Name() //表明

	for k := 0; k < rtype.NumField(); k++ { //遍历struct字段
		primaryPOS := strings.Index(rtype.Field(k).Tag.Get("dorm"), "PRIMARY;")
		if primaryPOS != -1 {
			primaryKey = "PRIMARY KEY (`" + rtype.Field(k).Tag.Get("dormCol") + "`)" //primaryKey
		}
		colName := rtype.Field(k).Tag.Get("dormCol") //字段名
		if colName == "" {                           //tag中没有就用字段名
			colName = ut.CalToUnder(rtype.Field(k).Name)
		}
		colType := rtype.Field(k).Tag.Get("dormMysqlType")
		if colType == "" { //如果 没有tag则使用默认匹配map
			colType = StructToMysqlMap[string(rtype.Field(k).Type.Kind().String())] //字段类型
		}
		colProperty := dormToSql(rtype.Field(k).Tag.Get("dorm")) //字段属性
		tpms := "`" + colName + "` " + colType + " " + colProperty

		if k < rtype.NumField()-1 || primaryKey != "" { //最后一句不加逗号，
			tpms = tpms + ","
		}

		flist = append(flist, tpms)

	}
	tableName = ut.CalToUnder(tableName)
	//根据模板生成
	ctx := map[string]interface{}{
		"tableName":  tableName,
		"field":      flist,
		"primaryKey": primaryKey,
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
	sql = strings.Replace(sql, splitFlag, " ", -1)
	sql = strings.Replace(sql, "PRIMARY", " ", -1)
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
		str := ut.UnderToCal(v["colName"].(string)) + " " + v["goType"].(string) + "  `dormCol:\"" + v["colName"].(string) + "\" " + "dormMysqlType:\"" + v["sqltype"].(string) + "\"" + " dorm:\"" + tmpstr + "\"`"
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

//获取字段及字段属性map
func (ml MysqlLexer) Field(tableStr string) []map[string]interface{} {
	field := make([]map[string]interface{}, 0)
	r := regexp.MustCompile(colLine)
	line := r.FindAllString(tableStr, -1)
	mysqlLexer := new(MysqlLexer)
	pk := mysqlLexer.Primarykey(tableStr)
	for _, v := range line {
		coln := getColnameByLine(v) //获取字段名

		if coln != "" { //字段名不能为空 不是字段
			colmap := make(map[string]interface{})
			colmap["colName"] = coln
			colmap["sqltype"], colmap["goType"], colmap["cockDBType"] = getColTypeByLine(v)
			colmap["property"] = getColptyByLine(v, pk)
			field = append(field, colmap)

		}
	}
	// fmt.Println(field)
	return field
}

//根据行获取熟悉数组
func getColptyByLine(str, pk string) []string {
	pty := ""
	r := regexp.MustCompile(property)
	ptylist := r.FindAllString(str, -1)
	// pp.Println(ut.ContainStrBysplit(pk, getColnameByLine(str), ","))
	if ut.ContainStrBysplit(pk, getColnameByLine(str), ",") {
		pty = "PRIMARY"
	}
	for _, v := range ptylist {

		if pty == "" {
			pty = strings.Replace(v, " ", splitFlag, -1)
		} else {
			pty = pty + "||" + strings.Replace(v, " ", splitFlag, -1)
		}
	}

	pty = strings.Replace(pty, ",", "", -1)
	return strings.Split(pty, "||")
}

//根据行取出字段类型
func getColTypeByLine(str string) (sqltype, gotype, cockDBtype string) {

	r := regexp.MustCompile(sqlType)
	sqltype = r.FindString(str)
	r = regexp.MustCompile(`\(.+\)`)
	tmap := r.ReplaceAllString(sqltype, "")
	gotype = MysqlToStructMap[tmap]
	cockDBtype = MysqlToCockroach[tmap]
	// pp.Println(gotype)
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

//获取某一个table的Uniquekey
func (ml MysqlLexer) Uniquekey(tableStr string) string {
	var uq string
	r := regexp.MustCompile(uniqueKeyLine)
	pline := r.FindAllString(tableStr, -1)

	for _, v := range pline {
		r = regexp.MustCompile(sContent)
		if uq == "" {
			uq = ut.PixContent(r.FindString(v), "`")
		} else {
			uq = uq + "," + ut.PixContent(r.FindString(v), "`")
		}
	}
	return uq
}

//获取某一个table的Indexkey
func (ml MysqlLexer) IndexKey(tableStr string) string {
	var index string
	r := regexp.MustCompile(keyIndex)
	pline := r.FindAllString(tableStr, -1)

	for _, v := range pline {
		r = regexp.MustCompile(sContent)
		if index == "" {
			index = ut.PixContent(r.FindString(v), "`")
		} else {
			index = index + "," + ut.PixContent(r.FindString(v), "`")
		}
	}
	return index
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
