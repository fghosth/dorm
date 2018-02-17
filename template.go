package main

const (
	TPL_TRANSPORT = `
package {{{pName}}}

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/golang/time/rate"
	"github.com/gorilla/mux"
	// "golang.org/x/time/rate"
	"github.com/go-kit/kit/auth/basic"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/ratelimit"
	kithttp "github.com/go-kit/kit/transport/http"
)

var errBadRoute = errors.New("bad route")
var User, Password string
var qps = 100000 //访问频率

func MakeHandler(bs {{{sname}}}, logger kitlog.Logger) http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r := mux.NewRouter()
	 {{#each serverfield}}
	{{{this}}}
	 {{/each}}

	{{#each handlefield}}
	{{{this}}}
	{{/each}}

	return r
}

{{#each decodeRequestfield}}
{{{this}}}
{{/each}}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	defer func(begin time.Time) {
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		level.Debug(util.KitLogger).Log(
			"method", f.Name(),
			"file", path.Base(file),
			"line", line,
			"response", response,
			"took", time.Since(begin).Nanoseconds()/1000,
		)
	}(time.Now())
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errBadRoute:
		w.WriteHeader(http.StatusNotFound)
	case errBadRoute:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": "1003",
		"msg":     err.Error(),
		"data":    nil,
	})
}
	`

	TPL_INSTRUMENTING = `
package {{{pName}}}

import (
	_ "fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram

	next ServerRoute
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s {{{sname}}}) {{{sname}}} {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}

{{#each funfield}}
	{{{this}}}
{{/each}}

`

	TPL_LOGGING = `
	package {{{pName}}}

	import (
		_ "fmt"
		"path"
		"runtime"
		"time"

		"github.com/go-kit/kit/log"
		"github.com/go-kit/kit/log/level"
	)

	type loggingService struct {
		logger log.Logger
		next   ServerRoute
	}

	func NewLoggingService(logger log.Logger, s {{{sname}}}) {{{sname}}} {
		return &loggingService{logger, s}
	}

	{{#each funfield}}
		{{{this}}}
	{{/each}}

`

	TPL_ENDPOINTS = `
package {{{pName}}}

import (
	"context"
	_ "fmt"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

{{#each funfield}}
	{{{this}}}
{{/each}}

type Response struct {
	Errcode string
	Msg     string
	Data    map[string]string
	Err     error
}
`

	TPL_SERVICE = `package {{{pName}}}
import(
	_ "fmt"
	"errors"
	)
type {{{sname}}} struct{}


{{#each funfield}}
	{{{this}}}
{{/each}}


`
	TPL_UTIL = `
	package util
import(
  // "fmt"
  "errors"
)
type Dstring struct{

}

func (ds *Dstring) FUPer(str string) (string,error){
  errEmpty := errors.New("字符串为空")
  v := []byte(str)
  if len(v) ==0 {
    return "",errEmpty
  }
  if v[0]<97 {
    v[0] += 32
  }
  return string(v),nil
}

`
)
