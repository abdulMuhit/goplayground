package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	log2 := logrus.New()

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log2.Out = file
	} else {
		log2.Info("Failed to log to file, using default stderr")
	}

	app.Get("/api/adjust/callback", func(c *fiber.Ctx) error {

		// Get the query parameters
		params := c.Queries()

		fmt.Println("params", params)

		// Get all the log from get request
		data := c.Request().URI().String()

		for key, value := range params {
			log2.WithFields(logrus.Fields{
				"Query Parameter": key,
				"Value":           value,
			}).Info("params")
		}

		// log all the data from get request

		// Return a JSON response
		return c.JSON(fiber.Map{
			"message": "Params logged successfully!",
			"params":  params,
			"data":    data,
		})
	})

	// Start the Fiber app on port 3000
	err = app.Listen(":3000")
	if err != nil {
		// Log the error
		log.Fatal(err)

	}
}
