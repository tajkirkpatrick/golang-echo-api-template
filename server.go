package main

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Structs
// type basicJSON struct {
//   Status string `json:"title"`
//   Message string `json:"message"`
// }

// Handlers
func homePage(c echo.Context) error {
  return c.String(http.StatusOK, "Connected to Homepage")
}

func gotify(c echo.Context) error {
	// Get title and message from the query string
	title := c.QueryParam("title")
	message := c.QueryParam("message")
  token := c.QueryParam("token")

  if len(strings.TrimSpace(title)) == 0 || len(strings.TrimSpace(message)) == 0 || len(strings.TrimSpace(token)) == 0 {
    return c.JSON(http.StatusBadRequest, struct {
      Status string `json:"status"`
      Message string `json:"message"`
    }{
      Status: "Bad Request",
      Message: "Missing Title, Message, or Token",
    })
  }

  http.PostForm("https://push.ldsnetwork.xyz/message?token=" + token,
    url.Values{"message": {message}, "title": {title}})

	return c.JSON(http.StatusOK, struct {
    Status string `json:"status"`
    Title string `json:"title"`
    Message string `json:"message"`
  }{
    Status: "Sent",
    Title: title,
    Message: message,
  })
}

func main() {

  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Healthcheck Endpoint For Docker Container
  e.GET("/healthcheck", func(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
  })

  e.GET("", homePage)
  e.GET("/", homePage)

  api := e.Group("/api")
  api.GET("", func(c echo.Context) error {
    return c.String(http.StatusOK, "Connected to API")
  })
  api.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Connected to API")
  })
  api.GET("/gotify", gotify)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

