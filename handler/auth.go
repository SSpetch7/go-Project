package handler

import (
	"fmt"
	"go-project/service"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authSrv service.AuthService
}

func NewAuthHandler(authSrv service.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) VerifyToken(c *fiber.Ctx) error {

	params := c.AllParams()

	fmt.Println("params", params)

	token, notNil := params["token"]

	if !notNil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot found token"})
	}

	payload, err := h.authSrv.VerifyToken(c.Context(), token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return handleSuccess(c, payload)
}
