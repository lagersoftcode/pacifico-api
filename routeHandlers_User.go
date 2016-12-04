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
		panic(parseErr)
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
		panic(err)
	}
}

func CreateUser(rw http.ResponseWriter, req *http.Request, routeData RouteData) {

	var request CreateUserRequest
	parseErr := json.Unmarshal(routeData.Body, &request)
	if parseErr != nil {
		panic(parseErr)
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
				Email:     request.LastName,
				IsLocked:  false,
				IsAdmin:   false,
			}
			db.Create(&user)
		}
	}

	response := CreateUserResponse{Response{http.StatusCreated}}
	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
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
