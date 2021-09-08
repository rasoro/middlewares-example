package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	n := negroni.Classic()

	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	fmt.Println("server listen at 8090")
	err := http.ListenAndServe(":8090", n)
	if err != nil {
		fmt.Println(err)
	}
}
