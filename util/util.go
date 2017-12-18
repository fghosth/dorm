package util

import (
	// "fmt"
	"errors"
	"os"
	"regexp"
	"strings"
)

type Dstring struct {
}

//获取``等符号中的内容
func (ds *Dstring) PixContent(str string, symbol string) string {
	reg := symbol + ".+" + symbol
	r := regexp.MustCompile(reg)
	content := r.FindString(str)
	content = strings.Replace(content, "`", "", -1)
	return content
}

//首字母转小写
func (ds *Dstring) FUPer(str string) (string, error) {
	errEmpty := errors.New("字符串为空")
	v := []byte(str)
	if len(v) == 0 {
		return "", errEmpty
	}
	if v[0] < 97 {
		v[0] += 32
	}
	return string(v), nil
}

//判断文件或目录是否存在
func (ds *Dstring) FileOrPathExists(path string) (bool, error) {
	errFileExist := errors.New(path + "文件已存在")
	_, err := os.Stat(path)
	if err == nil {
		return true, errFileExist
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//获取方法
func (ds *Dstring) GetMeth(mby string) string {
	mbyArr := strings.Split(mby, ",")
	mby = ""
	for _, v := range mbyArr { //循环通过空格取出每个参数，去除关键字，再拼接
		v = strings.TrimLeft(v, " ")
		if strings.Contains(v, " ") { //有关键字的才处理
			// fmt.Println(strings.Split(v, " ")[0])
			if strings.EqualFold(mby, "") { //字符串拼接
				mby = strings.Split(v, " ")[0]
			} else {
				mby = mby + "," + strings.Split(v, " ")[0]
			}

		} else {
			mby = mby + v
		}
	}
	if !strings.Contains(mby, ")") { //检查是否漏掉右括号
		mby = mby + ")"
	}
	return mby
}
