package main

import (
	"log"
	"net/http"
)

func startServer(port string) *http.Server {
	srv := &http.Server{Addr: port}

	http.HandleFunc("/", mainHandler)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("[http] listen failed ", err)
		}

		log.Println("[http] server started")
	}()

	return srv
}
