package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("{\"message\": \"noice\"}\n"))
		if err != nil {
			fmt.Println("error handling main", err.Error())
		}
	})

	fmt.Println("server listen at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
