package transport

import "github.com/gin-gonic/gin"

type checkSurvivorPermissionCtrl struct {
}

func NewCheckSurvivorPermissionCtrl() *checkSurvivorPermissionCtrl {
	return &checkSurvivorPermissionCtrl{}
}

func (ctrl *checkSurvivorPermissionCtrl) HandlerFunc(data interface{}, c *gin.Context) bool {
	return true
}
