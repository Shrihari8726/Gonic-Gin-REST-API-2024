// services/user_service.go
package services

import (
	"context"
	"errors"

	"example.com/gonicginrestapi/config"
	"example.com/gonicginrestapi/models"
	"example.com/gonicginrestapi/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(config.GetMongoURI())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		utils.LogError(err)
	}
	userCollection = client.Database("testdb").Collection("users")
}

func CreateUser(user models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	_, err := userCollection.InsertOne(context.TODO(), user)
	return err
}

func AuthenticateUser(email, password string) (string, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Email)
	return token, err
}

func FetchAllUsers() ([]models.User, error) {
	var users []models.User
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &users)
	return users, err
}
