package main

import (
	"net/http"
	"go_dev/video_server/api/defs"
	"encoding/json"
	"io"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErroResponse) {
	
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error)

	io.WriteString(w, string(resStr))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}