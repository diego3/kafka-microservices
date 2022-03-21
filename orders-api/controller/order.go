package controller

import (
	"io"
	"net/http"
)

func HandleCreateOrder(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)

	io.WriteString(w, "created")
}
