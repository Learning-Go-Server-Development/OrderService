package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Learning-Go-Server-Development/OrderService/handlers"
	"github.com/Learning-Go-Server-Development/OrderService/manager"
)

func TestServiceHandler_AddItem(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler
	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		res   manager.ResponseID
		code  int
		ctype string
		json  io.ReadCloser
		suc   bool
		want  int64
		want2 bool
	}{
		// TODO: Add test cases.
		{
			code:  200,
			ctype: "application/json",
			json:  io.NopCloser(bytes.NewBufferString(`{"orderId":1, "productId": 5}`)),
			suc:   true,
			want:  1,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.AddItemRes = &manager.ResponseID{
				ID:      tt.want,
				Success: tt.want2,
			}
			// aJSON := io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`))

			r, _ := http.NewRequest("POST", "/ffllist", tt.json)
			r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			hh.AddItem(w, r)

			var res manager.ResponseID
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_UpdateItem(t *testing.T) {
	var mm manager.MockServiceManager
	m := mm.New()
	var hh handlers.ServiceHandler
	hh.Manager = m

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w     http.ResponseWriter
		r     *http.Request
		res   manager.Response
		code  int
		ctype string
		json  io.ReadCloser
		suc   bool
		want  int64
		want2 bool
	}{
		// TODO: Add test cases.
		{
			code:  200,
			ctype: "application/json",
			json:  io.NopCloser(bytes.NewBufferString(`{"id":1, "orderId":1, "productId": 5}`)),
			suc:   true,
			want:  1,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm.UpdateItemRes = &manager.Response{
				Success: tt.want2,
			}
			// aJSON := io.NopCloser(bytes.NewBufferString(`{"orderNumber":"OD-11123", "customerID": 5}`))

			r, _ := http.NewRequest("PUT", "/ffllist", tt.json)
			r.Header.Set("Content-Type", tt.ctype)
			w := httptest.NewRecorder()
			// TODO: construct the receiver type.
			//var h handlers.ServiceHandler
			hh.UpdateItem(w, r)

			var res manager.Response
			body, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if w.Code != tt.code || res.Success != tt.suc {
				t.Fail()
			}
		})
	}
}
