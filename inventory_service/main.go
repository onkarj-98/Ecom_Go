package main

import (
	"context"
	"log"
	"time"

	"github.com/env-master"
)

// public api exposed is at 9000
var bindAddress = env.String("BIND_ADDRESS", false, ":9000", "Bind address for the server")

func main() {
	env.Parse()

	// mongodb connection

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://onkar_admin:<onkar>@cluster0-u0p80.mongodb.net/<inventory_service>?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)

}
