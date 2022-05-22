package controller

import (
	"fmt"
	"net/http"
)

type PlantListController interface {
}

type plantListController struct {
}

func NewPlantListController() PlantListController {
	return &plantListController{}
}

func (pl *plantListController) PlantList(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sample")
}
