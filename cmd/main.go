package main

import (
	"context"
	"log"
	"product/api"
	"product/internal/infrastructure"
)

func main() {
	// Inisialisasi koneksi MongoDB
	dbClient, err := infrastructure.NewMongoClient(context.Background(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer dbClient.Disconnect(context.Background())
	log.Println("Success connect to MongoDB")

	app := api.SetupApp(dbClient)

	log.Fatal(app.Listen(":3000"))
}
