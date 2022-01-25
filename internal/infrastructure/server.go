package infrastructure

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerMethod string

func (m *ServerMethod) ToString() string {
	return string(*m)
}

const (
	ServerGet     ServerMethod = http.MethodGet
	ServerPost    ServerMethod = http.MethodPost
	ServerPut     ServerMethod = http.MethodPut
	ServerPatch   ServerMethod = http.MethodPatch
	ServerDelete  ServerMethod = http.MethodDelete
	ServerOptions ServerMethod = http.MethodOptions
)

type appServer struct {
	router *gin.Engine
	auth   *jwt.GinJWTMiddleware
}

func NewServer(router *gin.Engine) *appServer {
	return &appServer{router: router, auth: nil}
}

func (s *appServer) ApplyCORS() {
	allowedOrigins := []string{
		"http://localhost:3000",
		"https://murilocosta.github.io",
		"https://agartha.stoplight.io",
	}

	allowedMethods := []string{
		string(ServerGet),
		string(ServerPost),
		string(ServerPut),
		string(ServerPatch),
		string(ServerDelete),
		string(ServerOptions),
	}

	allowedHeaders := []string{
		"Origin",
		"Authorization",
		"Authentication",
		"Content-Type",
	}

	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     allowedMethods,
		AllowHeaders:     allowedHeaders,
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Type"},
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
		s.registerImpl(handler)
	}
}

func (s *appServer) registerImpl(handler *ctrlHandler) {
	var handlerChain []gin.HandlerFunc
	if handler.Secured && s.auth != nil {
		handlerChain = append(handlerChain, s.auth.MiddlewareFunc())
	}
	handlerChain = append(handlerChain, handler.Ctrl.Execute)

	s.router.Handle(
		handler.HTTPMethod.ToString(),
		handler.URLPath,
		handlerChain...,
	)
}

func (s *appServer) Run() {
	s.router.Run()
}
