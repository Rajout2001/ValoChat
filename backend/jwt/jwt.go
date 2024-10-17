// Package jwt provides functions for creating and verifying JSON Web Tokens (JWT).
package jwt

import (
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

/*
	Source d'inspiration: https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60
*/

// secretKey is the secret key used for signing and verifying JWTs.
var secretKey = []byte("secret-key")

// CreateToken creates a new JWT token with the given ID and expiration time.
// It returns the token string, expiration time in Unix format, and any error encountered.
func CreateToken(id string) (string, int64, error) {
	// Set the expiration time to 3 days from now.
	exp := time.Now().Add(time.Hour * 72).Unix()

	// Create a new JWT token with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": exp,
		})

	// Sign the token with the secret key.
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}

// VerifyToken verifies the authenticity and validity of the given token string.
// It returns the token claims and a boolean indicating whether the token is valid.
func VerifyToken(tokenString string) (jwt.MapClaims, bool) {
	// Parse the token string and verify it using the secret key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Handle parsing errors.
	if err != nil {
		return nil, false
	}

	// Check if the token is valid and retrieve its claims.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		// Log an error if the token is invalid.
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
