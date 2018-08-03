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
	M_ERRLOCATION = errors.New("location不能为空:-l jvole.com/book/")
	m_modelPath   = "./model/"
	m_packageName = "model"
	m_base        = "base"
	location      string
)

/*sadfa
 */
func CmodelFile(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	location = c.String("location")
	file := c.String("file")
	if location == "" {
		fmt.Println("location不能为空")
		return M_ERRLOCATION
	}
	if file == "" {
		fmt.Println("您未指定struct文件")
		return M_ERRNOFILE
	}
	//在当前目录下生成md目录
	os.Mkdir(m_modelPath, os.ModePerm)
	os.Mkdir(m_modelPath+m_base, os.ModePerm)

	//创建model
	f, _ := os.Stat(file)
	if f.IsDir() { //如果是目录
		err := filepath.Walk(file, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if !f.IsDir() {
				createModel(path)
				createDAO(path)
				createBaseModel(dorm.Keytype) //创建basemodel

			}
			return nil
		})
		return err
	} else {
		err := createDAO(path)
		err = createModel(file)
		return err
	}

}

//生成DAO
func createDAO(file string) error {
	sl := new(lexer.StructLexer)
	fileStr := sl.GetStructFile(file)
	arrStruct := sl.StructStr(fileStr)

	for _, v := range arrStruct {
		Str := dorm.CreateDAO(location, m_packageName, v)
		tname := ut.CalToUnder(sl.StructName(v))
		//生成文件
		fileName := tname + "DAO.go"
		// fPath := m_modelPath + m_base + "/"
		exist, err := ut.FileOrPathExists(m_modelPath + fileName)
		if exist { //存在就不覆盖
			fmt.Println(util.Red(m_modelPath + fileName + "文件已存在"))
			return err
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

//生成basemodel
func createBaseModel(keytype string) error {
	Str := dorm.CreateModel(m_base, keytype)
	//生成文件
	fileName := "model.go"
	fPath := m_modelPath + m_base + "/"
	exist, err := ut.FileOrPathExists(fPath + fileName)
	if !COVRE && exist {
		fmt.Println(util.Red(fPath + fileName + "文件已存在"))
		return err
	}
	sf, err := os.Create(fPath + fileName)
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
		fmt.Println(fPath + fileName + "生成成功")
	}
	return nil
}

//生成model
func createModel(file string) error {
	sl := new(lexer.StructLexer)
	fileStr := sl.GetStructFile(file)
	arrStruct := sl.StructStr(fileStr)
	fPath := m_modelPath + m_base + "/"
	for _, v := range arrStruct {
		Str := dorm.CreateDorm(m_base, v)
		tname := ut.CalToUnder(sl.StructName(v))
		//生成文件
		fileName := tname + ".go"
		exist, err := ut.FileOrPathExists(fPath + fileName)
		if !COVRE && exist {
			fmt.Println(util.Red(fPath + fileName + "文件已存在"))
			return err
		}
		sf, err := os.Create(fPath + fileName)
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
			fmt.Println(fPath + fileName + "生成成功")
		}
	}
	return nil
}
