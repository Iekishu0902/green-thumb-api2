package entity

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// TVitalizerHistory represents a row from 'public.t_vitalizer_history'.
type TVitalizerHistory struct {
	VitalizerHistoryID int       `json:"vitalizer_history_id"` // vitalizer_history_id
	PlantID            int       `json:"plant_id"`             // plant_id
	VitalizerID        int       `json:"vitalizer_id"`         // vitalizer_id
	UserID             int       `json:"user_id"`              // user_id
	DeleteFlg          bool      `json:"delete_flg"`           // delete_flg
	CreateUser         string    `json:"create_user"`          // create_user
	CreateDatetime     time.Time `json:"create_datetime"`      // create_datetime
	UpdateUser         string    `json:"update_user"`          // update_user
	UpdateDatetime     time.Time `json:"update_datetime"`      // update_datetime
	Version            int       `json:"version"`              // version
}
