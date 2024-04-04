package app

import (
	"fmt"
	"log"

	"JWT/pkg/model"
	"encoding/json"
	"net/http"
)

func (i *Implementation) Registration(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /register")

	var p model.User

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error: json.NewDecoder: %v", err.Error()))
		return
	}

	err = i.services.RegisterUser(p.Login, p.Password)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error: i.services.RegisterUser: %v", err.Error()))
		return
	}

	SendSuccessResponse(w, "User registered successfully!")
}
