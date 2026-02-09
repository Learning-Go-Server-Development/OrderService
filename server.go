package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	px "github.com/GolangToolKits/go-http-proxy"
	gdb "github.com/GolangToolKits/go-mysql"
	mux "github.com/GolangToolKits/grrt"
	"github.com/Learning-Go-Server-Development/OrderService/database"
	"github.com/Learning-Go-Server-Development/OrderService/delegate"
	"github.com/Learning-Go-Server-Development/OrderService/handlers"
	"github.com/Learning-Go-Server-Development/OrderService/manager"
)

func main() {

	var dbHost string
	var dbUser string
	var dbPassword string

	if os.Getenv("DB_HOST") != "" {
		dbHost = os.Getenv("DB_HOST")
	} else {
		dbHost = "localhost:3306"
	}

	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	} else {
		dbUser = ""
	}

	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	} else {
		dbPassword = ""
	}

	//create database
	mm := &gdb.MyDB{
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		Database: "lgs_orders",
	}
	m := mm.New()
	m.Connect()
	var odb database.OrderDB
	odb.DB = m
	//----------

	//inject database into manager
	var sm manager.ServiceManager
	sm.DB = odb.New()

	//inject proxy
	var prx px.GoProxy
	sm.Proxy = prx.New()

	//inject delegate
	var del delegate.ServiceDelegate
	sm.Delegate = del.New()

	// set OrderService Host URL
	sm.OrderServiceHost = "http://localhost:3001/rs"

	var hh handlers.ServiceHandler
	hh.Manager = sm.New()

	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router := mux.NewRouter()

	h := hh.New()

	// configure routes-------------
	router.HandleFunc("/rs/order/add", h.AddOrder).Methods("POST")
	router.HandleFunc("/rs/order/update", h.UpdateOrder).Methods("PUT")
	router.HandleFunc("/rs/order/get/{id}", h.GetOrder).Methods("GET")
	router.HandleFunc("/rs/orders/current/{cid}", h.GetCurrentOrders).Methods("GET")
	router.HandleFunc("/rs/orders/past/{cid}", h.GetPastOrders).Methods("GET")
	router.HandleFunc("/rs/order/delete/{id}", h.DeleteOrder).Methods("DELETE")

	router.HandleFunc("/rs/item/add", h.AddItem).Methods("POST")
	router.HandleFunc("/rs/item/update", h.UpdateItem).Methods("PUT")
	router.HandleFunc("/rs/items/{oid}", h.GetItems).Methods("GET")
	router.HandleFunc("/rs/item/delete/{id}", h.DeleteItem).Methods("DELETE")

	router.HandleFunc("/rs/customer/get/{phone}", h.GetCustomer).Methods("GET")
	router.HandleFunc("/rs/customer/addresses/get/{cid}", h.GetCustomerAddresses).Methods("GET")

	router.HandleFunc("/rs/product/{sku}", h.GetProduct).Methods("GET")

	//-------------------------------

	fmt.Println("Order Service server is running on port " + port + "!")

	http.ListenAndServe(":"+port, router)

}

//  go mod init github.com/Learning-Go-Server-Development/OrderService
