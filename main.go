package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "createproject"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "derek",
			Email: "fghosth@163.com",
		},
	}
	app.Copyright = "(c) derek fan"
	app.Usage = "创建基于gokit的项目文件"
	app.UsageText = "创建基于gokit的项目文件"
	app.Commands = []cli.Command{
		{
			Name:    "micser",
			Aliases: []string{"ms"},
			Usage:   "根据service生成微服务项目文件",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "component, c",
					Value: "all",
					Usage: "要生成的组件，默认all.可选值：service,logging,proxying(TODO),transport,endpoint,instrumenting,util。多选请用分号分割 logging:proxying:endpoint",
				},
				cli.StringFlag{
					Name:  "cover",
					Value: "false",
					Usage: "是否覆盖已有文件，默认值false.可选值：true，false",
				},
			},
			Action: func(c *cli.Context) error {
				return CgokitFile(c)
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "根据service生成测试文件",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "component, c",
					Value: "none",
					Usage: "使用哪一个组建构建测试文件.可选值：none,goconvey,Testify。",
				},
			},
			Action: func(c *cli.Context) error {
				return CTestFile(c)

			},
		},

		{
			Name:    "createsql",
			Aliases: []string{"cs"},
			Usage:   "根据struct生成sql脚本。支持mysql，Cockroach",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "database, d",
					Value: "mysql",
					Usage: "需生成的数据库脚本.可选值：mysql，Cockroach",
				},
				cli.StringFlag{
					Name:  "file, f",
					Value: "",
					Usage: "struct文件.",
				},
				cli.StringFlag{
					Name:  "cover",
					Value: "false",
					Usage: "是否覆盖已有文件，默认值false.可选值：true，false",
				},
			},
			Action: func(c *cli.Context) error {
				return CsqlFile(c)
			},
		},

		{
			Name:    "createstruct",
			Aliases: []string{"cst"},
			Usage:   "根据数据库脚本生成struct。支持mysql，Cockroach",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "database, d",
					Value: "mysql",
					Usage: "选择数据类型.可选值：mysql，Cockroach",
				},
				cli.StringFlag{
					Name:  "file, f",
					Value: "",
					Usage: "sql文件.",
				},
				cli.StringFlag{
					Name:  "cover",
					Value: "false",
					Usage: "是否覆盖已有文件，默认值false.可选值：true，false",
				},
			},
			Action: func(c *cli.Context) error {
				return CstructFile(c)
			},
		},
	}

	app.Run(os.Args)
}
