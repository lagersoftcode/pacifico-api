package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/twinj/uuid"
)

func GiveTrophy(rw http.ResponseWriter, req *http.Request, routeData RouteData) {
	var request GiveTrophyRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	if len(request.TrophyId) > 1 && len(request.UserId) > 1 {
		var user User
		existingUser := db.Where(&User{ID: request.UserId}).First(&user)
		if existingUser.RecordNotFound() {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var trophy Trophy
		existingTrophy := db.Where(&Trophy{ID: request.TrophyId}).First(&trophy)
		if existingTrophy.RecordNotFound() {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var trophyGiven ScoreTransaction
		existingTrans := db.Where(&ScoreTransaction{UserID: request.UserId, ItemDataId: request.TrophyId}).First(&trophyGiven)

		if existingTrans.RecordNotFound() {
			scoreTransaction := ScoreTransaction{
				ID:              uuid.NewV4().String(),
				CreatedAt:       time.Now(),
				UserID:          request.UserId,
				TransactionType: TrophyTransaction,
				ItemDataId:      request.TrophyId,
				GivenBy:         routeData.Username,
				Points:          trophy.ScoreAmount,
			}

			db.Create(&scoreTransaction)

		}
		UpdateUserStats(request.UserId)
		response := Response{http.StatusCreated}
		rw.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}

	} else {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

}