package main

type User struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"unique_index:idx_username;type:varchar(30);"`
	Password string `sql:"type:char(60)"`
	IsLocked bool
	IsAdmin  bool
}
