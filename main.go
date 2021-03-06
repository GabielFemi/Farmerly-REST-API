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
	http.HandleFunc("/users", api.Users)
	http.HandleFunc("/posts", api.Posts)
	http.HandleFunc("/posts/fg", api.GetSinglePost)
	log.Println("Listening on 127.0.0.1:9000")
	_ = http.ListenAndServe(":9000", nil)
}
