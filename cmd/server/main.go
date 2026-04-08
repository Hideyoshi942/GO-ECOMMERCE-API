package main

import (
	"log"

	"go-ecomerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
