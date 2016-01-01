package main 

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Print("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
