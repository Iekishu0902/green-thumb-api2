package entity

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// TMyPlant represents a row from 'public.t_my_plants'.
type TMyPlant struct {
	MyPlantsID     int       `json:"my_plants_id"`    // my_plants_id
	PlantID        int       `json:"plant_id"`        // plant_id
	Purchase       time.Time `json:"purchase"`        // purchase
	UserID         int       `json:"user_id"`         // user_id
	DeleteFlg      bool      `json:"delete_flg"`      // delete_flg
	CreateUser     string    `json:"create_user"`     // create_user
	CreateDatetime time.Time `json:"create_datetime"` // create_datetime
	UpdateUser     string    `json:"update_user"`     // update_user
	UpdateDatetime time.Time `json:"update_datetime"` // update_datetime
	Version        int       `json:"version"`         // version
}