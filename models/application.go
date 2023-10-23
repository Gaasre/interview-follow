package models

type Application struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Stage       string `json:"stage"`
	UserId      string `json:"userid"`
	User        User   `json:"user" gorm:"foreignKey:UserId"`
}

type ApplicationRequest struct {
	Company     string `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Link        string `validate:"required,url"`
	Stage       string `validate:"required"`
}

type ApplicationResponse struct {
	Id          string `json:"id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Stage       string `json:"stage"`
	Description string `json:"description"`
}

func GetApplicationResponse(application Application) ApplicationResponse {
	return ApplicationResponse{
		Id:          application.Id,
		Company:     application.Company,
		Title:       application.Title,
		Link:        application.Link,
		Description: application.Description,
		Stage:       application.Stage,
	}
}
