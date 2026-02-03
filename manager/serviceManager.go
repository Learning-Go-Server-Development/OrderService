package manager

import (
	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderService/database"
)

type ServiceManager struct {
	DB               database.Database
	Proxy            px.Proxy
	OrderServiceHost string
}

// func (s *ServiceManager) New() Manager {
// 	return s
// }
