package main

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/owlsome-official/zlogres"
)

var (
	MY_REQUEST_ID_KEY      string = "myrequestlonglongid"
	MY_CONTEXT_MESSAGE_KEY string = "msg"
)

func main() {

	// Default
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(zlogres.New())

	app.Get("/", HandlerDefault)       // GET http://localhost:8000/
	app.Get("/msg/*", HandlerMsgParam) // GET http://localhost:8000/msg/{MESSAGE}

	fmt.Println("Listening on http://localhost:8000")
	fmt.Println("Try to send a request :D")
	go app.Listen(":8000")

	fmt.Println("// ----------------------------------------------- //")

	// Custom
	customApp := fiber.New(fiber.Config{DisableStartupMessage: true})
	customApp.Use(requestid.New(requestid.Config{ContextKey: MY_REQUEST_ID_KEY}))
	customApp.Use(zlogres.New(zlogres.Config{
		RequestIDContextKey: MY_REQUEST_ID_KEY,
		LogLevel:            "debug",
		ElapsedTimeUnit:     "nano",
		ContextMessageKey:   MY_CONTEXT_MESSAGE_KEY,
	}))

	customApp.Get("/msg_custom", HandlerCustomMsgParam) // GET http://localhost:8000/msg_custom

	fmt.Println("[CUSTOM] Listening on http://localhost:8001")
	fmt.Println("[CUSTOM] Try to send a request :D")
	customApp.Listen(":8001")

}

func HandlerDefault(c *fiber.Ctx) error {
	beautyCallLog("HandlerDefault")
	return c.SendString("Watch your app logs!")
}

func HandlerMsgParam(c *fiber.Ctx) error {
	beautyCallLog("HandlerMsgParam")

	msg := c.Params("*")
	c.Locals("message", msg) // Set context "message"

	return c.SendString("Watch your app logs! and see the difference (Hint: `message` will show on your logs)")
}

func HandlerCustomMsgParam(c *fiber.Ctx) error {
	beautyCallLog("HandlerCustomMsgParam")

	msg := "CUSTOM CONTEXT MESSAGE"
	c.Locals(MY_CONTEXT_MESSAGE_KEY, msg) // Set context "message" for zlogres

	return c.SendString("Watch your app logs! and see the difference (Hint: `message` will show on your logs)")
}

func beautyCallLog(called string) {
	m := regexp.MustCompile(".")
	dashed := "------------" + m.ReplaceAllString(called, "-") + "----"
	fmt.Println(dashed)
	fmt.Printf("--- Called: %v ---\n", called)
	fmt.Println(dashed)
}
