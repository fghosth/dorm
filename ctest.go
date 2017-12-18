package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"jvole.com/createProject/util"
	// "io/ioutil"
)

const (
	TEST_FILE = "service_test.go"
)

/*
CgokitFile
生成gokit文件
*/
func CTestFile(c *cli.Context) error {
	fmt.Println("===================获取接口文件")
	err := getSI()
	if err != nil {
		fmt.Println(util.Red("失败============获取接口文件:" + err.Error()))
		return err
	}
	fmt.Println(util.Blue("成功============获取接口文件"))

	fmt.Println("===================生成service的测试文件")
	err = createTestSV()
	if err != nil {
		fmt.Println(util.Red("失败============生成service的测试文件:" + err.Error()))
		return err
	}
	fmt.Println(util.Blue("成功============生成service的测试文件"))

	return nil
}

//*只能包含一个接口
func createTestSV() error {

	exist, err := ut.FileOrPathExists(path + TEST_FILE)
	if !COVRE && exist {
		return err
	}

	str := "package " + PName + "_test\n"
	str = str + `import(
	_ "fmt"
	"testing"
	)` + "\n"

	//添加接测试方法
	for mname, _ := range ServiceI {
		str = str + "func Test" + mname + "(t *testing.T) {\n"

		str = str + "}\n"
	}
	sf, err := os.Create(path + TEST_FILE)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		// fmt.Println(err)
		return err
	}

	return nil
}
