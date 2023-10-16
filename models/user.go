package models

type User struct {
	Id             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	HashedPassword string `json:"-"`
	Email          string `json:"email" gorm:"unique"`
}
