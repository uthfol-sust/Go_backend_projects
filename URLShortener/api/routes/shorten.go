package routes

import (
	"os"
	"strconv"
	"time"

	"urlshorten/database"
	"urlshorten/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRemaining     int           `json:"rate_limit"`
	XRateLimitRest int           `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	/* ---------------------------- RATE LIMIT ---------------------------- */

	r2 := database.CreateClient(0) // FIXED (DB1 removed)
	defer r2.Close()

	quota, _ := strconv.Atoi(os.Getenv("API_QUOTA"))

	// Check quota
	val, err := r2.Get(database.Ctx, c.IP()).Result()

	if err == redis.Nil {
		r2.Set(database.Ctx, c.IP(), quota, 30*time.Minute)
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			ttl := r2.TTL(database.Ctx, c.IP()).Val()

			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": int(ttl.Minutes()),
			})
		}
	}

	/* ----------------------------- VALIDATION ---------------------------- */

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "URL is required"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid URL format"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	/* ---------------------------- SHORT ID ---------------------------- */

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	// Check if ID already exists
	exists, _ := r.Get(database.Ctx, id).Result()
	if exists != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "URL custom short is already in use"})
	}

	// Default expiry: 24 hours
	if body.Expiry == 0 {
		body.Expiry = 24
	}

	// Save short -> URL
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unable to connect to server",
		})
	}

	/* ------------------------- SUCCESS RESPONSE ------------------------ */

	r2.Decr(database.Ctx, c.IP())

	remaining, _ := r2.Get(database.Ctx, c.IP()).Result()
	remainInt, _ := strconv.Atoi(remaining)
	ttl := r2.TTL(database.Ctx, c.IP()).Val()

	resp := response{
		URL:            body.URL,
		CustomShort:    os.Getenv("DOMAIN") + "/" + id,
		Expiry:         body.Expiry,
		XRemaining:     remainInt,
		XRateLimitRest: int(ttl.Minutes()),
	}

	return c.Status(200).JSON(resp)
}
