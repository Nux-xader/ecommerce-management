package repositories

import (
	"context"

	"github.com/Nux-xader/ecommerce-management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection

func InitOrderRepository(db *mongo.Database) {
	orderCollection = db.Collection("orders")
}

func GetAllOrders(userID primitive.ObjectID) ([]models.Order, error) {
	var orders []models.Order

	cursor, err := orderCollection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func AddOrder(order models.Order) error {
	_, err := orderCollection.InsertOne(context.Background(), order)
	return err
}

func SetOrderStatus(id, userID primitive.ObjectID, status string) (isSuccess bool, err error) {
	result, err := orderCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id, "user_id": userID},
		bson.M{"$set": bson.M{"status": status}},
	)
	if result.MatchedCount == 1 {
		isSuccess = true
	}
	return
}

func IsCompletedOrder(id, userID primitive.ObjectID) (isCompleted bool, err error) {
	matchedCount, err := orderCollection.CountDocuments(
		context.Background(),
		bson.M{"_id": id, "user_id": userID, "status": models.OrderStatusCompleted},
	)
	if err != nil {
		return
	}
	if matchedCount > 0 {
		isCompleted = true
	}
	return
}
