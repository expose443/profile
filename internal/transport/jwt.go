package transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expiryTime := time.Now().Add(5 * time.Minute)
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
		cookie, err := r.Cookie("Token")
		if err != nil {
			if err == http.ErrNoCookie {
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
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
		if !tkn.Valid {
			ErrorHandler(w, http.StatusUnauthorized)
			return
		}
		fmt.Println("valid")
		next.ServeHTTP(w, r)
	})
}

func key(t *jwt.Token) (interface{}, error) {
	return secretKey, nil
}
