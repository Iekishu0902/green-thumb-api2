package main

import (
	"green-thumb-api/config"
	"green-thumb-api/controller"
	"net/http"
)

var pl = controller.NewPlantListController()
var ro = controller.NewRouter(pl)

func main() {
	server := http.Server{
		Addr: ":" + config.Config.Port,
	}
	http.HandleFunc("/", ro.HandlePlantList)
	server.ListenAndServe()

}
