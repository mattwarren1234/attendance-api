package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mattwarren1234/attendance-api/auth"
	"github.com/mattwarren1234/attendance-api/member"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	// Routes
	e.GET("/", hello)

	e.POST("/auth/login", auth.CreateToken)
	// e.GET("/api/user", apiUser)

	e.GET("/members/", member.GetAll)
	e.GET("/members/:id/", member.GetByID)
	e.GET("/events/", member.GetAttendanceCountByDay)
	// g := e.Group("/members")
	// g.GET("/", member.GetAll)
	// g.GET("/:id/", member.GetByID)

	// eventGroup := e.Group("/events")
	// eventGroup.GET("/", member.GetAttendanceCountByDay)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
