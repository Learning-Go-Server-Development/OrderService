package database

import (
	"log"
)

func (d *OrderDB) AddItem(i *Item) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if i != nil {
		var a []any
		a = append(a, i.OrderID, i.ProductID)
		log.Println(a)
		suc, id = d.DB.Insert(insertItem, a...)
		log.Println("suc in add blog", suc)
		log.Println("id in add blog", id)
	}
	return suc, id
}
