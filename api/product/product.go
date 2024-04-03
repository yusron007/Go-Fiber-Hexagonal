package product

import (
	"context"
	"product/internal/domain"
	"product/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProducts(productService *service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mendapatkan nilai page dan limit
		page, _ := strconv.Atoi(c.Query("page", "1"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))

		ctx := context.Background()
		products, totalData, err := productService.GetAll(ctx, page, limit)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		startNo := (page - 1) * limit
		var numberedProducts []map[string]interface{}

		for i, product := range products {
			productMap := make(map[string]interface{})
			productMap["No"] = startNo + i + 1
			productMap["product_id"] = product.Id
			productMap["product_name"] = product.ProductName
			productMap["stock"] = product.Stock
			numberedProducts = append(numberedProducts, productMap)
		}

		//respons
		response := fiber.Map{
			"data": numberedProducts,
			"meta": fiber.Map{
				"totalData": totalData,
				"code":      fiber.StatusOK,
				"status":    "success",
				"message":   "Data Product",
			},
		}

		return c.JSON(response)
	}
}

func GetProductByID(productService *service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")

		// Convert product ID
		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}

		ctx := context.Background()
		product, err := productService.GetById(ctx, objID)
		if err != nil {
			if err == domain.ErrProductNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := fiber.Map{
			"data": product,
			"meta": fiber.Map{
				"code":    fiber.StatusOK,
				"status":  "success",
				"message": "Detail Data Product",
			},
		}

		return c.JSON(response)
	}
}

func CreateProducts(productService *service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var product domain.Product
		if err := c.BodyParser(&product); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		ctx := context.Background()
		err := productService.Create(ctx, &product)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := fiber.Map{
			"meta": fiber.Map{
				"code":    fiber.StatusOK,
				"status":  "Success",
				"message": "Add Product Success",
			},
		}

		return c.JSON(response)
	}
}

func UpdateProductByID(productService *service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")

		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}

		var product domain.Product
		if err := c.BodyParser(&product); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}

		ctx := context.Background()
		err = productService.Update(ctx, objID, &product)
		if err != nil {
			if err == domain.ErrProductNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := fiber.Map{
			"meta": fiber.Map{
				"code":    fiber.StatusOK,
				"status":  "Success",
				"message": "Update Product Success",
			},
		}

		return c.JSON(response)
	}
}

func DeleteProductByID(productService *service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")

		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}

		// Delete product
		ctx := context.Background()
		err = productService.Delete(ctx, objID)
		if err != nil {
			if err == domain.ErrProductNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := fiber.Map{
			"meta": fiber.Map{
				"code":    fiber.StatusOK,
				"status":  "Success",
				"message": "Delete Product Success",
			},
		}

		return c.JSON(response)
	}
}
