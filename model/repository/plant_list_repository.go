package repository

import (
	"green-thumb-api/model/entity"
	"log"
)

type PlantCategoryListRepository interface {
	GetPlantCategoryList() (plantCategoryList []entity.MPlantCategory, err error)
}

type plantCategoryListRepository struct{}

func NewPlantCategoryListRepository() PlantCategoryListRepository {
	return &plantCategoryListRepository{}
}

func (pr *plantCategoryListRepository) GetPlantCategoryList() (plantCategoryList []entity.MPlantCategory, err error) {
	plantCategoryList = []entity.MPlantCategory{}
	rows, err := Db.Query("select plant_category_name from m_plant_category")

	if err != nil {
		log.Fatalln(err)
		return
	}
	for rows.Next() {
		plant := entity.MPlantCategory{}
		err = rows.Scan(&plant.PlantcategoryName)
		if err != nil {
			log.Fatalln(err)
			return
		}
		plantCategoryList = append(plantCategoryList, plant)
	}

	return

}
