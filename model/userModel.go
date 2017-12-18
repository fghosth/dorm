package model

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MAXROWS = 30 //最多查出多少条,-1为不限制
)

var Db *sql.DB

type User struct {
	UserId     int64  `json:"user_id" dorm:"userid"`
	UserName   string `json:"user_name" dorm:"name"`
	Cash       string `json:"user_pos" dorm:"cash"`
	Gender     string `json:"user_pos" dorm:"gender"`
	CreateTime string `json:"user_pos" dorm:"create_time"`
	Payment    string `json:"user_pos" dorm:"payment"`
	Address    string `json:"user_pos" dorm:"address"`
	Card       string `json:"user_pos" dorm:"card"`
}

func init() {
	Db, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/orm?charset=utf8")
}

func (user User) Find(sql string, args ...interface{}) ([]User, error) { //如果使用u *user速度慢3倍
	rows, err := Db.Query(sql, args...)
	defer rows.Close()
	if err != nil {
		fmt.Println("sql open error ", err)
	}

	columns, _ := rows.Columns()
	au := make([]User, 0) //0为可变数组长度
	values := make([]interface{}, len(columns))
	reflectStruct := reflect.ValueOf(&user).Elem()
	rtype := reflect.TypeOf(&user).Elem()
	for i, v := range columns {
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
			values[i] = reflectStruct.FieldByName(tname).Addr().Interface()
		}

	}
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		rows.Scan(values...)
		au = append(au, user)
		num++
	}
	// pp.Println(au)

	return au, nil
}

func (u User) FindByreflect(sql string, args ...interface{}) ([]User, error) { //如果使用u *user速度慢3倍
	rows, err := Db.Query(sql, args...)
	defer rows.Close()
	if err != nil {
		fmt.Println("sql open error ", err)
	}

	columns, _ := rows.Columns()
	au := make([]User, 0) //0为可变数组长度
	values := make([]interface{}, len(columns))
	reflectStruct := reflect.ValueOf(&u).Elem()
	rtype := reflect.TypeOf(&u).Elem()
	for i, v := range columns {
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
			values[i] = reflectStruct.FieldByName(tname).Addr().Interface()
		}

	}
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		rows.Scan(values...)
		au = append(au, u)
		num++
	}
	// pp.Println(au)

	return au, nil
}
