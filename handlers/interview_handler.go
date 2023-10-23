package handlers

import (
	"interview-follow/db"
	"interview-follow/middleware"
	"interview-follow/models"
	"interview-follow/types"
	"interview-follow/utils"
	"interview-follow/validation"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteInterview(c *fiber.Ctx) error {
	id := c.Params("id")

	result := db.Database.Delete(&models.Interview{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(types.InterviewNotFound(id))
	}

	return c.JSON(types.InterviewDeleteSuccess)
}

func SetupInterviewRoutes(router fiber.Router) {
	interview := router.Group("interview")

	interview.Get("/all", middleware.DeserializeUser, GetAllInterviews)
	interview.Post("/new/:id", middleware.DeserializeUser, validation.ValidateNewInterview, NewInterview)
	interview.Put("/edit/:id", middleware.DeserializeUser, validation.ValidateEditInterview, EditInterview)
	interview.Delete("/delete/:id", middleware.DeserializeUser, DeleteInterview)
}

func EditInterview(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.InterviewRequest)
	c.BodyParser(&body)

	var interview models.Interview
	result := db.Database.First(&interview, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(types.InterviewNotFound(id))
	}

	parsedTime, err := time.Parse("2006-01-02T15:04:05", body.Date)
	if err != nil {
		log.Println(err)
	}

	interview.Date = parsedTime
	interview.Type = body.Type
	interview.Email = body.Email
	interview.Link = body.Link

	db.Database.Save(&interview)

	return c.JSON(types.ApplicationSuccess(models.GetInterviewResponse(interview)))
}

func NewInterview(c *fiber.Ctx) error {
	applicationId := c.Params("id")
	body := new(models.InterviewRequest)
	c.BodyParser(&body)

	parsedTime, err := time.Parse("2006-01-02T15:04:05", body.Date)
	if err != nil {
		log.Println(err)
	}

	newInterview := models.Interview{
		Id:            uuid.NewString(),
		Date:          parsedTime,
		Type:          body.Type,
		Email:         body.Email,
		EmailReceived: time.Now(),
		Link:          body.Link,
		ApplicationId: applicationId,
	}

	if err := db.Database.Create(&newInterview).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.InterviewCreateFailed)
	}

	return c.JSON(types.ApplicationSuccess(models.GetInterviewResponse(newInterview)))
}

func GetAllInterviews(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	model := db.Database.Joins("Application", db.Database.Where(&models.Application{UserId: user.Id})).Model(&models.Interview{})
	return c.JSON(types.InterviewSuccess(utils.Pg.With(model).Request(c.Request()).Response(&[]models.InterviewResponse{})))
}
