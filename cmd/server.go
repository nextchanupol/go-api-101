package main

import (
	"log"
	"net/http"

	"github.com/nextchanupol/go-api-101/pkg/app"

	pgsql12 "github.com/nextchanupol/go-api-101/pkg/database"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const port = ":8080"

func main() {

	// new server
	e := echo.New()

	// use middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Cross-Origin
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Call database connection
	err := pgsql12.Init()

	if err != nil {
		log.Fatalf("can not init database connection; %v", err)
	}

	e.POST("/members", app.CreateMember)
	e.GET("/members", app.GetMembers)

	// start server
	e.Logger.Fatal(e.Start(port))

}
