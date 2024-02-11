package controller

import (
	webUser "belajar-golang-restful-api/model/web/users"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

func CreateToken(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var user webUser.User
	_ = json.NewDecoder(request.Body).Decode(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, error := token.SignedString(JwtKey)

	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(writter).Encode(webUser.JwtToken{Token: tokenString})

}

// ValidateMiddleware validates the JWT token.
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			log.Println(bearerToken, "line 44")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return JwtKey, nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(webUser.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(webUser.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(webUser.Exception{Message: "An authorization header is required"})
		}
	})
}
