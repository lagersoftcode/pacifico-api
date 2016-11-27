package main

import "github.com/twinj/uuid"

func RunMigrations() {
	userMigrations()
	scoreMigrations()
}

func userMigrations() {
	db.AutoMigrate(&User{})
	var count int
	db.Table("users").Count(&count)
	if count == 0 {
		//pass: admin123
		user := User{ID: uuid.NewV4().String(), UserName: "admin", Password: "$2a$10$W.uffOh/uRdeiLhipDGwaOGcKhfV1ZXgLe3H09lIdomrAaFB9KCPu", IsAdmin: true}
		db.Create(&user)
	}
}

func scoreMigrations() {
	db.AutoMigrate(&ScoreTransaction{})
	db.AutoMigrate(&Medal{})
	db.AutoMigrate(&Trophy{})
}
