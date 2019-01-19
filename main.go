package main

import (
	// "fmt"
	// "github.com/dgrijalva/jwt-go"
	// "time"
	"log"
	"net/http"
)

func main() {
	// now := time.Now()
	// tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
	// 	UserName: "soham",
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: now.Add(1 * time.Minute).Unix(),
	// 		IssuedAt:  now.Unix(),
	// 	},
	// })
	// tknStr, _ := tkn.SignedString(jwtKey)

	// fmt.Println(tknStr)

	// claims := &MyClaims{}
	// tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return "seee", nil
	// })

	// fmt.Println("Decoded: ", tkn, err, claims)

	// "Signin" and "Signup" are handler that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))

}
