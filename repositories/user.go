package repositories

import (
	"context"
	"time"

	"github.com/Nux-xader/ecommerce-management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserRepository(db *mongo.Database) {
	userCollection = db.Collection("users")
}

func IsUsernameTaken(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	count, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func IsEmailTaken(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	count, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateUser(user *models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return &user, err
}

func UserProfile(id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}
