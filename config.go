package zlogres

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Optional. Default: "requestid"
	RequestIDContextKey string

	// Optional. Default: "info"
	LogLevel string

	// Optional. Default: "micro". Possible Value: ["nano", "micro", "milli"]
	ElapsedTimeUnit string

	// Optional. Default: "message". Use ContextMessageKey by set `c.Locals("message", "WHATEVER_MESSAGE")`
	ContextMessageKey string
}

var ConfigDefault = Config{
	Next:                nil,
	RequestIDContextKey: "requestid",
	LogLevel:            "info",
	ElapsedTimeUnit:     "micro",
	ContextMessageKey:   "message",
}

func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// set default
	if cfg.RequestIDContextKey == "" {
		cfg.RequestIDContextKey = ConfigDefault.RequestIDContextKey
	}

	if cfg.LogLevel == "" {
		cfg.LogLevel = ConfigDefault.LogLevel
	}

	if cfg.ElapsedTimeUnit == "" {
		cfg.ElapsedTimeUnit = ConfigDefault.ElapsedTimeUnit
	}

	if cfg.ContextMessageKey == "" {
		cfg.ContextMessageKey = ConfigDefault.ContextMessageKey
	}

	return cfg
}
