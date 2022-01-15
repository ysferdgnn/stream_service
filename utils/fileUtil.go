package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {

	if IsEmptyString(filePath) {
		return nil, errors.New("File path can not be empty")
	}

	file, err := ioutil.ReadFile(filePath)

	return file, err
}

func ReadFileWithChunkSize(filePath string, chunksize, startbyte int64) []byte {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 666)
	var fileSize int64 = 0
	fileBytes := make([]byte, chunksize)
	defer file.Close()
	if err != nil {
		log.Println("Dosya Açılamadı!")
	}
	info, errInfo := file.Stat()

	if errInfo != nil {
		log.Println("Dosya okunamadı!")
	}

	fileSize = int64(info.Size())

	if chunksize+startbyte <= fileSize {
		file.ReadAt(fileBytes, int64(startbyte))
	} else {
		fileBytes := make([]byte, fileSize-startbyte)
		file.ReadAt(fileBytes, startbyte)
	}
	return fileBytes
}

func CheckFile(contentName, contentType string) (string, bool) {
	isExist := true

	filename := "files/" + contentName
	_, errorStat := os.Stat(filename)

	if errorStat != nil {
		log.Println("File not found as ", filename)
		isExist = false

	}
	if isExist == false {
		filename = filename + contentType
		_, errorStatExtensionFileName := os.Stat(filename)
		if errorStatExtensionFileName != nil {
			if os.IsNotExist(errorStatExtensionFileName) {
				log.Println("File not found!!")
				return "", false
			}

		}
	}
	return filename, true
}
