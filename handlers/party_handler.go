package handlers

import (
	"partyhann/backend/models"

	fiber "github.com/gofiber/fiber/v2"
)

//PartyHandler ...
type PartyHandler struct{}

func NewParty() *PartyHandler {
	return &PartyHandler{}
}
func (ph *PartyHandler) All(c *fiber.Ctx) error {
	return c.JSON(models.Party{})
}
func (ph *PartyHandler) Find(c *fiber.Ctx) error {
	return nil
}
func (ph *PartyHandler) Users(c *fiber.Ctx) error {
	return nil
}
func (ph *PartyHandler) Crete(c *fiber.Ctx) error {
	return nil
}
func (ph *PartyHandler) Update(c *fiber.Ctx) error {
	return nil
}
func (ph *PartyHandler) Delete(c *fiber.Ctx) error {
	return nil
}
