package repository

import (
	"fmt"
	"log"
)

type PlantListRepository interface {
	GetPlantList()
}

type plantListRepository struct{}

func NewPlantListRepository() PlantListRepository {
	return &plantListRepository{}
}

func (pr *plantListRepository) GetPlantList() {
	rows, err := Db.Query("select * from m_plant_cotegory")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rows)

}

// plantList = []
