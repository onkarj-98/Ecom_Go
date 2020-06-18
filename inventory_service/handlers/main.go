package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/env-master"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// public api exposed is at 9000
var bindAddress = env.String("BIND_ADDRESS", false, ":9000", "Bind address for the server")

func main() {
	env.Parse()
	l := log.New(os.Stdout, "inventory_service", log.LstdFlags)

	// mongodb connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://onkar_admin:<onkar>@cluster0-u0p80.mongodb.net/<inventory_service>?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to database is Succesful!")
	fmt.Println(client)

	// creating http server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      nil,               // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	//create a new serve mux and register new handlers

}
