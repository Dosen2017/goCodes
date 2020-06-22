package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"go_dev/video_server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main()  {
	// c := make(chan int)
	go taskrunner.Start()  
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)  //这里就是一个阻塞的函数

	// <- c
}