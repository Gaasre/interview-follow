package models

type Application struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userid"`
	User        User   `json:"user" gorm:"foreignKey:UserId"`
}
