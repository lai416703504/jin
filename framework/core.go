package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心结构
type Core struct {
	router      map[string]*Tree // all routers
	middlewares []ControllerHandler
}

// 初始化框架核心结构
func NewCore() *Core {

	//定义二级map 将二级map写入一级map
	router := make(map[string]*Tree)
	router["GET"] = newTree()
	router["POST"] = newTree()
	router["PUT"] = newTree()
	router["DELETE"] = newTree()

	return &Core{
		router: router,
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

//region debug
func (c *Core) GetDfs() {

	for _, v := range c.router["GET"].root.childs {
		//log.Printf("key:%d,value:%v,value.segment:%v\n", k, v, v.segment)
		dfs(v, 1)
	}

	//for k, v := range c.router["GET"].root.childs {
	//	log.Printf("key:%d,value:%v,value.segment:%v\n", k, v, v.segment)
	//}

}

func dfs(node *node, depth int) {
	log.Printf("depth:%d,node:%v,node.segment:%v\n", depth, node, node.segment)

	if node.childs != nil {
		for _, child := range node.childs {
			dfs(child, depth+1)
		}
	}
	return
}

//endregion

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	//将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}

	//for k, v := range c.router["GET"].root.childs {
	//	log.Printf("key:%d,value:%v,value.segment:%v\n", k, v, v.segment)
	//}

}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	//将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	//将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	//将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) *node {
	uri := request.URL.Path
	method := request.Method

	upperMethod := strings.ToUpper(method)

	//查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		//查找第二层的map
		return methodHandlers.root.matchNode(uri)
	}

	return nil
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, requset *http.Request) {
	//log.Println("core.serveHTTP")

	//封装自定义context
	ctx := NewContext(requset, response)

	//// 一个简单的路由选择器，这里直接写死为测试路由foo
	//router := c.router["foo"]

	//寻找路由
	//router := c.FindRouteByRequest(requset)
	//
	//if router == nil {
	//	//如果没有找到，打印日志
	//	log.Printf("method[%s] router[%s] not found \n", requset.Method, requset.URL.Path)
	//	ctx.Json(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	//	return
	//}.

	//寻找路由
	node := c.FindRouteByRequest(requset)

	if node == nil {
		//如果没有找到，打印日志
		log.Printf("method[%s] router[%s] not found \n", requset.Method, requset.URL.Path)
		ctx.SetStatus(http.StatusNotFound).Json(http.StatusText(http.StatusNotFound))
		return
	}

	//设置context中的handlers字段
	ctx.SetHandlers(node.handlers)

	//if err := router(ctx); err != nil {
	//	ctx.Json(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	//}

	//设置路由参数
	params := node.parseParamsFromEndNode(requset.URL.Path)
	ctx.SetParams(params)

	if err := ctx.Next(); err != nil {
		ctx.SetStatus(http.StatusInternalServerError).Json(http.StatusText(http.StatusInternalServerError))
		return
	}
}
