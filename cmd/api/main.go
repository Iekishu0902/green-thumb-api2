package main

import (
	"green-thumb-api/config"
	"green-thumb-api/controller"
	"green-thumb-api/model/repository"
	"net/http"
)

var plantListRepo = repository.NewPlantListRepository()
var pl = controller.NewPlantListController(plantListRepo)
var ro = controller.NewRouter(pl)

func main() {
	server := http.Server{
		Addr: ":" + config.Config.Port,
	}
	http.HandleFunc("/", ro.HandlePlantList)
	server.ListenAndServe()

}
