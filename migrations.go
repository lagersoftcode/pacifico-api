package main

func RunMigrations() {
	db.AutoMigrate(&User{})
	db.Model(&User{}).AddIndex("idx_user_name", "user_name")
}
