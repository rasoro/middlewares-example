package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()

	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := map[string]string{"message": "noice"}
		resJSON, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(resJSON)
		if err != nil {
			fmt.Println("error in handleMain", err.Error())
		}
	}).Methods("GET")

	fmt.Println("server listen at 8090")
	err := http.ListenAndServe(":8090", n)
	if err != nil {
		fmt.Println(err)
	}
}
