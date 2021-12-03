package infrastructure

import "github.com/murilocosta/agartha/internal/infrastructure/transport"

type ctrlHandler struct {
	HTTPMethod ServerMethod
	URLPath    string
	Ctrl       transport.Ctrl
	Secured    bool
}

type serverHandlersConfig struct {
	Handlers []*ctrlHandler
}

func NewServerHandlersConfig() *serverHandlersConfig {
	return &serverHandlersConfig{}
}

func (c *serverHandlersConfig) AddHandler(method ServerMethod, path string, ctrl transport.Ctrl) {
	c.addHandlerImpl(method, path, ctrl, false)
}

func (c *serverHandlersConfig) AddSecuredHandler(method ServerMethod, path string, ctrl transport.Ctrl) {
	c.addHandlerImpl(method, path, ctrl, true)
}

func (c *serverHandlersConfig) addHandlerImpl(method ServerMethod, path string, ctrl transport.Ctrl, secured bool) {
	c.Handlers = append(c.Handlers, &ctrlHandler{
		HTTPMethod: method,
		URLPath:    path,
		Ctrl:       ctrl,
		Secured:    secured,
	})
}
