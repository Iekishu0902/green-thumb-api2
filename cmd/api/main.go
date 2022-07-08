package main

import (
	"green-thumb-api/config"
	"green-thumb-api/controller"
	"green-thumb-api/model/repository"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":" + config.Config.Port,
	}
	http.HandleFunc("/plant_category/", plantCategoryList)
	server.ListenAndServe()
}

func plantCategoryList(w http.ResponseWriter, r *http.Request) {
	var plantCategoryListRepository = repository.NewPlantCategoryListRepository()
	var plantCategoryListController = controller.NewPlantCategoryListController(plantCategoryListRepository)
	var router = controller.NewRouter(plantCategoryListController)
	router.HandlePlantCategoryList(w, r)
}
