package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router map[string]*Tree // all routers
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}
func (c Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}
func (c Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// ==== http method wrap end

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// FindRouterByRequest 需求2：匹配路由，如果没有匹配到，则返回nil
func (c *Core) FindRouterByRequest(r *http.Request) ControllerHandler {
	// url 和method全部转换为大写，保证大小写不敏感
	uri := r.URL.Path
	method := r.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {

		return methodHandlers.FindHandler(uri)
	}
	return nil

}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServeHTTP")
	// 封装自定义context
	ctx := NewContext(request, response)

	//寻找路由
	router := c.FindRouterByRequest(request)
	if router == nil {

		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")

	}
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
	}

}
