package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/urfave/cli"
	"jvole.com/createProject/gotests"
	"jvole.com/createProject/util"
	// "io/ioutil"
)

const (
	TEST_FILE = "service_test.go"
)

type Options struct {
	OnlyFuncs     string // Regexp string for filter matches.
	ExclFuncs     string // Regexp string for excluding matches.
	ExportedFuncs bool   // Only include exported functions.
	AllFuncs      bool   // Include all non-tested functions.
	PrintInputs   bool   // Print function parameters as part of error messages.
	Subtests      bool   // Print tests using Go 1.7 subtests
	WriteOutput   bool   // Write output to test file(s).
}

/*
CgokitFile
生成gokit文件
*/
func CTestFile(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	file := c.String("DIR")
	if file == "" {
		fmt.Println("您未指定目录")
		return ERRNOFILE
	}

	return createTestSV(file)

}

//*只能包含一个接口
func createTestSV(file string) error {
	popt := &Options{
		OnlyFuncs:     "",
		ExclFuncs:     "",
		ExportedFuncs: false,
		AllFuncs:      true,
		PrintInputs:   false,
		Subtests:      false,
		WriteOutput:   false,
	}
	gopt := parseOptions(os.Stdout, popt)
	if gopt == nil {
		fmt.Println("配置错误")
		return errors.New("配置错误")
	}

	gotests.GenerateTests(file, gopt)

	testout := gotests.TestOut
	if len(testout) == 0 {
		fmt.Println("测试文件已生成，如需重新生成请删除测试文件。")
		return errors.New("测试文件已生成，如需重新生成请删除测试文件。")
	}
	for _, v := range testout {
		filename := v.FileName
		testStr := v.TestStr
		exist, err := ut.FileOrPathExists(filename)

		if !COVRE && exist {
			fmt.Println(util.Red(filename + "文件已存在"))
			// return err
		}
		sf, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = sf.Write([]byte(testStr))
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !(!COVRE && exist) && err == nil {
			fmt.Println(filename + "生成成功")
		}
	}
	// pp.Println(gotests.TestOut)

	return nil
}

func parseOptions(out io.Writer, opt *Options) *gotests.Options {
	if opt.OnlyFuncs == "" && opt.ExclFuncs == "" && !opt.ExportedFuncs && !opt.AllFuncs {
		fmt.Fprintln(out, "Please specify either the -only, -excl, -export, or -all flag")
		return nil
	}
	onlyRE, err := parseRegexp(opt.OnlyFuncs)
	if err != nil {
		fmt.Fprintln(out, "Invalid -only regex:", err)
		return nil
	}
	exclRE, err := parseRegexp(opt.ExclFuncs)
	if err != nil {
		fmt.Fprintln(out, "Invalid -excl regex:", err)
		return nil
	}
	return &gotests.Options{
		Only:        onlyRE,
		Exclude:     exclRE,
		Exported:    opt.ExportedFuncs,
		PrintInputs: opt.PrintInputs,
		Subtests:    opt.Subtests,
	}
}

func parseRegexp(s string) (*regexp.Regexp, error) {
	if s == "" {
		return nil, nil
	}
	re, err := regexp.Compile(s)
	if err != nil {
		return nil, err
	}
	return re, nil
}
