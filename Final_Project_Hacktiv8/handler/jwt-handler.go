package handler

import (
	"fmt"
	"golang-crud-sql/helper"
	"golang-crud-sql/model"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var cfg *model.Config

func GenerateJWT(id int, username string, age int) (string, error) {
	cfg, err := helper.GetConfig()

	if err != nil {
		return err.Error(), err
	}

	var mySigningKey = []byte(cfg.JWT.Secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = username
	claims["age"] = age
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cfg, err := helper.GetConfig()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if strings.HasSuffix(r.URL.Path, "/users/register") ||
			strings.HasSuffix(r.URL.Path, "/users/login") {
			handler.ServeHTTP(w, r)
			return
		}

		clientToken := r.Header.Get("Authorization")
		if clientToken == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Write(helper.CreateErrorResponse("no authorization header provided"))
			return
		}

		splitToken := strings.Split(clientToken, "Bearer ")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(helper.CreateErrorResponse("invalid Token"))
			return
		}

		accessToken := splitToken[1]
		if len(accessToken) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(helper.CreateErrorResponse("invalid token"))
			return
		}

		var mySigningKey = []byte(cfg.JWT.Secret)

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(helper.CreateErrorResponse("your token has been expired!"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["id"]
			r.Header.Set("userId", fmt.Sprintf("%v", userId))
			handler.ServeHTTP(w, r)
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(helper.CreateErrorResponse("unauthorized"))
			return
		}
	})
}
