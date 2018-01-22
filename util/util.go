package util

import (
	// "fmt"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Dstring struct {
}

func (ds *Dstring) Md5Str(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (ds *Dstring) Checkerr(err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//查看字符串是否包含在另一个字符串中比如abc是否在abc,cde,aaa中
func (ds *Dstring) ContainStrBysplit(s, chars, sep string) bool {
	arrS := strings.Split(s, sep)
	isContain := false
	for _, v := range arrS {
		if v == chars {
			isContain = true
			break
		}
	}
	return isContain
}

//驼峰命名转下划线命名法
func (ds *Dstring) CalToUnder(str string) string {
	regstr := `[A-Z][a-z0-9]+` //所有大写字母开头的词
	tmpstr := ""
	r := regexp.MustCompile(regstr)
	words := r.FindAllString(str, -1)
	for _, v := range words {
		if tmpstr == "" {
			tmpstr = strings.ToLower(v)
		} else {
			tmpstr = tmpstr + "_" + strings.ToLower(v)
		}

	}
	return tmpstr
}

//下划线分割命名法转驼峰命名法
func (ds *Dstring) UnderToCal(str string) string {
	dst := new(Dstring)
	c := strings.Split(str, "_")
	tmpstr := ""
	for _, v := range c {
		tmpstr = tmpstr + dst.FUpRLow(v)
	}
	return tmpstr
}

//单词转换首字母大写,其他都小写
func (ds *Dstring) FUpRLow(str string) string {
	str = strings.ToLower(str)
	v := []byte(str)
	v[0] -= 32
	return string(v)
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
