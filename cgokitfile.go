package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/bndr/gotabulate"
	"github.com/urfave/cli"
	"jvole.com/createProject/lexer"
	"jvole.com/createProject/util"
	// "io/ioutil"
)

var ServiceI []lexer.FuncM

const (
	ENDPOINT   = "endpoints.go"
	TRANSPORT  = "transport.go"
	LOG        = "logging.go"
	SERVICE    = "service.go"
	PROXY      = "proxying.go"
	INSTRUMENT = "instrumenting.go"
	UTIL_FILE  = "util/util.go"
)

var (
	COVRE        = false  //已存在文件是否覆盖true为覆盖
	path         = "./"   //路径
	IName        = ""     //接口名称
	pat          = "[ ]+" //正则 查询接口名称去除多个空格用
	PName        = ""     //包名
	ut           = new(util.Dstring)
	SERVICE_FILE = "serviceI.go"
	SERVICE_NAME string //接口服务的名称
)

var (
	ErrNoInterface = errors.New("此文件不包含接口")
)

/*
CgokitFile
生成gokit文件
*/
func CgokitFile(c *cli.Context) error {
	var string_1, string_2, string_3, string_4, string_5, string_6 []string
	if c.String("cover") == "true" {
		COVRE = true
	} else {
		COVRE = false
	}
	if c.String("file") != "" {
		SERVICE_FILE = c.String("file")
	}
	//获取接口信息，初始化变量
	err := getSI()
	if err != nil {
		fmt.Println(util.Red("失败============获取接口文件:" + err.Error()))
		return err
	}

	if c.String("c") == "all" {

		fmt.Println("===================创建service.go文件")
		err = createService()
		if err != nil {
			string_6 = []string{"service.go", "failed", err.Error()}
		} else {
			string_6 = []string{"service.go", "success", ""}
		}

		// fmt.Println("===================创建util.go文件")
		// err = createUtil()
		// if err != nil {
		// 	string_1 = []string{"Util.go", "failed", err.Error()}
		// } else {
		// 	string_1 = []string{"Util.go", "success", ""}
		// }

		fmt.Println("===================创建endpoint.go文件")
		err = createEndpoint()
		if err != nil {
			string_2 = []string{"endpoint.go", "failed", err.Error()}
		} else {
			string_2 = []string{"endpoint.go", "success", ""}
		}

		fmt.Println("===================创建logging.go文件")
		err = createLogging()
		if err != nil {
			string_3 = []string{"logging.go", "failed", err.Error()}
		} else {
			string_3 = []string{"logging.go", "success", ""}
		}

		fmt.Println("===================创建instrumenting.go文件")
		err = createInstrument()
		if err != nil {
			string_4 = []string{"instrumenting.go", "failed", err.Error()}
		} else {
			string_4 = []string{"instrumenting.go", "success", ""}
		}

		fmt.Println("===================创建Transport.go文件")
		err = createTransport()
		if err != nil {
			string_5 = []string{"Transport.go", "failed", err.Error()}
		} else {
			string_5 = []string{"Transport.go", "success", ""}
		}

	} else {
		fmt.Println("added task: ", c.String("c"))
		list := strings.Split(c.String("c"), ":")
		for _, v := range list {
			// fmt.Println(v)
			switch v {
			case "service":
				fmt.Println("===================创建service.go文件")
				err := createService()
				if err != nil {
					string_6 = []string{"service.go", "failed", err.Error()}
				} else {
					string_6 = []string{"service.go", "success", ""}
				}
			case "logging":
				fmt.Println("===================创建logging.go文件")
				err := createLogging()
				if err != nil {
					string_3 = []string{"logging.go", "failed", err.Error()}
				} else {
					string_3 = []string{"logging.go", "success", ""}
				}
			case "transport":
				fmt.Println("===================创建Transport.go文件")
				err := createTransport()
				if err != nil {
					string_5 = []string{"transport.go", "failed", err.Error()}
				} else {
					string_5 = []string{"transport.go", "success", ""}
				}
			case "endpoint":
				fmt.Println("===================创建endpoint.go文件")
				err := createEndpoint()
				if err != nil {
					string_2 = []string{"endpoint.go", "failed", err.Error()}
				} else {
					string_2 = []string{"endpoint.go", "success", ""}
				}
			case "instrumenting":
				fmt.Println("===================创建instrumenting.go文件")
				err := createInstrument()
				if err != nil {
					string_4 = []string{"instrumenting.go", "failed", err.Error()}
				} else {
					string_4 = []string{"instrumenting.go", "success", ""}
				}
				// case "util":
				// 	fmt.Println("===================创建util.go文件")
				// 	err := createUtil()
				// 	if err != nil {
				// 		string_1 = []string{"util.go", "failed", err.Error()}
				// 	} else {
				// 		string_1 = []string{"util.go", "success", ""}
				// 	}
			}
		}
	}
	//输出统计结果
	// Create Object
	tabulate := gotabulate.Create([][]string{string_1, string_2, string_3, string_4, string_5, string_6})
	// Set Headers
	tabulate.SetHeaders([]string{"file", "result", "reason"})
	// Set Align (Optional)
	tabulate.SetAlign("right")
	// Set Max Cell Size
	tabulate.SetMaxCellSize(32)
	// Set the Empty String (optional)
	tabulate.SetEmptyString("None")
	// Turn On String Wrapping
	tabulate.SetWrapStrings(true)
	fmt.Println(tabulate.Render("grid"))
	return nil
}

