package main

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func homePage(c echo.Context) error {
  title := "Connected to Homepage"
  return c.String(http.StatusOK, title)
}

func gotify(c echo.Context) error {
	// Get title and message from the query string
	title := c.QueryParam("title")
	message := c.QueryParam("message")
  token := c.QueryParam("token")

  http.PostForm("https://push.ldsnetwork.xyz/message?token=" + token,
    url.Values{"message": {message}, "title": {title}})

	return c.String(http.StatusOK, "title:" + title + ", message:" + message)
}

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", homePage)
  e.GET("/gotify", gotify)

  // Healthcheck Endpoint For Docker Container
  e.GET("/healthcheck", func(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
  })

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

