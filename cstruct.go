package main

import (
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
	"github.com/urfave/cli"
	"jvole.com/createProject/model"
)

/*sadfa
 */
func CstructFile(c *cli.Context) error {
	sql := "select cash,gender,create_time,payment,address,userid,name,card from user"
	args := make([]interface{}, 0)
	user := new(model.User)
	// pp.Println(user)
	rows, err := model.Db.Query(sql, args...)
	defer rows.Close()
	if err != nil {
		fmt.Println("sql open error ", err)
	}
	columns, _ := rows.Columns()
	str := structFieldString(user, len(columns), columns)
	pp.Println(str)
	return nil
}

func structFieldString(obj interface{}, length int, col []string) []string {
	values := make([]string, length)
	// reflectStruct := reflect.ValueOf(obj).Elem()

	rtype := reflect.TypeOf(obj).Elem()
	for i, v := range col {
		tname := ""
		// fmt.Println(rtype.NumField())
		for k := 0; k < rtype.NumField(); k++ {
			if rtype.Field(k).Tag.Get("dorm") == v {
				// fmt.Println(v, "------", rtype.Field(k).Tag.Get("dorm"))
				tname = rtype.Field(k).Name
				break
			}
		}
		// tagn := rtype.Field(i).Tag.Get("bson")
		if tname != "" {
			// fmt.Println(tname, "===", v)
			sname, _ := ut.FUPer(rtype.Name())
			values[i] = "&" + sname + "." + tname
		}
	}
	return values
}

func test() error {
	sql := "select cash,gender,create_time,payment,address,userid,name,card from user"
	args := make([]interface{}, 0)
	// args[0] = "1"
	user := new(model.User)
	ulist, err := user.Find(sql, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pp.Println(ulist)
	return nil
}
