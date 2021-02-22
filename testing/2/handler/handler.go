package handler

import (
	"encoding/json"
	"github.com/fn-code/Go-Example/testing/2/service"
	"log"
	"net/http"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) RegisteUserPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {

		cmd := &service.RegisterUser{}
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(cmd)
		if err != nil {
			log.Printf("error decode requested body : %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if cmd.Email == "" {
			log.Println("error make new user, email is empty")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ok, err := h.svc.RegisterUser(cmd)
		if err != nil {
			log.Printf("error make new user : %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !ok {
			log.Println("cannot adding new user, because user already registered")
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		log.Println("error make request, method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
