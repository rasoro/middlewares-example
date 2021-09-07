package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "noice"}
	resJSON, _ := json.Marshal(res)
	_, err := w.Write(resJSON)
	if err != nil {
		fmt.Println("error handling main", err.Error())
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "ok"}
	resJson, _ := json.Marshal(res)
	_, err := w.Write(resJson)
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
