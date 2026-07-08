package models

import "time"

type Meter struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	SiteID    string    `json:"site_id" gorm:"type:uuid"`
	Type      string    `json:"type"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
