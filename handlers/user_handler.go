package handlers

import (
	"log"
	"partyhaan/backend/models"
	"partyhaan/backend/stores"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func NewUser() *UserHandler {
	return &UserHandler{}
}
func (uh *UserHandler) Authenticate(c *fiber.Ctx) error {
	authen := new(models.Authen)
	if err := c.BodyParser(authen); err != nil {
		log.Println(err)
	}
	authenStore := stores.NewAuthenStore()
	result, err := authenStore.FindByEmail(authen.Email)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
	return nil
}
func (uh *UserHandler) Register(c *fiber.Ctx) error {
	return nil
}
