package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(rw http.ResponseWriter, req *http.Request) {

	body, readErr := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if readErr != nil {
		panic(readErr)
	}

	var request LoginRequest
	parseErr := json.Unmarshal(body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	var response LoginResponse
	var user User
	db.Where(&User{UserName: request.Username}).First(&user)

	if user.ID == 0 {
		response.Success = false
		response.Message = "Invalid credentials"
		rw.WriteHeader(http.StatusUnauthorized)
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			response.Success = false
			response.Message = "Invalid credentials"
			rw.WriteHeader(http.StatusUnauthorized)
		} else {
			response.Success = true
			SetToken(rw, user.UserName, user.IsAdmin)
			rw.WriteHeader(http.StatusOK)
		}
	}

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func CreateUser(rw http.ResponseWriter, req *http.Request) {

	body, readErr := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if readErr != nil {
		panic(readErr)
	}

	var request CreateUserRequest
	parseErr := json.Unmarshal(body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	if len(request.Username) > 0 && len(request.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err == nil {
			user := User{
				UserName: request.Username,
				Password: string(hashedPassword),
				IsLocked: false,
				IsAdmin:  false,
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
