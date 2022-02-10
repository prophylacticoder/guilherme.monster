/* VERY Simple web server to serve static file */

package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("../templates/"))
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
