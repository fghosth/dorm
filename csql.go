package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func CsqlFile(c *cli.Context) error {
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	db := c.String("database")
	sqlfile := c.String("file")
	if db != "mysql" && db != "cockroach" {
		fmt.Println("目前只支持mysql和cockroach")
		return ERRNOFILE
	}
	if sqlfile == "" {
		fmt.Println("您未指定struct文件")
		return ERRNOSQL
	}
	if db == "mysql" {
		return genStructByMysql(sqlfile)
	}
	if db == "cockroach" {
		return genStructByCockroach(sqlfile)
	}
	return nil
}
