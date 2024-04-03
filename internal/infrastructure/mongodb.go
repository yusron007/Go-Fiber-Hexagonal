package infrastructure

import (
	"context"
	"log"
	"product/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type mongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(dbClient *mongo.Client) domain.ProductRepository {
	return &mongoProductRepository{
		collection: dbClient.Database("db_produk").Collection("produk"),
	}
}

func (repo *mongoProductRepository) GetAll(ctx context.Context, page, limit int) ([]domain.Product, int, error) {
	start := time.Now()

	var products []domain.Product

	// Hitung total data
	totalData, err := repo.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error while counting documents: %s\n", err)
		return nil, 0, err
	}

	offset := (page - 1) * limit

	cursor, err := repo.collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		log.Printf("Error while finding documents: %s\n", err)
		return nil, 0, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &products); err != nil {
		log.Printf("Error while decoding documents: %s\n", err)
		return nil, 0, err
	}

	//waktu selesai
	duration := time.Since(start)
	log.Printf("GetAll function executed in: %s\n", duration)

	return products, int(totalData), nil
}

func (repo *mongoProductRepository) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
	filter := bson.M{"_id": id}

	var product domain.Product
	err := repo.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		// Jika dokumen tidak ditemukan, kembalikan error
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}

	// Jika berhasil ditemukan, kembalikan produk
	return &product, nil
}

func (repo *mongoProductRepository) Create(ctx context.Context, product *domain.Product) error {
	_, err := repo.collection.InsertOne(ctx, product)
	return err
}

func (repo *mongoProductRepository) Update(ctx context.Context, id primitive.ObjectID, product *domain.Product) error {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"product_name": product.ProductName,
			"stock":        product.Stock,
		},
	}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *mongoProductRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}
