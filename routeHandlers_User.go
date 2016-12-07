package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/twinj/uuid"

	"golang.org/x/crypto/bcrypt"
)

func Login(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var request LoginRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		log.Println(parseErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	var response LoginResponse
	var user User
	db.Where(&User{UserName: request.Username}).First(&user)

	if len(user.ID) < 1 {
		response.Success = false
		response.Message = "User not found"
		rw.WriteHeader(http.StatusUnauthorized)
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			response.Success = false
			response.Message = "Invalid credentials"
			rw.WriteHeader(http.StatusUnauthorized)
		} else {
			response.Success = true
			response.AuthToken = GetToken(rw, user.UserName, user.ID, user.IsAdmin)
			rw.WriteHeader(http.StatusOK)
		}
	}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func CreateUser(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var request CreateUserRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		log.Println(parseErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	if len(request.Username) > 0 && len(request.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err == nil {
			user := User{
				ID:        uuid.NewV4().String(),
				UserName:  request.Username,
				Password:  string(hashedPassword),
				FirstName: request.FirstName,
				LastName:  request.LastName,
				Email:     request.Email,
				Category:  request.Category,
				IsLocked:  false,
				IsAdmin:   false,
			}
			db.Create(&user)
		}
	}

	response := CreateUserResponse{Response{http.StatusCreated}}
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetUsers(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	resultsPerPage := 6
	queryUser := req.URL.Query().Get("userSearch")
	queryPage := req.URL.Query().Get("page")
	page, err := strconv.Atoi(queryPage)

	if err != nil {
		page = 0
	} else {
		if page > 0 {
			page = page - 1
		}
	}

	var users []User
	var publicUsers []PublicUser
	db.Offset(page*resultsPerPage).Limit(resultsPerPage).Where("user_name like ?", "%"+queryUser+"%").Find(&users).Scan(&publicUsers) //possible sql injection?

	var totalUsers int
	db.Model(&User{}).Where("user_name like ?", "%"+queryUser+"%").Count(&totalUsers)
	totalPages := int(math.Ceil(float64(totalUsers) / float64(resultsPerPage)))

	response := GetUsersResponse{Users: publicUsers, CurrentPage: page + 1, TotalPages: totalPages}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetUserTrophies(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	queryUserId := req.URL.Query().Get("userId")
	if queryUserId == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User
	existingUser := db.Where(&User{ID: queryUserId}).First(&user)
	if existingUser.RecordNotFound() {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var transactions []UserTrophy
	db.Table("score_transactions").
		Where("score_transactions.user_id = ? and transaction_type = 'Trophy'", queryUserId).
		Select(`score_transactions.created_at, score_transactions.given_by, score_transactions.user_name, score_transactions.transaction_type,
			score_transactions.points, trophies.image, trophies.name, trophies.description`).
		Joins("inner join trophies on score_transactions.item_data_id = trophies.id").
		Order("created_at desc").
		Scan(&transactions)

	response := GetUserTrophiesResponse{Transactions: transactions}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetUserMedals(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	queryUserId := req.URL.Query().Get("userId")
	if queryUserId == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User
	existingUser := db.Where(&User{ID: queryUserId}).First(&user)
	if existingUser.RecordNotFound() {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var transactions []UserMedal
	db.Table("score_transactions").
		Where("score_transactions.user_id = ? and transaction_type = 'Medal'", queryUserId).
		Select(`score_transactions.created_at, score_transactions.given_by, score_transactions.user_name, score_transactions.transaction_type,
			score_transactions.points, medals.image, medals.name, medals.description, medals.material`).
		Joins("inner join medals on score_transactions.item_data_id = medals.id").
		Order("created_at desc").
		Scan(&transactions)

	response := GetUserMedalsResponse{Transactions: transactions}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetUserLastKudos(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	queryUserId := req.URL.Query().Get("userId")
	if queryUserId == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User
	existingUser := db.Where(&User{ID: queryUserId}).First(&user)
	if existingUser.RecordNotFound() {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var transactions []UserKudo
	db.Limit(10).Table("score_transactions").
		Where("score_transactions.user_id = ? and transaction_type = 'Kudo'", queryUserId).
		Select(`score_transactions.created_at, score_transactions.given_by, score_transactions.user_name, score_transactions.transaction_type,
			score_transactions.points`).
		Order("created_at desc").
		Scan(&transactions)

	response := GetUserLastKudosResponse{Transactions: transactions}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetUserStats(rw http.ResponseWriter, req *http.Request, routeData RouteData) {
	queryUserId := req.URL.Query().Get("userId")
	if queryUserId == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var publicUser PublicUser
	var user User
	existingUser := db.Where(&User{ID: queryUserId}).First(&user).Scan(&publicUser)
	if existingUser.RecordNotFound() {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(rw).Encode(publicUser); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
