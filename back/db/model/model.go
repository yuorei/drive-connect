package model

import (
	"time"
)

type User struct {
	ID              string `gorm:"primaryKey"`
	Email           string `gorm:"unique"`
	Name            string
	Password        string
	Driver          bool
	ProfileImageURL string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DriverBoard struct {
	ID          string `gorm:"primaryKey"`
	DriverID    string
	Title       string
	Description string
	Reward      string
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PassengerBoard struct {
	ID          string `gorm:"primaryKey"`
	PassengerID string
	Title       string
	Description string
	Reward      string
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Comment struct {
	ID          string `gorm:"primaryKey"`
	PostID      string
	CommenterID string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
