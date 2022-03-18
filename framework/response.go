package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// IResponse 代表返回方法
type IResponse interface {
	// Json 输出
	Json(obj interface{}) IResponse
	// Jsonp 输出
	Jsonp(obj interface{}) IResponse
	//xml 输出
	Xml(obj interface{}) IResponse
	// html 输出
	Html(template string, obj interface{}) IResponse
	//string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse
	//header
	SetHeader(key string, val string) IResponse
	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	// 设置状态码
	SetStatus(code int) IResponse
	// 设置 200 状态
	SetOkStatus() IResponse
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}

	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)

	return ctx
}

func (ctx *Context) Jsonp(obj interface{}) IResponse {
	//获取请求的参数callback
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成 XSS 攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}

	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}

	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}

	// 输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}

	ctx.SetHeader("Content-Type", "application/xml")
	ctx.responseWriter.Write(byt)

	return ctx
}

func (ctx *Context) Html(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}

	// 执行 Execute 方法将 obj 和模版进行结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "applicatin/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})

	return ctx
}

func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}
