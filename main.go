package main

import (
	"fmt"
	"time"
	_ "tripwala/routers"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

const (
	mySigningKey = "WOW,MuchShibe,ToDogge"
)

func ExampleNew(mySigningKey []byte) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}
func ExampleParse(myToken string, myKey string) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")
	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
	}
}
func main() {
	createdToken, err := ExampleNew([]byte(mySigningKey))
	if err != nil {
		fmt.Println("Creating token failed")
	}
	ExampleParse(createdToken, mySigningKey)
	beego.Run()

}
