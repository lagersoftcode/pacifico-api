package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetUserAvatarItems(rw http.ResponseWriter, req *http.Request, routeData RouteData) {
	var user User
	existingUser := db.Where(&User{ID: routeData.UserId}).First(&user)
	if existingUser.RecordNotFound() {
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	var avatarItems []AvatarItem
	db.Where("points_required <= ?", user.Stats_TotalScore).Order("points_required desc").Find(&avatarItems)

	response := &struct{ Items []AvatarItem }{Items: avatarItems}
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
