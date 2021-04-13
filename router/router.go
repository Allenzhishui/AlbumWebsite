package router

import (
	"AlbumWebsite+sql/control"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", control.IndexView)
	http.HandleFunc("/upload", control.UploadView)
	http.HandleFunc("/api/upload",control.ApiUpload)
	http.HandleFunc("/detail",control.DetailView)
	http.HandleFunc("/list",control.ListView)
	http.HandleFunc("/api/list",control.ApiList)
	http.HandleFunc("/api/del",control.ApiDel)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	log.Println("Run 8080 ...")
	http.ListenAndServe(":8080", nil)
}