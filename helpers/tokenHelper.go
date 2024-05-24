package helpers

import (
	"log"
	"os"
	"time"

	"github.com/Maliud/GO_JWT_Authentication_with_Gin_Gonic/database"
	"github.com/dgrijalva/jwt-go/v4"
	"go.mongodb.org/mongo-driver/mongo"
)




type SignedDetails struct {
	Email string
	First_name string
	Last_name string
	Uid string
	User_type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string) {

	currentTime := time.Now().Local()
	expirationTime := currentTime.Add(24 * time.Hour)

	currentTimeRefresh := time.Now().Local()
	expirationTimeRefresh := currentTimeRefresh.Add(168 * time.Hour)

	claims := &SignedDetails{
		Email:  email,
		First_name: firstName,
		Last_name: lastName,
		Uid: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expirationTime),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expirationTimeRefresh),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken
}