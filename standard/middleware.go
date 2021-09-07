package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{\"message\": \"noice\"}\n"))
	if err != nil {
		fmt.Println("error handling main", err.Error())
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{\"message\":\"ok\"}\n"))
	if err != nil {
		fmt.Println("error handling healthcheck", err.Error())
	}
}

func main() {
	http.HandleFunc("/", handleMain)

	http.HandleFunc("/healthcheck", handleHealthCheck)

	fmt.Println("server listen at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
