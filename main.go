package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

var client *auth.Client

func main() {
	var err error
	var app *firebase.App
	e := echo.New()
	e.HideBanner = true

	// create firebase client
	opt := option.WithCredentialsFile("fb.json")
	if app, err = firebase.NewApp(context.Background(), nil, opt); err != nil {
		e.Logger.Fatalf("error initializing app: %v", err)
	}

	if client, err = app.Auth(context.Background()); err != nil {
		e.Logger.Fatalf("error creating firebase client: %v", err)
	}

	e.Use(VerifyHeader)
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
			"error":   "no token",
			"details": "id-token missing in header",
		})
	}

	t, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return ctx.JSON(200, map[string]string{
			"error":   "invalid token",
			"message": " invalid id-token provided",
		})
	}

	fmt.Println("id-token verified")

	resp := make(map[string]string)
	resp["name"] = "Service A"
	resp["time"] = time.Now().Format(time.UnixDate)
	resp["uid"] = t.UID
	resp["project-url"] = t.Issuer
	resp["project-id"] = t.Audience

	return ctx.JSON(200, resp)
}
func serviceB(ctx echo.Context) error {
	idToken := ctx.Request().Header.Get("id-token")
	if idToken == "" {
		return ctx.JSON(200, map[string]string{
			"error": "no token",
		})
	}
	t, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return ctx.JSON(200, map[string]string{
			"error":   "invalid token",
			"message": " invalid id-token provided",
		})
	}
	fmt.Println("id-token verified")

	resp := make(map[string]string)
	resp["name"] = "Service B"
	resp["time"] = time.Now().Format(time.UnixDate)
	resp["uid"] = t.UID
	resp["project-url"] = t.Issuer
	resp["project-id"] = t.Audience

	return ctx.JSON(200, resp)
}

func server() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &http.Server{
		Addr:         ":" + port,       // port
		ReadTimeout:  20 * time.Minute, // read timeout
		WriteTimeout: 20 * time.Minute, // write timeout
	}
}

func VerifyHeader(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		log.Println(ctx.Request().RequestURI)
		return next(ctx)
	}
}
