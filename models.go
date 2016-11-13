package main

type User struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"index:idx_username"`
	Password string `sql:"type:char(60)"`
}
