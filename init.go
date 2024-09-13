package zlogres

import (
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
)

const ( // <-- this should be configurable
	// LogType    string = "info"
	// TimeLevel  string = "microsecond"

	// Note: This line below is not need to defined, bcoz time.RFC3339 is the default of zerolog
	TimeFieldFormat    string = time.RFC3339Nano
	TimestampFieldName string = "timestamp"

	EventTag              string = "event"
	EventNameTag          string = "name"
	URLTag                string = "url"
	MethodTag             string = "method"
	RequestTimeTag        string = "request_time"
	ResponseTimeTag       string = "response_time"
	ResponseStatusCodeTag string = "status_code"
	ResponseStatusTextTag string = "status_text"
	TimeUsageTag          string = "elapsed_time"
	TimeUsageUnitTag      string = "elapsed_time_unit"

	EventNameValue string = "zlogres"
)

func init() {
	zerolog.TimeFieldFormat = TimeFieldFormat
	zerolog.TimestampFieldName = TimestampFieldName
}

func New(config ...Config) fiber.Handler {
	// set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// start clock
		begin := time.Now()

		// let do the request as normal
		c.Next()

		// 'baby come back to me'
		interceptedResponse := c.Response()
		statusCode := interceptedResponse.StatusCode()
		elapsedTime := getTimeDuration(time.Since(begin), cfg.ElapsedTimeUnit)

		logger := getLogLevel(cfg.LogLevel)
		logger = logger.
			Interface(EventTag, map[string]interface{}{
				EventNameTag:          EventNameValue,
				URLTag:                c.OriginalURL(),
				MethodTag:             c.Method(),
				ResponseStatusCodeTag: statusCode,
				ResponseStatusTextTag: fasthttp.StatusMessage(statusCode),
				RequestTimeTag:        begin.Format(TimeFieldFormat),
				ResponseTimeTag:       time.Now().Format(TimeFieldFormat),
				TimeUsageTag:          elapsedTime,
				TimeUsageUnitTag:      cfg.ElapsedTimeUnit,
			})

		if reqID := c.Locals(cfg.RequestIDContextKey); reqID != nil {
			logger = logger.Str(strings.ReplaceAll(cfg.RequestIDContextKey, "-", "_"), reqID.(string))
		}

		msg := c.Locals(cfg.ContextMessageKey)
		if msg == nil {
			msg = ""
		}
		logger.Msgf("%v", msg)

		// Idk to return the same response; if you have the better way, please tell me.
		return c.Send(interceptedResponse.Body())
	}
}

func getTimeDuration(timeDuration time.Duration, unit string) int64 {
	switch unit {
	case "nano":
		return timeDuration.Nanoseconds()
	case "micro":
		return timeDuration.Microseconds()
	case "milli":
		return timeDuration.Milliseconds()
	default:
		return 0
	}
}

func getLogLevel(level string) *zerolog.Event {
	switch level {
	case "debug":
		return log.Logger.Debug()
	case "info":
		return log.Logger.Info()
	case "warn":
		return log.Logger.Warn()
	case "error":
		return log.Logger.Error()
	case "fatal":
		return log.Logger.Fatal()
	case "panic":
		return log.Logger.Panic()
	default:
		return log.Logger.Log()
	}
}
