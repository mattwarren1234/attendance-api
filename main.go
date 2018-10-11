package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mattwarren1234/attendance-api/auth"
	"github.com/mattwarren1234/attendance-api/member"
)

func main() {
	noauthPtr := flag.Bool("noauth", false, "disables auth")
	flag.Parse()
	disableAuth := *noauthPtr
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

	// NOTE: HAS TO GO AFTER AUTH.LOGIN.
	if !disableAuth {
		// don't run auth on the given paths
		allowedPaths := []string{"/", "/auth/login/"}
		e.Use(auth.Validate(allowedPaths))
	}

	// Routes
	e.GET("/", hello)
	// e.USE
	e.POST("/auth/login/", auth.CreateToken)
	e.GET("/test/", hello)
	members := e.Group("/members")
	members.GET("/", member.GetAll)
	members.GET("/:id/", member.GetByID)

	eventGroup := e.Group("/events")
	eventGroup.POST("/create/", member.CreateEvent)
	eventGroup.GET("/", member.GetAttendanceCountByDay)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
