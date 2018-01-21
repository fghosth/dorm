package tmp

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
)

var (
	MAXROWS = 30 //最多查出多少条,-1为不限制
	TABLE   = "user"
)

var Db *sql.DB //数据库连接

type User struct {
	UserID     int64  `json:"user_id" dormCol:"userid" dorm:"type:bigint(20);unsigned;not null;AUTO_INCREMENT"`
	UserName   string `json:"user_name" dormCol:"name"`
	Cash       string `json:"user_pos" dormCol:"cash"`
	Gender     string `json:"user_pos" dormCol:"gender"`
	CreateTime string `json:"user_pos" dormCol:"create_time"`
	Payment    string `json:"user_pos" dormCol:"payment"`
	Address    string `json:"user_pos" dormCol:"address"`
	Card       string `json:"user_pos" dormCol:"card"`
}

func init() {
	Db, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/orm?charset=utf8")
}

func (user User) Find(sqlHalf string, args ...interface{}) ([]User, error) { //如果使用u *user速度慢3倍
	sqlF := "select userid,name,cash,gender,card,create_time,payment,address from user "
	rows, err := Db.Query(sqlF+sqlHalf, args...)
	defer rows.Close()
	if err != nil {
		fmt.Println("sql open error ", err)
	}
	au := make([]User, 0) //0为可变数组长度
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	values[0] = &user.Cash
	values[1] = &user.Gender
	values[2] = &user.CreateTime
	values[3] = &user.Payment
	values[4] = &user.Address
	values[5] = &user.UserID
	values[6] = &user.UserName
	values[7] = &user.Card
	num := 0
	for rows.Next() {
		if num >= MAXROWS && MAXROWS != -1 {
			break
		}
		rows.Scan(values...)
		au = append(au, user)
		num++
	}
	pp.Println(au)

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
			if rtype.Field(k).Tag.Get("dormCol") == v {
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
