package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/karan/practicelapbs/database"
	"github.com/karan/practicelapbs/models"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/login", LoginHandler)
}

func LoginHandler(c *fiber.Ctx) error {

	godotenv.Load()
	var authSecret = os.Getenv("AUTH_SECRET")
	bodyBytes := c.Body()
	fmt.Println("Raw request body:", string(bodyBytes))
	var existingUser models.User

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == gorm.ErrRecordNotFound {
		database.DB.Create(&user)
	} else {
		user = &existingUser
	}

	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(authSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	fmt.Println("User:", user)

	return c.JSON(fiber.Map{
		"message": "Success",
		"user":    user,
		"token":   signedToken,
	})
}
