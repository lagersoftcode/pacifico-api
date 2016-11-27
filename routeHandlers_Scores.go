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
