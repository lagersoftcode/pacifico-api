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
			createScoreTransaction(user, routeData, TrophyTransaction, trophy.ID, trophy.ScoreAmount)
			UpdateUserStats(request.UserId)
		}

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

func GiveMedal(rw http.ResponseWriter, req *http.Request, routeData RouteData) {
	var request GiveMedalRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	if len(request.MedalId) > 1 && len(request.UserId) > 1 {
		var user User
		existingUser := db.Where(&User{ID: request.UserId}).First(&user)
		if existingUser.RecordNotFound() {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var medal Medal
		existingMedal := db.Where(&Medal{ID: request.MedalId}).First(&medal)
		if existingMedal.RecordNotFound() {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var medalGiven ScoreTransaction
		existingTrans := db.Where(&ScoreTransaction{UserID: request.UserId, ItemDataId: request.MedalId}).First(&medalGiven)

		if existingTrans.RecordNotFound() {
			createScoreTransaction(user, routeData, MedalTransaction, medal.ID, medal.ScoreAmount)
			UpdateUserStats(request.UserId)
		}

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

func GiveKudo(rw http.ResponseWriter, req *http.Request, routeData RouteData) {
	var request GiveKudoRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	if len(request.UserId) > 1 {
		var user User
		existingUser := db.Where(&User{ID: request.UserId}).First(&user)
		if existingUser.RecordNotFound() {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if user.ID == routeData.UserId {
			rw.WriteHeader(http.StatusConflict)
			return
		}

		thirtyMinutesAgo := time.Now().Add(-30 * time.Minute)
		existingTrans := db.Where("created_at > ? AND user_id = ? AND given_by_id = ?", thirtyMinutesAgo, user.ID, routeData.UserId).First(&ScoreTransaction{})

		if existingTrans.RecordNotFound() {
			createScoreTransaction(user, routeData, KudoTransaction, "", 20)
			UpdateUserStats(request.UserId)
			rw.WriteHeader(http.StatusCreated)
		} else {
			rw.WriteHeader(http.StatusNotAcceptable)
		}

		response := Response{http.StatusCreated}

		if err := json.NewEncoder(rw).Encode(response); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

	} else {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func createScoreTransaction(user User, routeData RouteData, transactionType string, itemId string, points uint) {
	scoreTransaction := ScoreTransaction{
		ID:              uuid.NewV4().String(),
		CreatedAt:       time.Now(),
		UserID:          user.ID,
		UserName:        user.UserName,
		TransactionType: transactionType,
		ItemDataId:      itemId,
		GivenBy:         routeData.Username,
		GivenById:       routeData.UserId,
		Points:          points,
	}
	db.Create(&scoreTransaction)
}

func GetLastActions(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var lastActions []ScoreTransaction
	db.Limit(5).Order("created_at desc").Find(&lastActions)

	var publicActions []PublicAction
	for _, transaction := range lastActions {
		publicActions = append(publicActions,
			PublicAction{
				CreatedAt:  transaction.CreatedAt,
				SourceUser: transaction.GivenBy,
				TargetUser: transaction.UserName,
				Item:       transaction.TransactionType})
	}

	response := GetLastActionsResponse{LastActions: publicActions}
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}
