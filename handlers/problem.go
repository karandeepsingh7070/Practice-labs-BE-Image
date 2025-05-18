package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karan/practicelapbs/utils"
)

func GetAllProblems(c *fiber.Ctx) error {
	return c.JSON(utils.Problems)
}

func GetProblemBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	for _, p := range utils.Problems {
		if p.TitleSlug == slug {
			return c.JSON(p)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Problem not found",
	})
}
