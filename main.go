package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.File("/", "public/index.html")

	grp := e.Group("/api/")
	grp.GET("serviceA", serviceA)
	grp.GET("serviceB", serviceB)

	e.Logger.Fatal(e.StartServer(server()))
}

func serviceA(ctx echo.Context) error {
	resp := make(map[string]string)
	resp["name"] = "Service A"
	return ctx.JSON(200, resp)
}
func serviceB(ctx echo.Context) error {
	resp := make(map[string]string)
	resp["name"] = "Service B"
	return ctx.JSON(200, resp)
}

func server() *http.Server {
	return &http.Server{
		Addr:         ":1313",          // port
		ReadTimeout:  20 * time.Minute, // read timeout
		WriteTimeout: 20 * time.Minute, // write timeout
	}
}
