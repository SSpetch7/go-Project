package handler

import (
	"fmt"
	"go-project/service"

	r "go-project/repository"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {

	users, err := h.userSrv.GetUsers()

	if err != nil {
		return err
	}

	return handleSuccess(c, users)
}

func (h userHandler) RegisterUser(c *fiber.Ctx) error {
	body := new(r.NewUserRequest)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}
	_, err := h.userSrv.RegisterUser(c.Context(), body)

	fmt.Println("err", err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return handleSuccess(c, body)

}
