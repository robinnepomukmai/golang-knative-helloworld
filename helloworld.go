package main

import (
	"net/http"
	"log"
	"fmt"
)

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	log.Print("recieved helloworld request")
	fmt.Fprintf(w, "hello world :)")
}