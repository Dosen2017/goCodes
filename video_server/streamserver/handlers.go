package main

import (
	"io"
	"os"
	"net/http"
	"io/ioutil"
	"html/template"
	"github.com/julienschmidt/httprouter"
	// "time"

	"log"

)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}


func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// vid := p.ByName("vid_id")
	// vl := ViDEO_DIR + vid + ".mp4"

	// log.Println(ViDEO_DIR + vid + ".mp4")

	// video, err := os.Open(vl)
	// if err != nil {
	// 	log.Printf("Error when try to open file %v", err)
	// 	sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
	// }

	// defer video.Close()

	// w.Header().Set("Content-Type", "video/mp4")
	// http.ServeContent(w, r, "", time.Now(), video)

	log.Println("Entered the streamHandler" + p.ByName("vid_id") + ".mp4")

	targetUrl := "http://dosen-video-t.oss-cn-zhangjiakou.aliyuncs.com/videos/" + p.ByName("vid_id") + ".mp4"

	log.Println(targetUrl)
	
	http.Redirect(w, r, targetUrl, 301)


}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE) ; err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file") // 
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error1")
	}
	data , err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file rror")
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error2")
	}

	fn := p.ByName("vid_id")
	err = ioutil.WriteFile(ViDEO_DIR + fn + ".mp4", data, 0666)
	if err != nil {
		log.Printf("Write file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error3")
		return 
	}

	ossfn := "videos/" + fn + ".mp4"
	path := ViDEO_DIR + fn + ".mp4"
	bn := "dosen-video-t"
	ret := UploadToOss(ossfn, path, bn)
	if !ret {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload Successfully")
}

