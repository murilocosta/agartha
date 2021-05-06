package transport

import "github.com/gin-gonic/gin"

type Ctrl interface {
	Execute(c *gin.Context)
}
