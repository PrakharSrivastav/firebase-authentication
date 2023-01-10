package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func main() {
	var err error
	var app *firebase.App
	ctx := context.Background()
	// create firebase client
	opt := option.WithCredentialsFile("fb.json")
	if app, err = firebase.NewApp(ctx, nil, opt); err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error creating firebase client: %v", err)
	}

	user, err := client.GetUserByEmail(ctx, "weather@test.no")
	if err != nil {
		log.Fatalf("client.GetUserByEmail: %v", err)
	}

	uid := user.UID
	claims := map[string]interface{}{
		"principal": "id-a",
		"scope":     "id-b",
		"companyID": "AOF",
	}

	err = client.SetCustomUserClaims(ctx, uid, claims)
	if err != nil {
		log.Fatalf("client.SetCustomUserClaims: %v", err)
	}
	log.Println("claims set")
}

//docker build  --platform=linux/amd64 . -t europe-north1-docker.pkg.dev/cip-shared-services/cip/firebase-ex:latest
//docker push europe-north1-docker.pkg.dev/cip-shared-services/cip/firebase-ex:latest
