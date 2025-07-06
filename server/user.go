package server

import "github.com/gofiber/fiber/v2"

func (s Server) LoginUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Send([]byte("Hello World!"))
}

func (s Server) LogoutUser(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
