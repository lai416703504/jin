package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handlers       []ControllerHandler
	index          int

	//是否超时标记位
	hasTimeout bool
	//写保护机制
	writerMux *sync.Mutex

	// url路由匹配的参数
	params map[string]string
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		index:          -1,
		writerMux:      &sync.Mutex{},
	}
}

//region base function

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// 核心函数，调用context的下一个函数
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}

	return nil
}

//endregion

func (ctx Context) BaseContext() context.Context {
	return ctx.request.Context()
}

//region implement context.Context
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

//endregion

//region query url

//func (ctx *Context) QueryInt(key string, def int) int {
//	params := ctx.QueryAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			intval, err := strconv.Atoi(vals[len-1])
//			if err != nil {
//				return def
//			}
//
//			return intval
//		}
//	}
//
//	return def
//}

//func (ctx *Context) QueryString(key string, def string) string {
//	params := ctx.QueryAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			return vals[len-1]
//		}
//	}
//
//	return def
//}

//func (ctx *Context) QueryArray(key string, def []string) []string {
//	params := ctx.QueryAll()
//	if vals, ok := params[key]; ok {
//		return vals
//	}
//
//	return def
//}

//
//func (ctx *Context) QueryAll() map[string][]string {
//	if ctx.request != nil {
//		return map[string][]string(ctx.request.URL.Query())
//	}
//
//	return map[string][]string{}
//}

//endregion

//region form post

//func (ctx *Context) FormInt(key string, def int) int {
//	params := ctx.FormAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			intval, err := strconv.Atoi(vals[len-1])
//			if err != nil {
//				return def
//			}
//
//			return intval
//		}
//	}
//
//	return def
//}

//func (ctx *Context) FormString(key string, def string) string {
//	params := ctx.FormAll()
//	if vals, ok := params[key]; ok {
//		len := len(vals)
//		if len > 0 {
//			return vals[len-1]
//		}
//	}
//
//	return def
//}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}

	return def
}

//func (ctx *Context) FormAll() map[string][]string {
//	if ctx.request != nil {
//		return map[string][]string(ctx.request.PostForm)
//	}
//
//	return map[string][]string{}
//}

//endregion

//region application/json post
//func (ctx *Context) BindJson(obj interface{}) error {
//	if ctx.request != nil {
//		body, err := ioutil.ReadAll(ctx.request.Body)
//
//		if err != nil {
//			return err
//		}
//		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
//
//		err = json.Unmarshal(body, obj)
//		if err != nil {
//			return err
//		}
//	} else {
//		return errors.New("ctx.request empty")
//	}
//
//	return nil
//}

//endregion

//region response
//func (ctx *Context) Json(status int, obj interface{}) error {
//	if ctx.HasTimeout() {
//		return nil
//	}
//
//	ctx.responseWriter.Header().Set("Content-Type", "application/json")
//	ctx.responseWriter.WriteHeader(status)
//	byt, err := json.Marshal(obj)
//	if err != nil {
//		ctx.responseWriter.WriteHeader(500)
//		return err
//	}
//	ctx.responseWriter.Write(byt)
//	return nil
//}

//func (ctx *Context) HTML(status int, obj interface{}, template string) error {
//	return nil
//}

//func (ctx *Context) Text(status int, obj string) error {
//	return nil
//}

//endregion
