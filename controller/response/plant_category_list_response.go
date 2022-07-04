package response

type PlantCategory struct {
	PlantcategoryId   int    `json:"plant_category_id"`
	PlantcategoryName string `json:"plant_category_name"`
}

type PlantCategoryListResponse struct {
	PlantCategoryList []PlantCategory `json:"plantCategoryList"`
}
