package handler

import (
	r "go-project/repository"
	"go-project/service"

	"github.com/gofiber/fiber/v2"
)

type urlHandler struct {
	urlSrv service.URLService
}

func NewURLHandler(urlSrv service.URLService) urlHandler {
	return urlHandler{urlSrv: urlSrv}
}

func (h urlHandler) CreateShortURL(c *fiber.Ctx) error {
	body := new(r.NewURLRequest)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	longURL, err := h.urlSrv.HashURL(body.LongURL)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot found long_url"})
	}

	return handleSuccess(c, longURL)
}
