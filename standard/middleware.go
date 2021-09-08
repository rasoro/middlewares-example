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
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(resJSON)
	if err != nil {
		fmt.Println("error handling main", err.Error())
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"message": "ok"}
	resJson, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(resJson)
	if err != nil {
		fmt.Println("error handling healthcheck", err.Error())
	}
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/healthcheck" {
			h.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || user != "admin" || pass != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "Unauthorized"}`)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		h.ServeHTTP(w, r)
	}
}

func applicationJSON(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", basicAuth(applicationJSON(handleMain)))

	http.HandleFunc("/healthcheck", applicationJSON(handleHealthCheck))

	fmt.Println("server listen at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
