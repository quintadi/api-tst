package main

import "github.com/gofiber/fiber/v2"

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}

func main() {
	app := fiber.New()

	app.Get("/signup", func(c *fiber.Ctx) error {
		req := new(SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		if req.Name == "" || req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Login inválido")
		}

		// Save info in DB

		// Create jwt token
		return nil
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "private"})
	})

	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "public"})
	})

	// app.Get("/", homePage() err)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

}

// func homePage(){
// 	return c.SendString("Hello, World 👋!")
// }
