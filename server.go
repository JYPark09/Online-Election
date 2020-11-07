package main

import (
	"fmt"
	"log"
	"net/http"
)

func startServer(port string) *http.Server {
	srv := &http.Server{Addr: port}

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/elect_view", electViewHandler)
	http.HandleFunc("/elect", electHandler)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("[http] listen failed ", err)
		}
	}()

	return srv
}

func produceMsg(writer http.ResponseWriter, msg string) {
	fmt.Fprintln(writer, "<script> alert('"+msg+"'); document.location = '/'; </script>")
}
