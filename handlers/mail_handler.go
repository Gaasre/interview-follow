package handlers

import (
	"interview-follow/db"
	"interview-follow/middleware"
	"interview-follow/models"
	"interview-follow/openai"
	"interview-follow/types"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Mail struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// TODO: This is just a test. Need to move it to email interception
func SetupMailRoutes(router fiber.Router) {
	mail := router.Group("mail")

	mail.Post("/new", middleware.DeserializeUser, func(c *fiber.Ctx) error {
		user := c.Locals("user").(models.UserResponse)
		body := new(Mail)
		c.BodyParser(&body)

		parsedMail := openai.ParseEmail(body.From, body.Subject, body.Body)

		if parsedMail.Company != "" {
			// Find the job based on the company
			var application models.Application
			result := db.Database.First(&application, "company LIKE ?", "%"+parsedMail.Company+"%")

			if result.RowsAffected > 0 {
				application.Stage = parsedMail.Stage
				db.Database.Save(&application)

				parsedTime, err := time.Parse("2006-01-02T15:04:05", parsedMail.Date)
				if err != nil {
					log.Println(err)
					parsedTime = time.Time{}
				}

				if parsedMail.Interview {
					newInterview := models.Interview{
						Id:            uuid.NewString(),
						Type:          parsedMail.Type,
						Email:         parsedMail.Summary,
						EmailReceived: time.Now(),
						Date:          parsedTime,
						Link:          parsedMail.Link,
						ApplicationId: application.Id,
					}

					if err := db.Database.Create(&newInterview).Error; err != nil {
						return c.Status(fiber.StatusInternalServerError).JSON(types.ApplicationCreateFailed)
					}

					return c.JSON(types.ApplicationSuccess(models.GetInterviewResponse(newInterview)))
				}
				return c.JSON(types.ApplicationUpdatedSuccess)

			} else {
				newApplication := models.Application{
					Id:          uuid.NewString(),
					Title:       parsedMail.Title,
					Description: "",
					Company:     parsedMail.Company,
					Link:        "",
					Stage:       parsedMail.Stage,
					UserId:      user.Id,
				}

				if err := db.Database.Create(&newApplication).Error; err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(types.ApplicationCreateFailed)
				}

				parsedTime, err := time.Parse("2006-01-02T15:04:05", parsedMail.Date)
				if err != nil {
					log.Println(err)
					parsedTime = time.Time{}
				}

				if parsedMail.Interview {
					newInterview := models.Interview{
						Id:            uuid.NewString(),
						Type:          parsedMail.Type,
						Email:         parsedMail.Summary,
						EmailReceived: time.Now(),
						Date:          parsedTime,
						Link:          parsedMail.Link,
						ApplicationId: newApplication.Id,
					}

					if err := db.Database.Create(&newInterview).Error; err != nil {
						return c.Status(fiber.StatusInternalServerError).JSON(types.InterviewCreateFailed)
					}
				}

				return c.JSON(types.ApplicationSuccess(models.GetApplicationResponse(newApplication)))
			}
		}

		return c.JSON(parsedMail)
	})
}
