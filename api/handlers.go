package api

import (
	"product/api/product"
	"product/internal/infrastructure"
	"product/internal/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupApp(dbClient *mongo.Client) *fiber.App {
	app := fiber.New()

	// Setup repository
	productRepo := infrastructure.NewMongoProductRepository(dbClient)

	// Setup service
	productService := service.NewProductService(productRepo)

	// Setup routes
	app.Get("/products", product.GetAllProducts(productService))
	app.Get("/products/:id", product.GetProductByID(productService))
	app.Post("/products", product.CreateProducts(productService))
	app.Put("/products/:id", product.UpdateProductByID(productService))
	app.Delete("/products/:id", product.DeleteProductByID(productService))

	return app
}
