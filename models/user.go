package models

type User struct {
	Id             string `gorm:"primaryKey"`
	Name           string
	HashedPassword string
	Email          string `gorm:"unique"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SignupRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Name     string `validate:"required"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func FilterPassword(user User) UserResponse {
	return UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
