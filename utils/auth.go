package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key should be in .env file
var secretKey = []byte("usdb396sf67ds6#783")

func GenerateToken(username, phone string) (string, error) {
	// Create a new Claims struct to hold the token claims
	claims := jwt.MapClaims{
		"user_name": username,
		"phone":     phone,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Set token expiration time (e.g., 24 hours from now)
	}

	// Create a new token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse and validate the JWT token using the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the "Authorization" header
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token string from the header
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if tokenString == authorizationHeader {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Verify the token
		token, err := VerifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Token is valid, pass it along to the next handler
		ctx := context.WithValue(r.Context(), "token", token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
