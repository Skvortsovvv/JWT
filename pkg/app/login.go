package app

import (
	"fmt"
	"log"

	"JWT/pkg/model"
	"JWT/pkg/store"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func (i *Implementation) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /login")

	var p model.User

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error: json.NewDecoder: %v", err.Error()))
		return
	}

	token, err := i.services.GetUserToken(p.Login, p.Password)
	if err != nil {
		if errors.Is(err, store.NotFoundError) {
			SendErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("error: i.services.GetUserToken: %v", err.Error()))

			return
		}
		SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error: i.services.GetUserToken: %v", err.Error()))
		return
	}

	SendTokenResponse(w, token)

}
