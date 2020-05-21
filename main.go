// This is the REST API for the Farmerly project
package main

import (
	"github.com/gabielfemi/farmerly_api/api"
	_ "github.com/gabielfemi/go-inform/farmerly"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/categories", api.Index)
	log.Println("Listening on 127.0.0.1:9000")
	_ = http.ListenAndServe(":9000", nil)
}
