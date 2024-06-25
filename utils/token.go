package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a new JWT token for the given user ID
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Local().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateQrToken(userId uint, cantidad int, puntos int, producto int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  userId,
			"cantidad": cantidad,
			"puntos":   puntos,
			"producto": producto,
			"exp":      time.Now().Local().Add(time.Minute * 5).Unix(),
		})
	tokenString, err := token.SignedString([]byte(SecretQR))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ExtractTokenID extracts the user ID from the JWT token
func ExtractTokenID(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return []byte(SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		//fmt.Println("usurio " + strconv.Itoa(int(userID)))
		return userID, nil
	}

	return 0, fmt.Errorf("invalid token")

}

// ExtractTokenID extracts the user ID from the JWT token
func ExtractQRTokenID(tokenString string) (uint, int, int, uint, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(SecretQR), nil
	})
	if err != nil {
		return 0, 0, 0, 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		cantidad := int(claims["cantidad"].(float64))
		puntos := int(claims["puntos"].(float64))
		producto := uint(claims["producto"].(float64))

		fmt.Println(cantidad, puntos)
		return userID, cantidad, puntos, producto, nil

	}

	//fmt.Errorf("invalid token")
	return 0, 0, 0, 0, errors.New("invalid token")

}

func TokenCookie(c *gin.Context) {
	// Get the token from the request header
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	// Set the token in the cookie
	c.SetCookie("token", tokenString, 0, "/", "localhost", false, true)
}
