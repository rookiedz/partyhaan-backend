package handlers

import fiber "github.com/gofiber/fiber/v2"

//AuthenHandler ...
type AuthenHandler struct{}

func NewAuthen() *AuthenHandler {
	return &AuthenHandler{}
}
func (ah *AuthenHandler) Login(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
func (ah *AuthenHandler) Register(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
