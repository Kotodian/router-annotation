package examples

import "github.com/gin-gonic/gin"

// Hello hello example
// @method: get
// @use
// @group: /api/v1
// @router: /api/v1/hello
func Hello(c *gin.Context) {
}

// HelloMiddleware hello middleware example
// @middleware
func HelloMiddleware(c *gin.Context) {

}
