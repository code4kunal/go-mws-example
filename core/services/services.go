package services


import (
"fmt"
"github.com/dgrijalva/jwt-go"
"time"
"../models"
)

var tokenEncodeString string = "something"

func createToken(user models.UserLoginDetails) (string, error) {
	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// set some claims
	token.Claims["username"] = user.UserName;
	token.Claims["password"] = user.Password;
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Create the token

	//Sign and get the complete encoded token as string
	return (token.SignedString(tokenEncodeString))
}

func parseToken(unparsedToken string, myKey string) (bool, string) {
	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")
	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
	}
}