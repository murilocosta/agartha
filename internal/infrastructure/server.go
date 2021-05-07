package infrastructure

import (
	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/infrastructure/transport"
)

type ServerMethod uint

const (
	ServerGet ServerMethod = iota
	ServerPost
	ServerPut
	ServerPatch
	ServerDelete
)

type appServer struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *appServer {
	return &appServer{router}
}

func (s *appServer) Register(method ServerMethod, path string, ctrl transport.Ctrl) {
	switch method {
	case ServerGet:
		s.router.GET(path, ctrl.Execute)
	case ServerPost:
		s.router.POST(path, ctrl.Execute)
	case ServerPut:
		s.router.PUT(path, ctrl.Execute)
	case ServerPatch:
		s.router.PATCH(path, ctrl.Execute)
	case ServerDelete:
		s.router.DELETE(path, ctrl.Execute)
	}
}

func (s *appServer) Run() {
	s.router.Run()
}
