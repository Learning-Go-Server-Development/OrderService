package manager_test

import (
	"testing"

	px "github.com/GolangToolKits/go-http-proxy"
	"github.com/Learning-Go-Server-Development/OrderService/manager"
)

func TestServiceManager_GetCustomer(t *testing.T) {
	var ss manager.ServiceManager

	var gpx px.GoProxy
	ss.Proxy = &gpx
	ss.OrderServiceHost = "http://localhost:3001/rs"

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		phone string
		want  string
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			phone: "154-555-7878",
			want:  "Roberts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			s := ss.New()
			got := s.GetCustomer(tt.phone)
			// TODO: update the condition below to compare got with tt.want.
			if got.LastName != tt.want {
				t.Errorf("GetCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServiceManager_GetCustomerAdresses(t *testing.T) {
	var s manager.ServiceManager

	var gpx px.GoProxy
	s.Proxy = &gpx
	s.OrderServiceHost = "http://localhost:3001/rs"

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cid  int64
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			cid:  12345,
			want: "121 Peachtree St",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			got := s.GetCustomerAdresses(tt.cid)
			// TODO: update the condition below to compare got with tt.want.
			if (*got)[0].Street != tt.want {
				t.Errorf("GetCustomerAdresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
