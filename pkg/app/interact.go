package app

import "net/http"

func (i *Implementation) Interact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
	w.WriteHeader(200)
	return
}
