package main

import (
	"fmt"
	"log"
	"net/http"
	"secureai-example-pg/graphql"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Worldkkk")
}

func main() {
	log.Println("Start service")
	http.HandleFunc("/demo-pg/hello", helloWorld)
	http.HandleFunc("/demo-pg/graphql", graphql.Graphql)
	http.ListenAndServe(":6552", nil)
}
