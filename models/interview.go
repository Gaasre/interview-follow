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

type InterviewRequest struct {
	Date  string `validate:"required"`
	Type  string `validate:"required"`
	Email string `validate:"required"`
	Link  string `validate:"required"`
}

type InterviewResponse struct {
	Id            string    `json:"id" gorm:"primaryKey"`
	Date          time.Time `json:"date"`
	Type          string    `json:"type"`
	Email         string    `json:"email"`
	EmailReceived time.Time `json:"emailReceived"`
	Link          string    `json:"link"`
}

func GetInterviewResponse(interview Interview) InterviewResponse {
	return InterviewResponse{
		Id:            interview.Id,
		Date:          interview.Date,
		Type:          interview.Type,
		Email:         interview.Email,
		EmailReceived: interview.EmailReceived,
		Link:          interview.Link,
	}
}
