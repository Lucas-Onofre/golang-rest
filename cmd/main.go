package main

import "github.com/Lucas-Onofre/golang-rest"

func main() {
	app := fiber.New()

	app.Get('/', func(c *fiber.Ctx) error {
		return c.sendString("Hello, World!")
	})

	app.Listen(":3000")
}