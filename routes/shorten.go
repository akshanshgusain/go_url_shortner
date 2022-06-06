package routes

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"urlshortner/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRaterRemaining int           `json:"rate_limit"`
	XRateLimitRest  time.Duration `json:"rate_limit_remaining"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "cannot parse JSON"})
	}

	// Rate limiting

	// check if the input is a valid url
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	// check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	// enforce TLS
	body.URL = helpers.EnforceHTTP(body.URL)
}
