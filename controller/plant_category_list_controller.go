package controller

import (
	"encoding/json"
	"green-thumb-api/controller/response"
	"green-thumb-api/model/repository"
	"net/http"
)

type PlantCategoryListController interface {
	PlantCategoryList(w http.ResponseWriter, r *http.Request)
}

type plantCategoryListController struct {
	plantCategoryListRepository repository.PlantCategoryListRepository
}

func NewPlantCategoryListController(plantCategoryListRepo repository.PlantCategoryListRepository) PlantCategoryListController {
	return &plantCategoryListController{plantCategoryListRepo}
}

func (plantCategoryListController *plantCategoryListController) PlantCategoryList(w http.ResponseWriter, r *http.Request) {
	plantCategoryList, err := plantCategoryListController.plantCategoryListRepository.GetPlantCategoryList()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var plantCategoryResponse []response.PlantCategory
	for _, v := range plantCategoryList {
		plantCategoryResponse = append(plantCategoryResponse, response.PlantCategory{PlantcategoryName: v.PlantcategoryName})
	}

	var plantCategoryListResponse response.PlantCategoryListResponse
	plantCategoryListResponse.PlantCategoryList = plantCategoryResponse

	output, _ := json.MarshalIndent(plantCategoryListResponse.PlantCategoryList, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}
