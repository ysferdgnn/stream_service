package main

import (
	"net/http"
	. "stream_service/handlers"

	"github.com/gorilla/mux"
)

func main() {

	gorillaRouter := mux.NewRouter()

	gorillaRouter.HandleFunc("/video", VideoHandler)
	gorillaRouter.HandleFunc("/video/", VideoHandler)
	gorillaRouter.HandleFunc("/video/{name}", VideoHandler)

	gorillaRouter.HandleFunc("/text", TextHandler)
	gorillaRouter.HandleFunc("/text/", TextHandler)
	gorillaRouter.HandleFunc("/text/{name}", TextHandler)

	gorillaRouter.HandleFunc("/music", MusicHandler)
	gorillaRouter.HandleFunc("/music/", MusicHandler)
	gorillaRouter.HandleFunc("/music/{name}", MusicHandler)

	gorillaRouter.HandleFunc("/image", ImageHandler)
	gorillaRouter.HandleFunc("/image/", ImageHandler)
	gorillaRouter.HandleFunc("/image/{name}", ImageHandler)

	http.ListenAndServe("localhost:8080", gorillaRouter)

}
