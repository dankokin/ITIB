package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"itib/lab9/handlers"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func openBrowser(url string) {
	time.Sleep(time.Second * 1)
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {


	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(handlers.JsonContentType)

	apiRouter.HandleFunc("/", handlers.ApiHomeHandler)
	apiRouter.HandleFunc("/area", handlers.ApiGetArea).Queries("id", "{id:[0-9]+}", "dist_id", "{dist_id:[0-9]+}").Methods("GET")
	apiRouter.HandleFunc("/area", handlers.AddArea).Methods("POST")
	apiRouter.HandleFunc("/area", handlers.ApiClearArea).Methods("DELETE")
	apiRouter.HandleFunc("/point", handlers.AddPoint).Methods("POST")
	apiRouter.HandleFunc("/cluster", handlers.AddCluster).Methods("POST")
	apiRouter.HandleFunc("/train", handlers.Learn).Methods("POST")

	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))


	go openBrowser("http://127.0.0.1:8080/")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("server not started: %s", err.Error())
	}
}
