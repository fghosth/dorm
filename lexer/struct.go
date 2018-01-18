package lexer

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	//所有struct体字符串
	structStr = `type[\w\W]+?struct[\w\W]+?}`
	//type定义struct到{位子的字符串
	typeStructStr = `type.+{`
	//所有成员变量行
	structFieldStr = "[a-zA-Z].+"
	//所有go变量类型
	keywordStr = `\b(int|int8|int16|int32|int64|string|float32|float64|\[\]byte|bool)\b`
	//所有导出类单词
	oupPutWord = `[A-Z][a-zA-z\\b]+`
	//所有tag
	tagStr = "`.+`"
)

type StructLexer struct{}

//解析tag
func (sl *StructLexer) Taglex(tag string) (tagmap map[string]string) {
	tagmap = make(map[string]string)
	tag = strings.Replace(tag, "`", "", -1)
	leng := len(strings.Split(tag, " "))

	for i := 0; i < leng; i++ {
		r := regexp.MustCompile(`(.+?:".+?")`) //匹配类似dormCol:"secret_key"
		str := r.FindString(tag)
		// fmt.Println(str)
		arr := strings.Split(str, ":")
		if len(arr) == 2 {
			tagmap[strings.Replace(arr[0], " ", "", -1)] = strings.Replace(arr[1], "\"", "", -1)
		}
		tag = strings.Replace(tag+" ", str, "", -1)
	}

	return tagmap
}

//获取所有导出成员变量
func (sl *StructLexer) FieldName(structStr string) []map[string]string {
	field := make([]map[string]string, 0)
	r := regexp.MustCompile(structFieldStr)
	line := r.FindAllString(structStr, -1)
	for _, v := range line {
		isCol, err := regexp.MatchString(keywordStr, v)
		if err != nil {
			fmt.Println(err)
		}
		if isCol {
			//处理每个字符串去除制表符tab等。把多个空格换成一个,去除左右2边所有空格
			l := strings.Replace(v, "\t", "", -1)
			l = strings.TrimRight(l, " ")
			l = strings.TrimLeft(l, " ")
			r = regexp.MustCompile("[ ]+")
			l = r.ReplaceAllString(l, " ")
			colmap := make(map[string]string)
			colmap["field"], colmap["type"] = getStructField(l)
			colmap["tag"] = getStructTag(l)
			field = append(field, colmap)
		}

	}
	return field
}

//获取单行tag
func getStructTag(str string) string {
	tag := ""
	r := regexp.MustCompile(tagStr)
	tag = r.FindString(str)
	return tag
}

//获取单行成员名,类型
func getStructField(str string) (fieldName string, fieldType string) {
	arr := strings.Split(str, " ")
	if len(arr) < 2 {
		fmt.Println("传入的数据错误：" + str)
	} else {
		fieldName = arr[0]
		fieldType = arr[1]
	}
	return fieldName, fieldType
}

//获取所有struct字符串数组
func (sl *StructLexer) StructStr(fileStr string) []string {
	r := regexp.MustCompile(structStr)
	sline := r.FindAllString(fileStr, -1)
	return sline
}

//获取struct名字
func (sl *StructLexer) StructName(sStr string) string {
	r := regexp.MustCompile(typeStructStr)
	sline := r.FindString(sStr)
	r = regexp.MustCompile(oupPutWord)
	sname := r.FindString(sline)
	return sname
}

//读取文件
func (ml *StructLexer) GetStructFile(file string) string {
	var err error
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	return string(dat)
}
