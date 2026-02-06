package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Learning-Go-Server-Development/OrderService/manager"
)

func (h *ServiceHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	Ok := h.checkContent(r)
	if !Ok {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var i manager.Item
		pi, err := h.processBody(r, &i)
		log.Println("bs: ", pi)
		log.Println("err: ", err)
		if !pi || err != nil {
			http.Error(w, "Trouble parsing body", http.StatusBadRequest)
		} else {
			ir := h.Manager.AddItem(&i)
			log.Println("oo: ", ir)
			if ir.Success && ir.ID != 0 {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(ir)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func (h *ServiceHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	Ok := h.checkContent(r)
	if !Ok {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var i manager.Item
		pi, err := h.processBody(r, &i)
		log.Println("bs: ", pi)
		log.Println("err: ", err)
		if !pi || err != nil {
			http.Error(w, "Trouble parsing body", http.StatusBadRequest)
		} else {
			ir := h.Manager.UpdateItem(&i)
			log.Println("oo: ", ir)
			if ir.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(ir)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
