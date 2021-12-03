package infrastructure

import (
	"github.com/murilocosta/agartha/internal/infrastructure/transport"
)

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

func (c *serverHandlersConfig) Add(method ServerMethod, path string, ctrl transport.Ctrl) {
	c.addHandlerImpl(method, path, ctrl, false)
}

func (c *serverHandlersConfig) Get(path string, ctrl transport.Ctrl) {
	c.Add(ServerGet, path, ctrl)
}

func (c *serverHandlersConfig) Post(path string, ctrl transport.Ctrl) {
	c.Add(ServerPost, path, ctrl)
}

func (c *serverHandlersConfig) Put(path string, ctrl transport.Ctrl) {
	c.Add(ServerPut, path, ctrl)
}

func (c *serverHandlersConfig) Delete(path string, ctrl transport.Ctrl) {
	c.Add(ServerDelete, path, ctrl)
}

func (c *serverHandlersConfig) AddProtected(method ServerMethod, path string, ctrl transport.Ctrl) {
	c.addHandlerImpl(method, path, ctrl, true)
}

func (c *serverHandlersConfig) GetProtected(path string, ctrl transport.Ctrl) {
	c.AddProtected(ServerGet, path, ctrl)
}

func (c *serverHandlersConfig) PostProtected(path string, ctrl transport.Ctrl) {
	c.AddProtected(ServerPost, path, ctrl)
}

func (c *serverHandlersConfig) PutProtected(path string, ctrl transport.Ctrl) {
	c.AddProtected(ServerPut, path, ctrl)
}

func (c *serverHandlersConfig) DeleteProtected(path string, ctrl transport.Ctrl) {
	c.AddProtected(ServerDelete, path, ctrl)
}

func (c *serverHandlersConfig) addHandlerImpl(method ServerMethod, path string, ctrl transport.Ctrl, secured bool) {
	c.Handlers = append(c.Handlers, &ctrlHandler{
		HTTPMethod: method,
		URLPath:    path,
		Ctrl:       ctrl,
		Secured:    secured,
	})
}
