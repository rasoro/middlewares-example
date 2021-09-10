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

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}

func main() {
	n := negroni.Classic()
	n.Use(applicationJSON())

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
