package models

import "time"

type Interview struct {
	Id            string      `json:"id" gorm:"primaryKey"`
	Date          time.Time   `json:"date"`
	Type          string      `json:"type"` // Coding Challenge/HR Screening/etc...
	Email         string      `json:"email"`
	EmailReceived time.Time   `json:"emailReceived"`
	Link          string      `json:"link"`
	ApplicationId string      `json:"applicationid"`
	Application   Application `json:"application" gorm:"foreignKey:ApplicationId"`
}
