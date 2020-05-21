// This is the REST API for the Farmerly project
package main

import (
	"github.com/gabielfemi/farmerly_api/api"
	_ "github.com/gabielfemi/go-inform/farmerly"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Index)
	_ = http.ListenAndServe(":9000", nil)
}
