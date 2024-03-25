package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	// Decimal

	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB

	// Binary

	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
	TiB = 1024 * GiB
	PiB = 1024 * TiB
)

var root string

func main() {
	root = os.Getenv("ROOT_DIR")
	if root == "" {
		root = "./files"
	}
	log.Print("listening at http://localhost:5958 ")
	http.ListenAndServe(":5958", http.HandlerFunc(ReceiveFile))
}
func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseMultipartForm(4 * GB) // limit your max input length!
	log.Printf("upload file %s", r.URL.Path)

	// in your case removeFile would be fileupload
	removeFile, header, err := r.FormFile("file")
	if err != nil {
		log.Print(err)
		return
	}
	defer removeFile.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	p := path.Join(root, r.URL.Path)
	dir := path.Dir(p)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		log.Print(err)
		return
	}

	localFile, err := os.Create(p)
	if err != nil {
		log.Print(err)
		return
	}
	// Copy the file data to my buffer
	_, err = io.Copy(localFile, removeFile)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