/*
生成transport.go

*/
func createTransport() error {
	exist, err := ut.FileOrPathExists(path + TRANSPORT)
	if !COVRE && exist {
		return err
	}
	serverfield := make([]string, len(ServiceI))
	handlefield := make([]string, len(ServiceI))
	decodeRequestfield := make([]string, len(ServiceI))
	for _, v := range ServiceI {
		str := `
e` + v.FunName + ` := make` + v.FunName + `Endpoint(bs)
e` + v.FunName + ` = basic.AuthMiddleware(User, Password, "")(e` + v.FunName + `)
e` + v.FunName + ` = middleware.ValidMiddleware()(e` + v.FunName + `)
e` + v.FunName + ` = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e` + v.FunName + `)
` + v.FunName + `Handler := kithttp.NewServer(
	e` + v.FunName + `,
	decode` + v.FunName + `Request,
	encodeResponse,
	opts...,
)
		`
		serverfield = append(serverfield, str)
	}

	for _, v := range ServiceI {
		str := `r.Handle("/v1/` + v.FunName + `", ` + v.FunName + `Handler).Methods("POST")`
		handlefield = append(handlefield, str)
	}

	for _, v := range ServiceI {
		str := `
func decode` + v.FunName + `Request(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {

	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	defer func(begin time.Time) {
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		level.Debug(util.KitLogger).Log(
			"method", f.Name(),
			"file", path.Base(file),
			"line", line,
			"request", body,
			"took", time.Since(begin).Nanoseconds()/1000,
		)
	}(time.Now())
	return ` + v.FunName + `Request{

	}, nil
}
	`
		decodeRequestfield = append(decodeRequestfield, str)
	}

	//根据模板生成
	ctx := map[string]interface{}{
		"pName":              PName,
		"sname":              SERVICE_NAME,
		"serverfield":        serverfield,
		"handlefield":        handlefield,
		"decodeRequestfield": decodeRequestfield,
	}

	str, err := raymond.Render(TPL_TRANSPORT, ctx)
	if err != nil {
		fmt.Println(err)
	}
	sf, err := os.Create(path + TRANSPORT)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

/*
生成 instrumenting.go
*/
func createInstrument() error {
	exist, err := ut.FileOrPathExists(path + INSTRUMENT)
	if !COVRE && exist {
		return err
	}
	funfield := make([]string, len(ServiceI))
	for _, v := range ServiceI {
		str := `
			func (s *instrumentingService) ` + v.FunName + v.Args + v.Returns + ` {
				defer func(begin time.Time) {
					s.requestCount.With("method", "Insert").Add(1)
					s.requestLatency.With("method", "` + v.FunName + `").Observe(time.Since(begin).Seconds())
				}(time.Now())
				return s.next.` + v.FunName + lexer.GetFunFields(v.Args) + `
			}
		`
		funfield = append(funfield, str)
	}
	//根据模板生成
	ctx := map[string]interface{}{
		"pName":    PName,
		"sname":    SERVICE_NAME,
		"funfield": funfield,
	}

	str, err := raymond.Render(TPL_INSTRUMENTING, ctx)
	if err != nil {
		fmt.Println(err)
	}

	sf, err := os.Create(path + INSTRUMENT)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

//生成logging.go
func createLogging() error {
	exist, err := ut.FileOrPathExists(path + LOG)
	if !COVRE && exist {
		return err
	}
	funfield := make([]string, len(ServiceI))
	for _, v := range ServiceI {
		str := `
		func (s *loggingService) ` + v.FunName + v.Args + v.Returns + ` {
			defer func(begin time.Time) {
				pc, file, line, _ := runtime.Caller(1)
				f := runtime.FuncForPC(pc)
				level.Info(s.logger).Log(
					"method", f.Name(),
					"file", path.Base(file),
					"line", line,
					"took", time.Since(begin).Nanoseconds()/1000,
				)
			}(time.Now())

			return s.next.` + v.FunName + lexer.GetFunFields(v.Args) + `
		}
	`
		funfield = append(funfield, str)
	}
	//根据模板生成
	ctx := map[string]interface{}{
		"pName":    PName,
		"sname":    SERVICE_NAME,
		"funfield": funfield,
	}

	str, err := raymond.Render(TPL_LOGGING, ctx)
	if err != nil {
		fmt.Println(err)
	}

	sf, err := os.Create(path + LOG)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		// fmt.Println(err)
		return err
	}

	return nil
}

//生成endpoint.go文件
func createEndpoint() error {
	exist, err := ut.FileOrPathExists(path + ENDPOINT)
	if !COVRE && exist {
		return err
	}
	funfield := make([]string, len(ServiceI))
	for _, v := range ServiceI {
		str := `
			type ` + v.FunName + `Request struct {
				Uid   uint64
				Tags  map[string]string
				Field map[string]interface{}
				Table string
			}

			func make` + v.FunName + `Endpoint(s ` + SERVICE_NAME + `) endpoint.Endpoint {
				return func(ctx context.Context, request interface{}) (interface{}, error) {
					req := request.(` + v.FunName + `Request)
					return Response{Errcode: "0", Msg: "ok", Data: nil, Err: nil}, nil
				}
			}
		`
		funfield = append(funfield, str)
	}
	//根据模板生成
	ctx := map[string]interface{}{
		"pName":    PName,
		"sname":    SERVICE_NAME,
		"funfield": funfield,
	}

	str, err := raymond.Render(TPL_ENDPOINTS, ctx)
	if err != nil {
		fmt.Println(err)
	}

	sf, err := os.Create(path + ENDPOINT)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

//获得接口信息
func getSI() error {
	data, err := ioutil.ReadFile(path + SERVICE_FILE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	ifLexer := new(lexer.InterfaceLexer)
	ifstr := ifLexer.GetInterfaceStr(string(data))
	ServiceI = ifLexer.GetIfFuncM(ifstr)
	SERVICE_NAME = ifLexer.GetServiceName(ifstr)
	PName = ifLexer.GetPackageName(string(data))
	return nil
	// br :=

}

//创建util文件
func createUtil() error {
	exist, err := ut.FileOrPathExists(path + UTIL_FILE)
	// fmt.Println(err)
	if !COVRE && exist {
		// fmt.Println(err)
		return err
	}
	exist, err = ut.FileOrPathExists(path + "util")
	if !exist {
		err = os.Mkdir(path+"util", os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			// fmt.Println(err)
			return err
		}
	}

	sf, err := os.Create(path + UTIL_FILE)
	str := TPL_UTIL
	if err != nil {
		// fmt.Println(err)
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

//*只能包含一个接口
func createService() error {
	var str string
	iname, _ := ut.FUPer(SERVICE_NAME)
	funfield := make([]string, len(ServiceI))
	for _, v := range ServiceI {
		str := v.Comment + `
			func (s ` + iname + `) ` + v.FunName + v.Args + v.Returns + ` {
				return nil
			}
		`
		funfield = append(funfield, str)
	}

	//根据模板生成
	ctx := map[string]interface{}{
		"pName":    PName,
		"sname":    iname,
		"funfield": funfield,
	}

	str, err := raymond.Render(TPL_SERVICE, ctx)
	if err != nil {
		fmt.Println(err)
	}

	sf, err := os.Create(path + SERVICE)
	if err != nil {
		return err
	}
	_, err = sf.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
