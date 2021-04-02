package repository

import (
	"discount-service/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetProductById(productID string) (domain.Product, error)
}

type productRepository struct {
	productsCollection *mongo.Collection
}

func NewProductRepository(productsCollection *mongo.Collection) ProductRepository {
	return &productRepository{productsCollection}
}

func (c *productRepository) GetProductById(productID string) (domain.Product, error) {
	product := domain.Product{
		ID: "xxx",
	}

	return product, nil
}
