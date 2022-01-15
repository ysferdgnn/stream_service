package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	. "stream_service/utils"

	"github.com/gorilla/mux"
)

func VideoHandler(w http.ResponseWriter, r *http.Request) {

	requestParams := mux.Vars(r)
	contentName := requestParams["name"]

	if IsEmptyString(contentName) {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	filename, isExist := CheckFile(contentName, ".mp4")

	if isExist == false {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "video/mp4; charset=utf-8")
	w.Header().Add("Connection", "Keep-Alive")

	writeContentHttpWriter(w, filename)

}

func TextHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	contentName := requestParams["name"]

	if IsEmptyString(contentName) {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	filename, isExist := CheckFile(contentName, ".txt")

	if isExist == false {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Connection", "Keep-Alive")
	writeContentHttpWriter(w, filename)
}

func MusicHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	contentName := requestParams["name"]

	if IsEmptyString(contentName) {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	filename, isExist := CheckFile(contentName, ".mp3")

	if isExist == false {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "audio/mpeg")
	w.Header().Add("Connection", "Keep-Alive")
	writeContentHttpWriter(w, filename)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	contentName := requestParams["name"]

	if IsEmptyString(contentName) {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	filename, isExist := CheckFile(contentName, ".jpg")

	if isExist == false {
		w.Write([]byte("Not Found!"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "image/jpeg")
	w.Header().Add("Connection", "Keep-Alive")
	writeContentHttpWriter(w, filename)

}

func writeContentHttpWriter(w http.ResponseWriter, contentName string) {

	fileInfo, errorInfo := os.Stat(contentName)

	if errorInfo != nil {
		log.Println("write content de hata var")
		return
	}

	filesize := fileInfo.Size()

	var chunksize int64 = 1000
	var counter int64 = 0
	var bytepivot int64 = 0
	var loopCount int64 = filesize / chunksize
	var exceptionBytes int64 = filesize % chunksize
	bytessend := 0
	log.Default().Printf(fmt.Sprintf("FileSize: %v", filesize))
	log.Default().Printf(fmt.Sprintf("Loop Count->%v", loopCount))
	w.Header().Add("Content-Length", fmt.Sprintf("%v", filesize))

	for counter < loopCount {
		fmt.Println(fmt.Sprintf("Counter-> %v", counter))

		fileBytes := ReadFileWithChunkSize(contentName, chunksize, bytepivot)
		w.Write(fileBytes)

		counter++
		bytepivot = chunksize + bytepivot
		bytessend = int(bytepivot)
		log.Default().Printf(fmt.Sprintf("ChunkNew : %v", bytepivot))
	}

	if exceptionBytes > 0 {
		chunksize = exceptionBytes
		bytessend += int(chunksize)
		fileBytes := ReadFileWithChunkSize(contentName, chunksize, bytepivot)
		w.Write(fileBytes)
	}

	counter = 0
	log.Default().Println("Bytes sent", bytessend)
	log.Default().Println("File Size byte", filesize)
}
