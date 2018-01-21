package main

import (
	_ "database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

var (
	cocksqlFileName = "./mysql2Cockroach_gen.sql"
)

/*sadfa
 */
func cockroachsqlFromMysql(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}

	sqlfile := c.String("file")

	if sqlfile == "" {
		fmt.Println("您未指定sql文件")
		return ERRNOFILE
	}
	return genCockroachSqlFromMysql(sqlfile)

}

//根据cockroach脚本生成struct
func genCockroachSqlFromMysql(sqlfile string) error {
	mysqlLexer := new(lexer.MysqlLexer)
	sqlStr := mysqlLexer.SqlString(sqlfile)
	str := mysqlLexer.CreateTableString(sqlStr)
	var createTableStr string
	for _, v := range str {
		createTableStr = createTableStr + "\n" + mysqlLexer.CreateCockroachSqlFromMysql(v)
	}
	insertStr := "\n" + mysqlLexer.CreateCockInsertSqlFromMysql(mysqlLexer.InsertStr(sqlStr))
	//生成文件
	exist, err := ut.FileOrPathExists(cocksqlFileName)
	if !COVRE && exist {
		fmt.Println(util.Red(cocksqlFileName + "文件已存在"))
		// return err
	}
	sf, err := os.Create(cocksqlFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(createTableStr + insertStr))
	if err != nil {
		fmt.Println(err)
		return err
	}
	// pp.Println(err)
	if !(!COVRE && exist) && err == nil {
		fmt.Println(cocksqlFileName + "生成成功")
	}
	return nil
}
