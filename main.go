package main

import (
	"avoxi-interview/api"
	ip_db "avoxi-interview/ip-db"
	"avoxi-interview/utility"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

import "github.com/gorilla/mux"

func main() {
	//config
	config, err := utility.LoadConfig("")
	if err != nil {
		log.Panic("Cannot load configuration!!")
	}

	//logging
	cleanup, err := utility.InitLogging()
	if err != nil {
		log.Panic("Cannot load configuration!!")
	}
	defer cleanup()

	//ip-db
	err = ip_db.InitDb()
	if err != nil {
		log.Panic("Cannot init DB")
	}

	//router
	r := mux.NewRouter()
	r.HandleFunc("/ipcountrycheck", api.IpCountryCheck)
	http.Handle("/ipcountrycheck", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.ListeningAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
