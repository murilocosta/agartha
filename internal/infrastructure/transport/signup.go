package transport

import "github.com/gin-gonic/gin"

type survivorSignUpCtrl struct {
}

func NewSurvivorSignUpCtrl() *survivorSignUpCtrl {
	return &survivorSignUpCtrl{}
}

func (ctrl *survivorSignUpCtrl) Execute(c *gin.Context) {

}
