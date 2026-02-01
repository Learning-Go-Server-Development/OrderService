package database

type Database interface {
	AddOrder(o *Order) (bool, int64)
	UpdateOrder(o *Order) bool
	GetOrder(cid int64) *Order
	GetAllOrders(cid int64) *[]Order
	DeleteOrder(id int64) bool
}
