package infrastructure

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/murilocosta/agartha/internal/infrastructure/transport"
)

type ServerMethod string

const (
	ServerGet     ServerMethod = "GET"
	ServerPost    ServerMethod = "POST"
	ServerPut     ServerMethod = "PUT"
	ServerPatch   ServerMethod = "PATCH"
	ServerDelete  ServerMethod = "DELETE"
	ServerOptions ServerMethod = "OPTIONS"
)

type appServer struct {
	router *gin.Engine
	auth   *jwt.GinJWTMiddleware
}

func NewServer(router *gin.Engine) *appServer {
	return &appServer{router: router, auth: nil}
}

func (s *appServer) ApplyCORS() {
	allowedMethods := []string{
		string(ServerGet),
		string(ServerPost),
		string(ServerPut),
		string(ServerPatch),
		string(ServerDelete),
		string(ServerOptions),
	}

	s.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     allowedMethods,
		AllowHeaders:     []string{"Origin", "Authentication"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func (s *appServer) RegisterAuthHandlers(auth *jwt.GinJWTMiddleware) {
	s.auth = auth
	s.router.POST("/api/login", s.auth.LoginHandler)
	s.router.GET("/api/refresh-token", s.auth.RefreshHandler)
}

func (s *appServer) RegisterCtrlHandlers(config *serverHandlersConfig) {
	for _, handler := range config.Handlers {
		if handler.Secured && s.auth != nil {
			s.router.Use(s.auth.MiddlewareFunc())
			{
				s.registerImpl(handler.HTTPMethod, handler.URLPath, handler.Ctrl)
			}
		} else {
			s.registerImpl(handler.HTTPMethod, handler.URLPath, handler.Ctrl)
		}
	}
}

func (s *appServer) registerImpl(method ServerMethod, path string, ctrl transport.Ctrl) {
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
