package dorm

//
// import (
// 	_ "database/sql"
// 	"fmt"
// 	"reflect"
// 	"strconv"
//
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/k0kubun/pp"
// 	"github.com/urfave/cli"
// 	"jvole.com/createProject/model"
// )
//
// const (
// 	COLTAG = "dormCol"
// )
//
// var (
// 	sqlFind = `func (user User) Find(sql string, args ...interface{}) ([]User, error) { //如果使用u *user速度慢3倍
// 		rows, err := Db.Query(sql, args...)
// 		defer rows.Close()
// 		if err != nil {
// 			fmt.Println("sql open error ", err)
// 		}
// 		au := make([]User, 0) //0为可变数组长度
// 		columns, _ := rows.Columns()
// 		values := make([]interface{}, len(columns))
// 		{{#each val}}
// 		{{{this}}}
// 	   {{/each}}
// 		num := 0
// 		for rows.Next() {
// 			if num >= MAXROWS && MAXROWS != -1 {
// 				break
// 			}
// 			rows.Scan(values...)
// 			au = append(au, user)
// 			num++
// 		}
// 		// pp.Println(au)
//
// 		return au, nil
// 	}`
// )
//
// /*sadfa
//  */
// func CstructFile(c *cli.Context) error {
// 	// sql := "select cash,gender,create_time,payment,address,userid,name,card from user where userid = ?"
// 	args := make([]interface{}, 1)
// 	args[0] = "1"
// 	user := new(model.User)
// 	vala, sqlstr := structFieldString(user)
// 	pp.Println(vala)
// 	pp.Println(sqlstr)
//
// 	// ctx := map[string]interface{}{
// 	// 	"val": vala,
// 	// }
// 	// output, err := raymond.Render(sqlFind, ctx)
// 	// pp.Println(output)
// 	// pp.Println(sql)
// 	// user.Find(sql, args...)
//
// 	return nil
// }
//
// //根据数据库中的字段对应struct字段,sqlstr
// func structFieldString(obj interface{}) ([]interface{}, string) {
// 	sqlStr := "" //sql需查出的字段
// 	rtype := reflect.TypeOf(obj).Elem()
// 	// sname, _ := ut.FUPer(rtype.Name())
//
// 	sql := "select * from " + model.TABLE + " limit 1"
// 	rows, err := model.Db.Query(sql)
// 	if err != nil {
// 		fmt.Println("sql open error ", err)
// 	}
// 	columns, _ := rows.Columns()
// 	values := make([]interface{}, len(columns))
// 	// reflectStruct := reflect.ValueOf(obj).Elem()
//
// 	for i, v := range columns {
// 		tname := ""
// 		// fmt.Println(rtype.NumField())
// 		for k := 0; k < rtype.NumField(); k++ {
// 			if rtype.Field(k).Tag.Get(COLTAG) == v {
// 				// fmt.Println(v, "------", rtype.Field(k).Tag.Get("dorm"))
// 				tname = rtype.Field(k).Name
// 				break
// 			}
// 		}
//
// 		// tagn := rtype.Field(i).Tag.Get("bson")
// 		if tname != "" {
// 			// fmt.Println(tname, "===", v)
// 			if sqlStr == "" { //sql需查出的字段
// 				sqlStr = v
// 			} else {
// 				sqlStr = sqlStr + "," + v
// 			}
// 			values[i] = "values[" + strconv.Itoa(i) + "]=&" + model.TABLE + "." + tname
// 		}
// 	}
// 	return values, sqlStr
// }
//
// // func test() error {
// // 	sql := "select cash,gender,create_time,payment,address,userid,name,card from user"
// // 	args := make([]interface{}, 0)
// // 	// args[0] = "1"
// // 	user := new(model.User)
// // 	ulist, err := user.Find(sql, args...)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return err
// // 	}
// // 	pp.Println(ulist)
// // 	return nil
// // }
