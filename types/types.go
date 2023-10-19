package types

import (
	"fmt"
	"interview-follow/models"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// APPLICATION

func ApplicationSuccess(application interface{}) ApiResponse {
	return ApiResponse{
		Status: "success",
		Data:   application,
	}
}

func ApplicationNotFound(id string) ApiResponse {
	return ApiResponse{
		Status:  "success",
		Message: fmt.Sprintf("Cannot find application with id: %s", id),
	}
}

var ApplicationDeleteSuccess = ApiResponse{
	Status:  "failed",
	Message: "Application deleted successfully",
}

var ApplicationCreateFailed = ApiResponse{
	Status:  "failed",
	Message: "Couldn't create application",
}

//USER

func UserSuccess(user models.UserResponse) ApiResponse {
	return ApiResponse{
		Status: "success",
		Data:   user,
	}
}

func UsersFoundSuccess(users []models.UserResponse) ApiResponse {
	return ApiResponse{
		Status: "success",
		Data:   users,
	}
}

func SignInSuccess(token string) ApiResponse {
	return ApiResponse{
		Status:  "success",
		Message: "Signed in successfully",
		Data:    token,
	}
}

var InvalidCredentials = ApiResponse{
	Status:  "failed",
	Message: "Invalid Credentials",
}

var EmailAlreadyInUse = ApiResponse{
	Status:  "failed",
	Message: "Email already in use",
}

var HashError = ApiResponse{
	Status:  "failed",
	Message: "Couldn't hash password",
}

var UserCreationFailed = ApiResponse{
	Status:  "failed",
	Message: "Couldn't create user",
}

var UserCreationSuccess = ApiResponse{
	Status:  "success",
	Message: "User created successfully",
}

// JWT

var Unauthorized = ApiResponse{
	Status:  "failed",
	Message: "You are not authorized",
}

var InvalidUser = ApiResponse{
	Status:  "failed",
	Message: "The user belonging to this token no longer exists",
}

var InvalidTokenClaim = ApiResponse{
	Status:  "failed",
	Message: "Invalid token claim",
}

var InvalidateToken = ApiResponse{
	Status:  "failed",
	Message: "Invalidate token",
}
