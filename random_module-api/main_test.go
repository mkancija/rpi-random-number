package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRandomModuleGet(t *testing.T) {

	tests := []struct {
		info         string
		route        string // route path to test
		responseCode int    // HTTP code
	}{
		{
			info:         "HTTP 200 - only route that exists",
			route:        "/number",
			responseCode: 200,
		},
		{
			info:         "Http code 404 - update route does not exist",
			route:        "/data",
			responseCode: 404,
		},
		{
			info:         "Http code 404 - update route does not exist",
			route:        "/update",
			responseCode: 404,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/number", func(c *fiber.Ctx) error {

		return c.SendString("Random Number")
	})

	// Our test list:
	for _, test := range tests {
		// GET request for our test route
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// latency = -1 - no latency
		resp, _ := app.Test(req, -1)
		t.Log("testing route: ", test.route, " | status coude", resp.StatusCode)

		// Check response status codes.
		assert.Equalf(t, test.responseCode, resp.StatusCode, test.info)
	}

}
