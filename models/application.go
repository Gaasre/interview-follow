package models

type Application struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userid"`
	User        User   `json:"user" gorm:"foreignKey:UserId"`
}

type ApplicationRequest struct {
	Company     string `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
}

type ApplicationResponse struct {
	Id          string `json:"id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetApplicationResponse(application Application) ApplicationResponse {
	return ApplicationResponse{
		Id:          application.Id,
		Company:     application.Company,
		Title:       application.Title,
		Description: application.Description,
	}
}
