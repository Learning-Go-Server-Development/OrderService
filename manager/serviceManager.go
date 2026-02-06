package manager

import (
	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderService/database"
	"github.com/Learning-Go-Server-Development/OrderService/delegate"
)

type ServiceManager struct {
	DB               database.Database
	Proxy            px.Proxy
	Delegate         delegate.Delegate
	OrderServiceHost string
}

func (s *ServiceManager) New() Manager {
	return s
}
