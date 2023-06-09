package transport

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

const keyUser = "username"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expiryTime := time.Now().Add(120 * time.Minute)
	claims := Claims{
		Username: username,
	}
	claims.ExpiresAt = expiryTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (h *Handlers) MiddlewareJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		fmt.Println(r.Cookies())
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println(5)
				ErrorHandler(w, http.StatusUnauthorized)
				return
			}
			ErrorHandler(w, http.StatusBadRequest)
			return
		}
		tokenStr := cookie.Value
		claims := Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, &claims, key)
		fmt.Println("CLAIMS", claims)
		if err != nil {
			fmt.Println(3)
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
		if !tkn.Valid {
			fmt.Println(4)
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), keyUser, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func key(t *jwt.Token) (interface{}, error) {
	return secretKey, nil
}
