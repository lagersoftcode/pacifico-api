package main

import "time"

const (
	MedalTransaction  = "Medal"
	TrophyTransaction = "Trophy"
	KudoTransaction   = "Kudo"
)

const (
	Bronze   string = "Bronze"
	Silver   string = "Silver"
	Gold     string = "Gold"
	Diamond  string = "Diamond"
	Platinum string = "Platinum"
)

type User struct {
	ID                  string `gorm:"primary_key;type:char(36)"`
	UserName            string `gorm:"unique_index:idx_username;type:varchar(30);"`
	Password            string `sql:"type:char(60)"`
	FirstName           string `gorm:"type:varchar(20)"`
	LastName            string `gorm:"type:varchar(20)"`
	Email               string `gorm:"type:varchar(60)"`
	Category            string `gorm:"index:idx_category;type:varchar(30)"`
	About               string `gorm:"type:varchar(150)"`
	AvatarId            string `gorm:"type:char(36)"`
	IsLocked            bool
	IsAdmin             bool
	Stats_TotalTrophies uint
	Stats_TotalMedals   uint
	Stats_TotalKudos    uint
	Stats_TotalScore    uint
}

type ScoreTransaction struct {
	ID              string    `gorm:"primary_key;type:char(36)"`
	CreatedAt       time.Time `gorm:"index:idx_createdAt"`
	TransactionType string    `gorm:"index:idx_transactionType"`
	UserID          string    `gorm:"index:idx_userId;type:char(36)"`
	UserName        string    `gorm:"type:char(30)"`
	ItemDataId      string    `gorm:"type:char(36)"`
	GivenBy         string    `gorm:"type:varchar(30)"`
	GivenById       string    `gorm:"type:char(36)"`
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

type Avatar struct {
	ID              string `gorm:"primary_key;type:char(36)"`
	Head            string `gorm:"type:char(36)"`
	Body            string `gorm:"type:char(36)"`
	LeftHand        string `gorm:"type:char(36)"`
	RightHand       string `gorm:"type:char(36)"`
	WeaponRightHand string `gorm:"type:char(36)"`
	WeaponLeftHand  string `gorm:"type:char(36)"`
}

type AvatarItem struct {
	ID             string `gorm:"primary_key;type:char(36)"`
	image          string `gorm:"type:varchar(200)"`
	Type           string `gorm:"type:varchar(15)"`
	UnlockedByItem string `gorm:"type:char(36)"`
	Orientation    string `gorm:"type:varchar(15)"`
	PointsRequired uint
}
