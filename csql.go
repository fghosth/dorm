package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

var (
	sqlFileName = "auto_gen.sql"
)

func CsqlFile(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	db := c.String("database")
	file := c.String("file")
	//处理file如果是目录的情况。确保目录后有『/』
	hasf, _ := regexp.MatchString(`\/$`, file)
	if !hasf {
		file = file + "/"
	}
	if db != "mysql" && db != "cockroach" {
		fmt.Println("目前只支持mysql和cockroach")
		return ERRNOFILE
	}
	if file == "" {
		fmt.Println("您未指定struct文件")
		return ERRNOSQL
	}
	if db == "mysql" {
		return genMysqlByStruct(file)

	}
	if db == "cockroach" {
		return genCockroachSqlByStruct(file)
	}
	return nil
}

//生成mysql的脚本
func genMysqlByStruct(file string) error {
	exist, err := ut.FileOrPathExists(sqlFileName)
	str := ""
	if !COVRE && exist {
		fmt.Println(util.Red(sqlFileName + "文件已存在"))
		return err
	}
	f, _ := os.Stat(file)
	if f.IsDir() { //如果是目录
		str = getAllMysqlString(file)
	} else { //如果是单文件
		mysqlLexer := new(lexer.MysqlLexer)
		sl := new(lexer.StructLexer)
		fileStr := sl.GetStructFile(file)
		slist := sl.StructStr(fileStr)
		for _, v := range slist {
			str = str + "\n" + mysqlLexer.CreateSqlByStructStr(v)
		}
	}
	sf, err := os.Create(sqlFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !(!COVRE && exist) && err == nil {
		fmt.Println(sqlFileName + "生成成功")
	}
	return nil
}

//生成cockroach脚本
func genCockroachSqlByStruct(file string) error {
	exist, err := ut.FileOrPathExists(sqlFileName)
	str := ""
	if !COVRE && exist {
		fmt.Println(util.Red(sqlFileName + "文件已存在"))
		return err
	}
	f, _ := os.Stat(file)
	if f.IsDir() { //如果是目录
		str = getAllCockDBString(file)
	} else { //如果是单文件
		cockDBLexer := new(lexer.CockDBLexer)
		sl := new(lexer.StructLexer)
		fileStr := sl.GetStructFile(file)
		slist := sl.StructStr(fileStr)
		for _, v := range slist {
			str = str + "\n" + cockDBLexer.CreateSqlByStructStr(v)
		}
	}
	sf, err := os.Create(sqlFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !(!COVRE && exist) && err == nil {
		fmt.Println(sqlFileName + "生成成功")
	}
	return nil
}

//cockroach用遍历目录拼接所有文件并去除生成sql文件string
func getAllCockDBString(fp string) (allStr string) {
	err := filepath.Walk(fp, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			cockDBLexer := new(lexer.CockDBLexer)
			sl := new(lexer.StructLexer)
			fileStr := sl.GetStructFile(fp + f.Name())
			slist := sl.StructStr(fileStr)
			str := ""
			for _, v := range slist {
				str = str + "\n" + cockDBLexer.CreateSqlByStructStr(v)
			}
			allStr = allStr + "\n" + str
			// fmt.Println(str)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return allStr
}

//mysqlsql用遍历目录拼接所有文件并去除生成sql文件string
func getAllMysqlString(fp string) (allStr string) {
	err := filepath.Walk(fp, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			mysqlLexer := new(lexer.MysqlLexer)
			sl := new(lexer.StructLexer)
			fileStr := sl.GetStructFile(fp + f.Name())
			slist := sl.StructStr(fileStr)
			str := ""
			for _, v := range slist {
				str = str + "\n" + mysqlLexer.CreateSqlByStructStr(v)
			}
			allStr = allStr + "\n" + str
			// fmt.Println(str)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return allStr
}
