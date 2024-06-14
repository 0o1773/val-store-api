package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type ReqBody struct {
	discordId string
}

func handler(w http.ResponseWriter, r *http.Request) {
	projectId := os.Getenv("firebase_projectId")

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	log.Print(string(body))

	reqBody := ReqBody{}
	json.Unmarshal(body, &reqBody)

	discordId := reqBody.discordId

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	
	_, err:= client.Collection("users").get()

}
