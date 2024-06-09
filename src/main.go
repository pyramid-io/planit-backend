package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    http.HandleFunc("/", helloHandler)
    
    port := "8080"
    if value, exists := os.LookupEnv("PORT"); exists {
        port = value
    }

    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}