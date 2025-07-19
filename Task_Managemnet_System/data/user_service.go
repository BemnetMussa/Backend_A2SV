package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("jwt_secrete_code_temporary") // replace --<
var UserCollection *mongo.Collection

func SetUserCollection(c *mongo.Collection) {
	UserCollection = c
}

func RegisterUser( name string, email string, password string) error {
	// Check if user already exists in MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser models.User
	err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Count how many users exist
	count, err := UserCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return errors.New("failed to count users")
	}

	role := "user"
	if count == 0 {
		role = "admin" // First user = admin
	}

	// Insert into MongoDB
	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	fmt.Println(newUser)
	_, err = UserCollection.InsertOne(ctx, newUser)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func LoginUser(email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Email,
		"email":   user.Email,
		"role": user.Role,
	})

	

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return jwtToken, nil
}

func PromoteUserByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("failed to promote user")
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
