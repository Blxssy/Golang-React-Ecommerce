package token

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const accessTokenDuration = time.Minute * 15
const refreshTokenDuration = time.Hour * 24 * 7

func InitJWTKey() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

var jwtKey []byte

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GetNewTokens(userID uint) (string, string, error) {
	// TODO: Вынести TTL в конфиг
	accessToken, err := NewToken(userID, accessTokenDuration)
	if err != nil {
		log.Fatal(err)
		return "", "", nil
	}

	refreshToken, err := NewToken(userID, refreshTokenDuration)
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

func UpdateToken(refreshTokenString string) (string, string, error) {
	claims := &Claims{}
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", "", err
	}

	if !refreshToken.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	return GetNewTokens(claims.UserID)
}

func ValidateToken(refreshToken string) bool {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false
	}
	return true
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
