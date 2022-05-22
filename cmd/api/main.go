package main

import (
	"green-thumb-api/controller"
	"net/http"
)

var pl = controller.NewPlantListController()
var ro = controller.NewRouter(pl)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", ro.HandlePlantList)
	server.ListenAndServe()

}
