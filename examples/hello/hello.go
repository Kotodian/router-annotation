package hello

import "github.com/gin-gonic/gin"

// Middleware 中间件测试
func Middleware(c *gin.Context) {

}

// Middleware2 中间件测试2
func Middleware2(c *gin.Context) {

}

// Hello hello example
// @method: get
// @group: V1
// @router: /hello
func Hello(c *gin.Context) {
}

// Hello2 hello example
// @method: get
// @group: V2
// @use: Middleware
// @router: /hello2
func Hello2(c *gin.Context) {

}

// Hello3 hello example
// @method: get
// @group: V2
// @router: /hello3
func Hello3(c *gin.Context) {

}
