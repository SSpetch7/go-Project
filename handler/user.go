package handler

import (
	"go-project/service"

	r "go-project/repository"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv service.UserService
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	user, err := h.userSrv.RegisterUser(c.Context(), body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return handleSuccess(c, user)

}

func (h userHandler) Login(c *fiber.Ctx) error {
	var body loginReq

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	user, err := h.userSrv.Login(c.Context(), body.Email, body.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return handleSuccess(c, user)
}
