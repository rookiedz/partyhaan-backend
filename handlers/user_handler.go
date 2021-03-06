package handlers

import (
	"errors"
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
	return c.JSON(result)
}
func (uh *UserHandler) Register(c *fiber.Ctx) error {
	authen := new(models.Authen)
	if err := c.BodyParser(authen); err != nil {
		log.Println(err)
	}
	authenStore := stores.NewAuthenStore()
	a, err := authenStore.FindByEmail(authen.Email)
	if err != nil {
		log.Println(err)
	}
	authenEmpty := models.Authen{}
	if a == authenEmpty {
		return errors.New("ผู้ใช้งานนี้มีคนใช้แล้ว")
	}
	id, err := authenStore.Create(*authen)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(id)
}
