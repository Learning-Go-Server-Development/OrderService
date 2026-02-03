package database_test

import (
	"testing"

	gdb "github.com/GolangToolKits/go-mysql"
	"github.com/Learning-Go-Server-Development/OrderService/database"
)

func TestOrderDB_AddItem(t *testing.T) {
	// mm := &gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "",
	// 	Password: "",
	// 	Database: "lgs_orders",
	// }
	// m := mm.New()
	mm := &gdb.MyDBMock{}
	mm.MockTestRow = &gdb.DbRow{
		Row: []string{"0"},
	}
	mm.MockConnectSuccess = true
	mm.MockUpdateSuccess1 = true
	m := mm.New()
	m.Connect()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		i     *database.Item
		want  bool
		want2 int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var d database.OrderDB
			got, got2 := d.AddItem(tt.i)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("AddItem() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("AddItem() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
