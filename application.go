/* VERY Simple web server to serve static file */

package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fileServer := http.FileServer(http.Dir("../templates/"))
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
