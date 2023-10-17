package handlers

import (
	"fmt"
	"interview-follow/db"
	"interview-follow/middleware"
	"interview-follow/models"
	"interview-follow/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/morkid/paginate"
)

var pg = paginate.New()

func SetupApplicationRoutes(router fiber.Router) {
	application := router.Group("application")

	application.Get("/all", middleware.DeserializeUser, GetAllApplications)
	application.Post("/new", middleware.DeserializeUser, validation.ValidateNewApplication, NewApplication)
	application.Put("/edit/:id", middleware.DeserializeUser, validation.ValidateEditApplication, EditApplication)
	application.Delete("/delete/:id", middleware.DeserializeUser, DeleteApplication)
}

func DeleteApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	result := db.Database.Delete(&models.Application{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Cannot find application with id: %s", id)})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Application deleted successfully"})
}

func EditApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(models.ApplicationRequest)
	c.BodyParser(&body)

	var application models.Application
	result := db.Database.First(&application, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": fmt.Sprintf("Cannot find application with id: %s", id)})
	}

	application.Company = body.Company
	application.Description = body.Description
	application.Title = body.Title

	db.Database.Save(&application)

	return c.JSON(fiber.Map{"status": "success", "message": "Application updated successfully", "data": models.GetApplicationResponse(application)})
}

func NewApplication(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	body := new(models.ApplicationRequest)
	c.BodyParser(&body)

	newApplication := models.Application{
		Id:          uuid.NewString(),
		Title:       body.Title,
		Description: body.Description,
		Company:     body.Company,
		UserId:      user.Id,
	}

	if err := db.Database.Create(&newApplication).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create application"})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Application created successfully", "data": models.GetApplicationResponse(newApplication)})
}

func GetAllApplications(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	model := db.Database.Where("user_id = ?", user.Id).Model(&models.Application{})
	return c.JSON(pg.With(model).Request(c.Request()).Response(&[]models.ApplicationResponse{}))
}
