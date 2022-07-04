package main

import (
	"green-thumb-api/config"
	"green-thumb-api/controller"
	"green-thumb-api/model/repository"
	"net/http"
)

var plantCategoryListRepo = repository.NewPlantCategoryListRepository()
var plantCategoryListController = controller.NewPlantCategoryListController(plantCategoryListRepo)
var router = controller.NewRouter(plantCategoryListController)

func main() {
	server := http.Server{
		Addr: ":" + config.Config.Port,
	}
	http.HandleFunc("/", router.HandlePlantCategoryList)
	server.ListenAndServe()

}
