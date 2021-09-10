package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "noice"}
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(resJSON)
	if err != nil {
		fmt.Printf("error in handleMain: %s", err.Error())
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "ok"}
	resJson, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(resJson)
	if err != nil {
		fmt.Printf("error in handleMain: %s", err.Error())
	}
}

func main() {
	n := negroni.Classic()

	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	r.HandleFunc("/", handleMain).Methods("GET")

	r.HandleFunc("/healthcheck", handleHealth).Methods("GET")

	fmt.Println("server listen at 8090")
	err := http.ListenAndServe(":8090", n)
	if err != nil {
		fmt.Println(err)
	}
}
