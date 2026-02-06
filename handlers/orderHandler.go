package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Learning-Go-Server-Development/OrderService/manager"
)

func (h *ServiceHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var o manager.Order
		ps, err := h.processBody(r, &o)
		log.Println("bs: ", ps)
		log.Println("err: ", err)
		if !ps || err != nil {
			http.Error(w, "Trouble parsing body", http.StatusBadRequest)
		} else {
			or := h.Manager.AddOrder(&o)
			log.Println("oo: ", or)
			if or.Success && or.ID != 0 {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(or)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
