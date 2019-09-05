package server

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

var (
	SignMethod = jwt.SigningMethodRS256
)

type JWTKeys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func ReadKeys() (*JWTKeys, error) {
	data, err := ioutil.ReadFile("./keys/private-key.pem")
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(data)
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadFile("./keys/public-key.pem")
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(data)
	if err != nil {
		return nil, err
	}

	return &JWTKeys{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

func (keys *JWTKeys) NewAccessToken(userId string, now time.Time) (string, error) {
	if keys == nil || keys.PrivateKey == nil {
		return "", errors.New("private key is not set")
	}
	if userId == "" {
		return "", errors.New("user id is empty")
	}

	claims := CreateCustomClaim(userId, AccessToken, now)

	token := jwt.NewWithClaims(SignMethod, claims)
	tokenString, err := token.SignedString(keys.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (keys *JWTKeys) NewRefreshToken(userId string, now time.Time) (string, error) {
	if keys == nil || keys.PrivateKey == nil {
		return "", errors.New("private key is not set")
	}
	if userId == "" {
		return "", errors.New("user id is empty")
	}

	claims := CreateCustomClaim(userId, RefreshToken, now)

	token := jwt.NewWithClaims(SignMethod, claims)
	tokenString, err := token.SignedString(keys.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (keys *JWTKeys) ParseToken(tokenString string) (*jwt.Token, error) {
	if keys == nil || keys.PrivateKey == nil {
		return nil, errors.New("private key is not set")
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, err := token.Method.(*jwt.SigningMethodRSA); !err {
			return nil, errors.New("unexpected signing method")
		}
		return keys.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errors.New("token is not valid")
	}

	return parsedToken, nil
}
