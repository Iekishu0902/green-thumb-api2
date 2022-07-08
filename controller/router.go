package controller

import (
	"net/http"
)

type Router interface {
	HandlePlantCategoryList(w http.ResponseWriter, r *http.Request)
}

type router struct {
	pl PlantCategoryListController
}

func NewRouter(pl PlantCategoryListController) Router {
	return &router{pl}
}

func (ro *router) HandlePlantCategoryList(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/plant_category/list":
		ro.pl.PlantCategoryList(w, r)
	}

}
