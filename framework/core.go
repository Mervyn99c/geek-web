package framework

import "net/http"

// Core 框架核心架构
type Core struct {
}

// NewCore 初始化框架核心架构
func NewCore() *Core {
	return &Core{}
}

// 框架核心架构实现Handler接口
func (c *Core) ServeHTTP(w http.ResponseWriter, r http.Request) {
	// todo
}
