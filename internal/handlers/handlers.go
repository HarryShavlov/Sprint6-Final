package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/interna/service"
)

func MainHandle(res http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("index.html")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	res.Write(data)
}

func UploadHandle(res http.ResponseWriter, req *http.Request) {

	err := req.ParseMultipartForm(10 << 20)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handle, err := req.FormFile("file")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(handle.Filename)
	fileName := time.Now().UTC().Format("20260504102500") + ext

	outPutFile, err := os.Create(fileName)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write([]byte(result))

	defer outPutFile.Close()
	defer file.Close()
}
