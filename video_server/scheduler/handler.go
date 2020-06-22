package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"go_dev/video_server/scheduler/dbops"
)

func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	vid := p.ByName("vid-id")

	if  len(vid) == 0 {
		SendResponse(w, "video id should not be empty", 400)
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)

	if err != nil {
		SendResponse(w, "Internal server error", 500)
		return
	}
	SendResponse(w, "", 200)
}