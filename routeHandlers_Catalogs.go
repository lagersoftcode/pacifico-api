package main

import (
	"encoding/json"
	"net/http"

	"github.com/twinj/uuid"
)

func CreateTrophy(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var request CreateTrophyRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	if len(request.Name) > 0 && len(request.Image) > 0 {
		trophy := Trophy{
			ID:          uuid.NewV4().String(),
			Name:        request.Name,
			Description: request.Description,
			Image:       request.Image,
			ScoreAmount: request.ScoreAmount,
		}
		db.Create(&trophy)
	}

	response := Response{http.StatusCreated}
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func GetTrophies(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var trophies []Trophy
	db.Find(&trophies)
	response := GetTrophiesResponse{Trophies: trophies}
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func CreateMedal(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var request CreateMedalRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	if len(request.Name) > 0 && len(request.Image) > 0 {
		medal := Medal{
			ID:          uuid.NewV4().String(),
			Name:        request.Name,
			Description: request.Description,
			Image:       request.Image,
			Material:    MedalMaterial(request.Material),
			ScoreAmount: request.ScoreAmount,
		}
		db.Create(&medal)
	}

	response := Response{http.StatusCreated}
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func GetMedals(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var medals []Medal
	db.Find(&medals)
	response := GetMedalsResponse{Medals: medals}
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}