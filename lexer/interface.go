package lexer

import (
	"regexp"
	"strings"
)

var (
	//所有intreface 片段
	interfaceArr = `(?i)(type ).+(interface)[ ]+{[\W\w]+?(}[ ]*\n)`
	//注释
	commentArr = `(/\*)[\W\w]+?(\*/)`
	//function字符串
	functionStr = `[A-Z].*\(.*\).*`
	//functin名字
	fnameStr = `[A-Z][a-zA-Z0-9_-]*`
	//functing内的参数:含()
	argStr = `\(.*?\)`
	//返回值字符串 包含) 需去除
	returnStr = `\).*`
	//获得包名行
	packageName = `(package).*`
	//获取interface行
	interfaceLine = `(type).*(interface)[ ]*{`
)

type InterfaceLexer struct {
}

type FuncM struct {
	FunName string
	Comment string
	Args    string
	Returns string
}

var IfFuncM []FuncM

func (il InterfaceLexer) GetIfFuncM(str string) []FuncM {
	ifstr := il.GetInterfaceStr(str)
	ifM := &FuncM{}
	funLine := getFunctionLine(ifstr)
	commentLine := getComment(ifstr)
	for k, v := range funLine {
		ifM.Comment = commentLine[k]
		ifM.FunName = getFunName(v)
		ifM.Args = getFunArgs(v)
		ifM.Returns = getFunReturns(v)
		IfFuncM = append(IfFuncM, *ifM)
	}
	return IfFuncM
}

func (il InterfaceLexer) GetInterfaceStr(str string) string {
	r := regexp.MustCompile(interfaceArr)
	s := r.FindString(str)
	return s
}

func (il InterfaceLexer) GetPackageName(filestr string) string {
	r := regexp.MustCompile(packageName)
	s := r.FindString(filestr)
	s = strings.Split(s, " ")[1]
	return s
}
func (il InterfaceLexer) GetServiceName(ifstr string) string {
	r := regexp.MustCompile(interfaceLine)
	s := r.FindString(ifstr)
	if len(s) > 10 {
		s = strings.TrimSpace(strings.Split(s, " ")[1])
	}
	return s
}

//根据参数行得到里面的字段
func GetFunFields(funLine string) string {
	var str string
	funLine = strings.Replace(funLine, "(", "", -1)
	funLine = strings.Replace(funLine, ")", "", -1)
	flarr := strings.Split(funLine, ",") //,号分割后去除空格后面的关键字
	for i := 0; i < len(flarr); i++ {
		flarr[i] = strings.Split(strings.TrimSpace(flarr[i]), " ")[0]
	}
	str = "(" + strings.Join(flarr, ", ") + ")"
	return str
}

//获取所有注释
func getComment(str string) []string {
	r := regexp.MustCompile(commentArr)
	s := r.FindAllString(str, -1)
	return s
}

//获取所有functin行
func getFunctionLine(ingerfacestr string) []string {
	r := regexp.MustCompile(functionStr)
	s := r.FindAllString(ingerfacestr, -1)
	return s
}

//通过单行fun获得单个名字
func getFunName(funstr string) string {
	r := regexp.MustCompile(fnameStr)
	s := r.FindString(funstr)
	return s
}

//通过单行fun获得参数
func getFunArgs(funstr string) string {
	r := regexp.MustCompile(argStr)
	s := r.FindString(funstr)
	return s
}

//通过单行fun获得返回值
func getFunReturns(funstr string) string {
	r := regexp.MustCompile(returnStr)
	s := r.FindString(funstr)
	s = strings.Replace(s, ")", "", 1)
	return s
}
