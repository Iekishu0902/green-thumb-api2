package repository

import (
	"green-thumb-api/model/entity"
	"log"
)

type PlantListRepository interface {
	GetPlantList() (plantList []entity.MPlantCategory, err error)
}

type plantListRepository struct{}

func NewPlantListRepository() PlantListRepository {
	return &plantListRepository{}
}

func (pr *plantListRepository) GetPlantList() (plantList []entity.MPlantCategory, err error) {
	plantList = []entity.MPlantCategory{}
	rows, err := Db.Query("select plant_category_name from m_plant_category")

	if err != nil {
		log.Fatalln(err)
		return
	}
	for rows.Next() {
		plant := entity.MPlantCategory{}
		err = rows.Scan(&plant.PlantcategoryName)
		if err != nil {
			log.Print(err)
			return
		}
		plantList = append(plantList, plant)
	}

	return

}

// plantList = []
