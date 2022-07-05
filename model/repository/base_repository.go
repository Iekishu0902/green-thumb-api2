package repository

import (
	"database/sql"
	"fmt"
	"green-thumb-api/config"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	fmt.Println("base_repository.init()")
	LoadDataSource()
}

func LoadDataSource() {
	var err error
	dataSourceName := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=disable",
		config.Config.User, config.Config.DbName, config.Config.Password, config.Config.DbPort,
	)

	Db, err = sql.Open(config.Config.SQLDriver, dataSourceName)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
