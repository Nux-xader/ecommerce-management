package repositories

import (
	"context"

	"github.com/Nux-xader/ecommerce-management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

func InitProductRepository(db *mongo.Database) {
	productCollection = db.Collection("products")
}

func GetAllProducts(
	userID primitive.ObjectID,
	IDs ...primitive.ObjectID,
) ([]models.Product, error) {
	var products []models.Product
	filter := bson.M{"user_id": userID}

	// Check if IDs parameter is provided and not empty
	if len(IDs) > 0 {
		filter["_id"] = bson.M{"$in": IDs}
	}

	cursor, err := productCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(product *models.Product) error {
	_, err := productCollection.InsertOne(context.Background(), product)
	return err
}

func UpdateProduct(
	id, userID primitive.ObjectID,
	product *models.Product,
) (isSuccess bool, err error) {
	result, err := productCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id, "user_id": userID},
		bson.M{"$set": product},
	)
	if result.MatchedCount == 1 {
		isSuccess = true
	}
	return
}

func DeleteProduct(id, userID primitive.ObjectID) (isSuccess bool, err error) {
	result, err := productCollection.DeleteOne(
		context.Background(),
		bson.M{"_id": id, "user_id": userID},
	)
	if result.DeletedCount == 1 {
		isSuccess = true
	}
	return
}
