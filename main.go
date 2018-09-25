package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	member "github.com/mattwarren1234/attendance-api/member"
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
	// e.GET("/api/user", apiUser)

	g := e.Group("/members")
	g.GET("/", getAll)
	g.GET("/:id/", getByID)
	// e.GET("/api/articles/:slug/comments", articles)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func getByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	attendance, err := member.GetAttendanceByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, attendance)
}

func getAll(c echo.Context) error {
	members, err := member.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, members)
}

// func apiUser(c echo.Context) error {
// 	// response is ?
// 	response := []string{"some stuff"}
// 	val := map[string]interface{}{
// 		"user": response}
// 	return c.JSON(http.StatusOK, val)
// }

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
