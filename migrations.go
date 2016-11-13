package main

func RunMigrations() {
	userMigrations()
}

func userMigrations() {
	db.AutoMigrate(&User{})
	var count int
	db.Table("users").Count(&count)
	if count == 0 {
		//pass: admin123
		user := User{UserName: "admin", Password: "$2a$10$W.uffOh/uRdeiLhipDGwaOGcKhfV1ZXgLe3H09lIdomrAaFB9KCPu", IsAdmin: true}
		db.Create(&user)
	}
}
