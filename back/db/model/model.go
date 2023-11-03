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

type Board struct {
	ID                   string `gorm:"primaryKey"`
	Type                 string
	UserID               string
	Description          string
	DepartureLatitude    float64
	DepartureLongitude   float64
	DestinationLatitude  float64
	DestinationLongitude float64
	Reward               string
	StartTime            time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type Comment struct {
	ID          string `gorm:"primaryKey"`
	PostID      string
	CommenterID string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
