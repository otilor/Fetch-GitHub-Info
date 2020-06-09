package main

import (
	"github.com/GabielFemi/Fetch-GitHub-Info/githubInfo"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router:= mux.NewRouter()
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.HandleFunc("/", githubInfo.Index)
	 server := &http.Server{
	 	Addr: "127.0.0.1:8000",
	 	Handler: router,
	 }
	 logrus.Info("Application has started on ", server.Addr)
	 logrus.Fatalln(server.ListenAndServe())
}