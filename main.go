package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/logging"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	//fmt.Println("Endpoint Hit: homePage")
}
func main() {
	ctx := context.Background()
	// Sets your Google Cloud Platform project ID.
	projectID := "ingka-fulfilment-ordermgt-dev"
	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	logName := "test-api-log"
	l := client.Logger(logName)
	http.HandleFunc("/", homePage)
	l.StandardLogger(logging.Warning).Println("Welcome to the HomePage!")
	l.StandardLogger(logging.Error).Println("Starting server on PORT:8080")
	l.StandardLogger(logging.Critical).Fatal(http.ListenAndServe(":8080", nil))

}
