package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i *Implementation) Registration(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /register")

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("error: json.NewDecoder: %v", err.Error())
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = i.services.RegisterUser(p.Login, p.Password)
	if err != nil {
		log.Printf("error: i.services.RegisterUser: %v", err.Error())
		http.Error(w, "User already registered", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("got person data: Login: %v, Password: %v", p.Login, p.Password)))

	w.WriteHeader(200)

}
