package main

import (
	"log"
	"net/http"

)

func main(){
	http.HandleFunc("/t1", http.HandlerFunc(RenderTemplate))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
