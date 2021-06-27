package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"app/GoSample/logger"

	jwt "github.com/dgrijalva/jwt-go"
)

var JWT JwtWrapper

func init() {
	jwtKey := os.Getenv("JWTSecretKey")
	JWT = JwtWrapper{
		SecretKey: jwtKey,
		Issuer:    "AuthService",
		ExpirationHours: 2,
	}
}

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	UserId string
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(userId uint) (signedToken string, err error) {
	claims := &JwtClaim{
		UserId: fmt.Sprint(userId),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		logger.ErrorLog("An error occured while generating auth token - UserId:", userId, "- Error:", err.Error())
	}
	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		logger.ErrorLog("An error occered while parsing claims")
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		logger.ErrorLog("JWT is expired - Token:", signedToken)
		err = errors.New("JWT is expired")
	}
	return
}
