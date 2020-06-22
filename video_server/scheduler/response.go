package main

import (
	"net/http"
	"io"
)

func SendResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}