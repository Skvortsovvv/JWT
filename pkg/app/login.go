package app

import (
	"JWT/pkg/store"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func (i *Implementation) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /login")

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("error: json.NewDecoder: %v", err.Error())
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := i.services.GetUserToken(p.Login, p.Password)
	if err != nil {
		if errors.Is(err, store.NotFoundError) {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		log.Printf("error: i.services.GetUserToken: %v", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(TokenResponse{Token: token})

	if err != nil {
		log.Fatalf("json.NewEncoder(w).Encode: %v", err)
	}
}
