package main

import (
	_ "database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

var (
	ERRNOFILE   = errors.New("未指定sql文件：-f ./user.sql")
	ERRNOSQL    = errors.New("未指定数据库类型:-d mysql(cockroach)")
	structPath  = "./ormstruct/"
	packageName = "ormstruct"
)

/*sadfa
 */
func CstructFile(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	db := c.String("database")
	sqlfile := c.String("file")
	if db != "mysql" && db != "cockroach" && db != "mariadb" && db != "postgresql" {
		fmt.Println("目前只支持mysql,mariadb,postgresql,cockroach")
		return ERRNOSQL
	}
	if sqlfile == "" {
		fmt.Println("您未指定sql文件")
		return ERRNOFILE
	}
	if db == "mysql" || db == "mariadb" {
		return genStructByMysql(sqlfile)
	}
	if db == "cockroach" || db == "postgresql" {
		return genStructByCockroach(sqlfile)
	}
	return nil
}

//根据cockroach脚本生成struct
func genStructByCockroach(sqlfile string) error {
	cockDBlexer := new(lexer.CockDBLexer)
	cocksqlStr := cockDBlexer.SqlString(sqlfile)

	str := cockDBlexer.CreateTableString(cocksqlStr)
	os.Mkdir(structPath, os.ModePerm) //在当前目录下生成md目录

	for _, v := range str {
		tname := cockDBlexer.TableName(v)
		field := cockDBlexer.Field(v)
		structStr := cockDBlexer.CreateStruct(packageName, tname, field)
		//生成文件
		fileName := tname + ".go"
		exist, err := ut.FileOrPathExists(structPath + fileName)
		if !COVRE && exist {
			fmt.Println(util.Red(structPath + fileName + "文件已存在"))
			// return err
		}
		sf, err := os.Create(structPath + fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = sf.Write([]byte(structStr))
		if err != nil {
			fmt.Println(err)
			return err
		}
		// pp.Println(err)
		if !(!COVRE && exist) && err == nil {
			fmt.Println(structPath + fileName + "生成成功")
		}
	}
	return nil
}

//根据mysql脚本生成struct
func genStructByMysql(sqlfile string) error {
	mysqlLexer := new(lexer.MysqlLexer)
	sqlStr := mysqlLexer.SqlString(sqlfile)
	str := mysqlLexer.CreateTableString(sqlStr)
	os.Mkdir(structPath, os.ModePerm) //在当前目录下生成md目录

	for _, v := range str {
		tname := mysqlLexer.TableName(v)
		field := mysqlLexer.Field(v)
		structStr := mysqlLexer.CreateStruct(packageName, tname, field)
		//生成文件
		fileName := tname + ".go"
		exist, err := ut.FileOrPathExists(structPath + fileName)
		if !COVRE && exist {
			fmt.Println(util.Red(structPath + fileName + "文件已存在"))
			// return err
		}
		sf, err := os.Create(structPath + fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = sf.Write([]byte(structStr))
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !(!COVRE && exist) && err == nil {
			fmt.Println(structPath + fileName + "生成成功")
		}
	}
	return nil
}
