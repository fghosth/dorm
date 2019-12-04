package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

type dir struct {
	name   string
	Subdir []string
}

//initProject 生成目录
func initProject(c *cli.Context) error {
	pathP := c.String("file")
	var err error
	//创建目录
	for _, dv := range getDirList() { //遍历dir struct
		d := pathP + "/" + dv.name
		err = os.Mkdir(d, os.ModePerm)
		if err != nil {
			return err
		}
		if dv.Subdir != nil { //如果子目录不为空
			for _, sdv := range dv.Subdir { //遍历子目录
				sd := d + "/" + sdv
				err = os.Mkdir(sd, os.ModePerm)
				if err != nil {
					return err
				}
			}
		}
	}

	//创建文件
	readmeFile := pathP + "/README.md"
	readmeCon := []byte("# 项目说明")
	err = writeFile(readmeFile, readmeCon)
	if err != nil {
		return err
	}
	gomodFile := pathP + "/go.mod"
	var module string
	if pathP == "./" {
		module = getCurrentPath()
	} else {
		pathArr := strings.Split(pathP, "/")
		module = pathArr[len(pathArr)-1]
	}
	gomodCon := []byte(`module ` + module + `

go 1.12
replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.38.0
	github.com/go-tomb/tomb => gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
	go.opencensus.io => github.com/census-instrumentation/opencensus-go v0.21.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.1.0
	go.uber.org/zap => github.com/uber-go/zap v1.10.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190507160741-ecd444e8653b
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190508025753-952990169864
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.4.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190502173448-54afdca5d873
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
	gopkg.in/alecthomas/kingpin.v2 => github.com/alecthomas/kingpin v2.2.6+incompatible
	gopkg.in/mgo.v2 => github.com/go-mgo/mgo v0.0.0-20180705113738-7446a0344b78
	gopkg.in/vmihailenco/msgpack.v2 => github.com/vmihailenco/msgpack v4.0.4+incompatible
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v2.1.0+incompatible
	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20180705113738-7446a0344b78
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
	`)
	err = writeFile(gomodFile, gomodCon)
	if err != nil {
		return err
	}

	//makefile
	mFile := "Makefile"
	err = writeFile(mFile, makefile)
	if err != nil {
		return err
	}

	return err
}

func getDirList() []dir {
	d := make([]dir, 6)
	d[0] = dir{"cmd", nil}
	d[1] = dir{"conf", nil}
	d[2] = dir{"context", nil}
	d[3] = dir{"domain", []string{"data", "role", "object"}}
	d[4] = dir{"infrastructure", []string{"util", "mysql", "mq"}}
	d[5] = dir{"service", nil}
	return d
}

//writeFile 写入文件
func writeFile(file string, content []byte) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	defer f.Close()
	return err
}

func getCurrentPath() string {
	s, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}
	i := strings.LastIndex(s, "/")
	path := string(s[i+1 : len(s)])
	return path
}

var makefile = []byte(`
APP = xxx
CONF = conf.yaml
MAINDIR = ./cmd
default:
	@echo 'Usage of make: [ build | linux_build | windows_build | clean ]'

build:
	@go build -o ./dist/$(APP) $(MAINDIR)
	@cp ./conf/$(CONF) ./dist

linux_build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/$(APP) $(MAINDIR)
	@cp ./conf/$(CONF) ./dist

windows_build:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/$(APP).exe $(MAINDIR)
	@cp ./conf/$(CONF) ./dist

run: build
	@./dist/$(APP) -f ./dist/$(CONF)

install: build
	@mv ./dist/$(APP) $(GOPATH)/$(APP)

clean:
	@rm -f ./dist/*
`)
