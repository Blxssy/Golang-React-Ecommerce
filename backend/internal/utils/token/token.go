package token

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	jwtKey = []byte(os.Getenv("JWT_KEY"))
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func GetNewTokens(db *gorm.DB, refreshToken string) (string, string, error) {
	id, err := VerifyToken(refreshToken)
	if err != nil {
		log.Fatal("Invalid referesh token")
	}

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			log.Println("User not found")
			return "", "", nil
		} else {
			log.Fatal(err)
		}
	}

	// TODO: Вынести TTL в конфиг
	accessToken, err := NewToken(user.ID, time.Hour*24)
	if err != nil {
		log.Fatal(err)
		return "", "", nil
	}

	refreshToken, err = NewToken(user.ID, time.Hour*24*7)
	if err != nil {
		log.Fatal(err)
		return "", "", nil
	}

	return accessToken, refreshToken, nil
}

func NewToken(userID uint, ttl time.Duration) (string, error) {
	expirationTime := time.Now().Add(ttl)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (uint, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, fmt.Errorf("invalid token signature")
		}
		return 0, err
	}

	if !token.Valid {
		log.Println("VerifyToken: invalid token")
		return 0, fmt.Errorf("VerifyToken: invalid token")
	}

	return claims.UserID, nil
}
