package middleware

import (
	"mvp-shop-backend/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateToken is a JWT New With Claims SigningMethodHS256
func GenerateToken(customer models.CustomerClaims) (string, error) {
	jwtExpired := os.Getenv("JWT_EXPIRED")
	uom := jwtExpired[len(jwtExpired)-1:]
	expiredToken, _ := strconv.Atoi(strings.Replace(jwtExpired, uom, "", -1))

	timeHourMinutes := time.Hour
	if uom == "d" {
		timeHourMinutes = time.Hour * 24
	}

	expiredAt := time.Now().Add(timeHourMinutes * time.Duration(expiredToken)).Unix()
	claims := models.CustomerClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    "dbo-is-backend",
			Subject:   "customer",
		},
		ID:     customer.ID,
		Name:   customer.Name,
		Email:  customer.Email,
		Status: customer.Status,
	}

	var signingKey = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JwtClaim is a JWT Parse With Claims Token
func JwtClaim(token string) (customer *models.CustomerClaims, err error) {
	var signingKey = []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	id := claims["id"].(string)
	name := claims["name"].(string)
	email := claims["email"].(string)
	status := claims["status"].(string)
	customer = &models.CustomerClaims{
		ID:     id,
		Name:   name,
		Email:  email,
		Status: models.Status(status),
	}
	return
}
