package controller

import (
	"fmt"
	"green-thumb-api/model/repository"
	"net/http"
)

type PlantListController interface {
	PlantList(w http.ResponseWriter, r *http.Request)
}

type plantListController struct {
	pr repository.PlantListRepository
}

func NewPlantListController(plantListRepo repository.PlantListRepository) PlantListController {
	return &plantListController{plantListRepo}
}

func (pl *plantListController) PlantList(w http.ResponseWriter, r *http.Request) {
	plantList, err := pl.pr.GetPlantList()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Print(plantList)
}
