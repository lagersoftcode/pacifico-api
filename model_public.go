package main

import "time"

type PublicUser struct {
	ID                  string
	UserName            string
	Stats_TotalTrophies uint
	Stats_TotalMedals   uint
	Stats_TotalKudos    uint
	Stats_TotalScore    uint
}

type PublicAction struct {
	CreatedAt  time.Time
	SourceUser string
	TargetUser string
	Item       string
}

type UserTrophy struct {
	CreatedAt       time.Time
	UserName        string
	GivenBy         string
	TransactionType string
	Image           string
	Points          int
	Name            string
	Description     string
}

type UserMedal struct {
	CreatedAt       time.Time
	UserName        string
	GivenBy         string
	TransactionType string
	Image           string
	Points          int
	Name            string
	Descrtiption    string
	Material        string
}

type UserKudo struct {
	CreatedAt       time.Time
	UserName        string
	GivenBy         string
	TransactionType string
}
