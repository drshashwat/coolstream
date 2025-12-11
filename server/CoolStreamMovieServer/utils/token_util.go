package utils

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/database"
)

type SignedDetails struct {
	FirstName string
	LastName  string
	Email     string
	UID       string
	Role      string
	jwt.RegisteredClaims
}

var (
	SECRET_KEY         string = os.Getenv("SECRET_KEY")
	SECRET_REFRESH_KEY        = os.Getenv("SECRET_REFRESH_KEY")

	userCollection *mongo.Collection = database.OpenCollection("users")
)

func GenerateAllTokens(email, firstName, lastName, role, userID string) (string, string, error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		UID:       userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "CoolStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshClaims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		UID:       userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "CoolStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(SECRET_REFRESH_KEY))
	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil
}

func UpdateAllTokens(userID, token, refershToken string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateData := bson.M{
		"$set": bson.M{
			"token":         token,
			"refresh_token": refershToken,
			"updated_at":    updateAt,
		},
	}
	_, err = userCollection.UpdateOne(ctx, bson.M{"user_id": userID}, updateData)
	return
}
