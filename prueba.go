package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := "4444"

	fmt.Println(os.Args[0])

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	http.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Minute * 1)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
