// This is the REST API for the Farmerly project
package main

import (
	"github.com/gabielfemi/farmerly_api/api"
	"github.com/gabielfemi/gabbymux"
	_ "github.com/gabielfemi/go-inform/farmerly"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Index)

	port := "8080"
	portIsTaken := gabbymux.IsPortTaken(port)

	if portIsTaken{
		log.Fatal("Port is taken. Try another port")
	} else {
		log.Println("Listening on port :", port)
		_ = http.ListenAndServe(port, nil)
	}
}
