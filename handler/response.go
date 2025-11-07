package handler

import "github.com/gofiber/fiber/v2"

type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func handleSuccess(ctx *fiber.Ctx, data any) error {
	res := newResponse(true, "Success", data)
	return ctx.Status(fiber.StatusOK).JSON(res)
}
