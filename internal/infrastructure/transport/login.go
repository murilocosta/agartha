package transport

import "github.com/gin-gonic/gin"

type survivorLoginCtrl struct {
}

func NewSurvivorLoginCtrl() *survivorLoginCtrl {
	return &survivorLoginCtrl{}
}

func (ctrl *survivorLoginCtrl) HandlerFunc(c *gin.Context) (interface{}, error) {
	return nil, nil
}
