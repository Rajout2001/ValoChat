package main

import (
	"fmt"
	"log"
	"net/http"
	database "pc3r/database"
	"pc3r/prisma"
	"pc3r/services"
	ws "pc3r/webSocket"

	"github.com/jasonlvhit/gocron"
	"github.com/rs/cors"
)

var mux = http.NewServeMux()

func main() {

	services.UseHttpRouter(mux)
	ws.UseSocketRouter(mux)
	// Establish connection with the remote database using Prisma ORM
	prisma.Init()

	cors_options := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	})

	handler := cors_options.Handler(mux)
	
	// lancer le serveur dans un Thread séparé pour maximiser la concurrence
	go func() {
		fmt.Println("Server running on PORT 5001")
		err := http.ListenAndServe(":5001", handler)
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	// Push data once in a hour
	
	sched := gocron.NewScheduler()
	sched.Every(1).Hour().Do(func() {
		database.PushData()
	})
	
	<- sched.Start()
	
	
	select{}
}

