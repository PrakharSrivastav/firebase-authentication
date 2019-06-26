package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Static("/", "public")

	grp := e.Group("/api/")
	grp.GET("serviceA", serviceA)
	grp.GET("serviceB", serviceB)

	e.Logger.Fatal(e.StartServer(server()))
}

func serviceA(ctx echo.Context) error {
	idToken := ctx.Request().Header.Get("id-token")
	if idToken == "" {
		return ctx.JSON(200, map[string]string{
			"error": "no token",
		})
	}
	resp := make(map[string]string)
	resp["name"] = "Service A"
	resp["time"] = time.Now().Format(time.UnixDate)
	return ctx.JSON(200, resp)
}
func serviceB(ctx echo.Context) error {
	idToken := ctx.Request().Header.Get("id-token")
	if idToken == "" {
		return ctx.JSON(200, map[string]string{
			"error": "no token",
		})
	}
	resp := make(map[string]string)
	resp["name"] = "Service B"
	resp["time"] = time.Now().Format(time.UnixDate)
	return ctx.JSON(200, resp)
}

func server() *http.Server {
	return &http.Server{
		Addr:         ":1313",          // port
		ReadTimeout:  20 * time.Minute, // read timeout
		WriteTimeout: 20 * time.Minute, // write timeout
	}
}
