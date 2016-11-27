package main

import "time"

type ScoreTransactionType int

const (
	MedalTransaction  ScoreTransactionType = 1
	TrophyTransaction ScoreTransactionType = 2
	KudoTransaction   ScoreTransactionType = 3
	BossTransaction   ScoreTransactionType = 4
)

type MedalMaterial int

const (
	Bronze   MedalMaterial = 1
	Silver   MedalMaterial = 2
	Gold     MedalMaterial = 3
	Diamond  MedalMaterial = 4
	Platinum MedalMaterial = 5
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
	Description     string               `gorm:"type:varchar(50)"`
	ItemDataId      string               `gorm:"type:char(36)"`
	GivenBy         string               `gorm:"type:char(36)"`
	ScoreAmount     uint
}

type Medal struct {
	ID          string `gorm:"primary_key;type:char(36)"`
	Name        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:varchar(50)"`
	Material    MedalMaterial
	ScoreAmount uint
}

type Trophy struct {
	ID          string `gorm:"primary_key;type:char(36)"`
	Name        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:varchar(50)"`
	ScoreAmount uint
}
