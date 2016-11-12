package main

func RunMigrations() {
	db.AutoMigrate(&User{})
}
