package database

import (
	"fmt"
	"log"
	"strconv"

	gdb "github.com/GolangToolKits/go-mysql"
)

type OrderDB struct {
	DB gdb.Database
}

func (d *OrderDB) testConnection() bool {
	log.Println("in testConnection")
	var rtn = false
	var a []any
	log.Println("d.DB: ", fmt.Sprintln(d.DB))
	rowPtr := d.DB.Test("select count(*) from orders ", a...)
	log.Println("rowPtr", *rowPtr)
	log.Println("after testConnection test", *rowPtr)
	if len(rowPtr.Row) != 0 {
		foundRow := rowPtr.Row
		int64Val, err := strconv.ParseInt(foundRow[0], 10, 0)
		log.Print("Records found during test ")
		log.Println("Records found during test :", int64Val)
		if err != nil {
			log.Println(err)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

// func (d *OrderDB)AddOrder(o *Order) (bool, int64){

// }
