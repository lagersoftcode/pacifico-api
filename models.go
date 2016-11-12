package main

type User struct {
	UserName string `gorm:"index:idx_username"`
}
