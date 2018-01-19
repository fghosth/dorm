package main

import (
	_ "database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"jvole.com/createProject/dorm"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
)

var (
	M_ERRNOFILE   = errors.New("未指定struct文件：-f ./user.sql")
	M_ERRNOSQL    = errors.New("未指定数据库类型:-d mysql(cockroach)")
	m_modelPath   = "./model/"
	m_packageName = "model"
)

/*sadfa
 */
func CmodelFile(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	db := c.String("database")
	file := c.String("file")
	if db != "mysql" && db != "cockroach" {
		fmt.Println("目前只支持mysql和cockroach")
		return M_ERRNOSQL
	}
	if file == "" {
		fmt.Println("您未指定struct文件")
		return M_ERRNOFILE
	}
	createBaseModel()
	f, _ := os.Stat(file)
	if f.IsDir() { //如果是目录
		err := filepath.Walk(file, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if !f.IsDir() {
				createModel(path, db)
			}
			return nil
		})
		return err
	} else {
		return createModel(file, db)
	}

}

func createBaseModel() {
	Str := dorm.CreateModel(m_packageName)
	os.Mkdir(structPath, os.ModePerm) //在当前目录下生成md目录
	//生成文件
	fileName := "model.go"
	exist, err := ut.FileOrPathExists(m_modelPath + fileName)
	if !COVRE && exist {
		fmt.Println(util.Red(m_modelPath + fileName + "文件已存在"))
		// return err
	}
	sf, err := os.Create(m_modelPath + fileName)
	if err != nil {
		fmt.Println(err)
		// return err
	}
	_, err = sf.Write([]byte(Str))
	if err != nil {
		fmt.Println(err)
		// return err
	}
	// pp.Println(err)
	if !(!COVRE && exist) && err == nil {
		fmt.Println(m_modelPath + fileName + "生成成功")
	}

}

//根据cockroach脚本生成struct
func createModel(file, db string) error {
	sl := new(lexer.StructLexer)
	fileStr := sl.GetStructFile(file)
	arrStruct := sl.StructStr(fileStr)
	os.Mkdir(structPath, os.ModePerm) //在当前目录下生成md目录

	for _, v := range arrStruct {
		Str := dorm.CreateDorm(m_packageName, db, v)
		tname := ut.CalToUnder(sl.StructName(v))
		//生成文件
		fileName := tname + ".go"
		exist, err := ut.FileOrPathExists(m_modelPath + fileName)
		if !COVRE && exist {
			fmt.Println(util.Red(m_modelPath + fileName + "文件已存在"))
			// return err
		}
		sf, err := os.Create(m_modelPath + fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = sf.Write([]byte(Str))
		if err != nil {
			fmt.Println(err)
			return err
		}
		// pp.Println(err)
		if !(!COVRE && exist) && err == nil {
			fmt.Println(m_modelPath + fileName + "生成成功")
		}
	}
	return nil
}
