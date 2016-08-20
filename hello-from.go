package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		hostPort = "8080"
	}
	log.Printf(fmt.Sprintf("Setting host port to %s...\n", hostPort))

	hostName, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s\n", hostName)
	})

	log.Printf(fmt.Sprintf("Go Web server listening on port %s...\n", hostPort))
	log.Fatal(http.ListenAndServe(":"+hostPort, nil))
}
