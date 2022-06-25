package controller

import (
	"fmt"
	"net/http"
)

type Router interface {
	HandlePlantList(w http.ResponseWriter, r *http.Request)
}

type router struct {
	pl PlantListController
}

func NewRouter(pl PlantListController) Router {
	return &router{pl}
}

func (ro *router) HandlePlantList(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		fmt.Println("HandlePlantList()")
		ro.pl.PlantList(w, r)
	}

}
