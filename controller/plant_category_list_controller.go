package controller

import (
	"encoding/json"
	"green-thumb-api/controller/response"
	"green-thumb-api/model/repository"
	"net/http"
)

// 外部に公開するインターフェース
type PlantCategoryListController interface {
	PlantCategoryList(w http.ResponseWriter, r *http.Request)
}

// 非公開のPlantCategoryListControllerの構造体
type plantCategoryListController struct {
	plantCategoryListRepository repository.PlantCategoryListRepository
}

// PlantCategoryListControllerのコンストラクタ
func NewPlantCategoryListController(plantCategoryListRepo repository.PlantCategoryListRepository) PlantCategoryListController {
	return &plantCategoryListController{plantCategoryListRepo}
}

// PlantNameのリストを取得する
func (plantCategoryListController *plantCategoryListController) PlantCategoryList(w http.ResponseWriter, r *http.Request) {
	// リポジトリ呼び出し
	plantCategoryList, err := plantCategoryListController.plantCategoryListRepository.GetPlantCategoryList()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// 取得したentityをresponseオブジェクトに格納
	var plantCategoryResponse []response.PlantCategory
	for _, v := range plantCategoryList {
		plantCategoryResponse = append(plantCategoryResponse, response.PlantCategory{PlantcategoryName: v.PlantcategoryName})
	}

	var plantCategoryListResponse response.PlantCategoryListResponse
	plantCategoryListResponse.PlantCategoryList = plantCategoryResponse

	// jsonに変換
	output, _ := json.MarshalIndent(plantCategoryListResponse.PlantCategoryList, "", "\t\t")

	// jsonを返却
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}
