package handler

import (
	r "go-project/repository"
	"go-project/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type urlHandler struct {
	urlSrv service.URLService
}

func NewURLHandler(urlSrv service.URLService) urlHandler {
	return urlHandler{urlSrv: urlSrv}
}

func (h urlHandler) CreateShortURL(c *fiber.Ctx) error {
	body := new(r.OriginalURLInsert)

	userProfile := c.Locals("userAuth")
	claims := userProfile.(jwt.MapClaims)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	userId := int(claims["id"].(float64))

	longURL, err := h.urlSrv.CreateShortURL(body.OriginalURL, userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot found long_url"})
	}

	return handleSuccess(c, longURL)
}
