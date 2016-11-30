package main

import "time"

type ScoreTransactionType int

const (
	MedalTransaction  ScoreTransactionType = 1
	TrophyTransaction ScoreTransactionType = 2
	KudoTransaction   ScoreTransactionType = 3
)

const (
	Bronze   string = "Bronze"
	Silver   string = "Silver"
	Gold     string = "Gold"
	Diamond  string = "Diamond"
	Platinum string = "Platinum"
)

type User struct {
	ID       string `gorm:"primary_key;type:char(36)"`
	UserName string `gorm:"unique_index:idx_username;type:varchar(30);"`
	Password string `sql:"type:char(60)"`
	IsLocked bool
	IsAdmin  bool
}

type ScoreTransaction struct {
	ID              string               `gorm:"primary_key;type:char(36)"`
	CreatedAt       time.Time            `gorm:"index:idx_createdAt"`
	TransactionType ScoreTransactionType `gorm:"index:idx_transactionType"`
	UserID          string               `gorm:"index:idx_userId;type:char(36)"`
	ItemDataId      string               `gorm:"type:char(36)"`
	GivenBy         string               `gorm:"type:varchar(30)"`
	Points          uint
}

type Medal struct {
	ID          string `gorm:"primary_key;type:char(36)"`
	Name        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:varchar(50)"`
	Image       string `gorm:"type:varchar(200)"`
	Material    string
	ScoreAmount uint
}

type Trophy struct {
	ID          string `gorm:"primary_key;type:char(36)"`
	Name        string `gorm:"type:varchar(50)"`
	Image       string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:varchar(50)"`
	ScoreAmount uint
}

type UserStatus struct {
	UserId        string `gorm:"primary_key;type:char(36)"`
	TotalMedals   uint
	TotalTrophies uint
	TotalKudos    uint
	TotalScore    uint
}
