package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type RateLimiterConfig struct {
	Max        int
	Expiration time.Duration
	Message    string
}

func NewCustomRateLimiter(cfg RateLimiterConfig) fiber.Handler {
	if cfg.Max == 0 {
		cfg.Max = 10
	}
	if cfg.Expiration == 0 {
		cfg.Expiration = 1 * time.Minute
	}
	if cfg.Message == "" {
		cfg.Message = "Too Many Requests"
	}

	return limiter.New(limiter.Config{
		Max:        cfg.Max,
		Expiration: cfg.Expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": cfg.Message,
			})
		},
	})
}

func AuthRateLimiter() fiber.Handler {
	return NewCustomRateLimiter(RateLimiterConfig{
		Max:        3,
		Expiration: 1 * time.Minute,
		Message:    "Too many requests. Please try again later.",
	})
}

func APIRateLimiter() fiber.Handler {
	return NewCustomRateLimiter(RateLimiterConfig{
		Max:        100,
		Expiration: 1 * time.Minute,
		Message:    "API rate limit exceeded.",
	})
}
