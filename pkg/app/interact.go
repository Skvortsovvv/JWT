package app

import (
	"fmt"
	"log"

	"JWT/pkg/middleware"
	"net/http"
)

func (i *Implementation) Interact(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /interact")

	userId, err := middleware.GetUserId(r.Context())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error: middleware.GetUserId: %v", err.Error()))
		return
	}

	user, err := i.storage.GetUserByID(userId.String())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error: i.storage.GetUserByID: %v", err.Error()))
		return
	}

	SendSuccessResponse(w, fmt.Sprintf("You successfully made authorized request, %s", user.Login))

	return
}
