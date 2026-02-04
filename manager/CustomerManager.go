package manager

import (
	"log"
	"net/http"
	"strconv"
)

func (s *ServiceManager) GetCustomer(phone string) *Customer {
	var rtn Customer
	if phone != "" {
		var pcus ProxyCustomer
		req, err := http.NewRequest(http.MethodGet, s.OrderServiceHost+"/customer/get/"+phone, nil)
		if err == nil {
			suc, stat := s.Proxy.Do(req, &pcus)
			log.Println("suc: ", suc)
			log.Println("stat: ", stat)
			if suc && stat == http.StatusOK {
				rtn.ID = pcus.ID
				rtn.FirstName = pcus.FirstName
				rtn.LastName = pcus.LastName
				rtn.PhoneNumber = pcus.PhoneNumber
			}
		}
	}

	return &rtn
}

func (s *ServiceManager) GetCustomerAdresses(cid int64) *[]Address {
	var rtn = []Address{}
	if cid != 0 {
		var pads []ProxyAddress
		scid := strconv.FormatInt(cid, 10)
		req, err := http.NewRequest(http.MethodGet, s.OrderServiceHost+"/addresses/get/"+scid, nil)
		if err == nil {
			suc, stat := s.Proxy.Do(req, &pads)
			log.Println("suc: ", suc)
			log.Println("stat: ", stat)
			if suc && stat == http.StatusOK {
				for _, pa := range pads {
					var a Address
					a.ID = pa.ID
					a.CID = pa.CID
					a.Street = pa.Street
					a.City = pa.City
					a.State = pa.State
					a.ZipCode = pa.ZipCode
					rtn = append(rtn, a)
				}
			}
		}
	}
	return &rtn
}
