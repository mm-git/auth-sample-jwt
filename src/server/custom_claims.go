package server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"
)

const (
	Issuer            = "auth-sample-jwt"
	AccessToken       = "AccessToken"
	AccessExpireTime  = time.Second * 20
	RefreshToken      = "RefreshToken"
	RefreshExpireTime = time.Second * 60
)

type CustomClaims struct {
	UserId string `json:"uid,omitempty"`
	jwt.StandardClaims
}

// testで任意の時刻をセットできるように現在時刻を引数にしている
func CreateCustomClaim(userId string, tokenType string, now time.Time) *CustomClaims {
	var expire int64
	switch tokenType {
	case AccessToken:
		expire = now.Add(AccessExpireTime).Unix()
	case RefreshToken:
		expire = now.Add(RefreshExpireTime).Unix()
	}

	claims := &CustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expire,
			Id:        uuid.NewV4().String(),
			IssuedAt:  now.Unix(),
			Issuer:    Issuer,
			Subject:   tokenType,
		},
	}

	return claims
}

func (c CustomClaims) Valid() error {
	err := c.StandardClaims.Valid()

	if err != nil {
		return err
	}

	if c.Issuer != Issuer {
		return &jwt.ValidationError{
			Inner:  fmt.Errorf("bad issuer"),
			Errors: jwt.ValidationErrorIssuer,
		}
	}

	if c.Subject != AccessToken && c.Subject != RefreshToken {
		return &jwt.ValidationError{
			Inner:  fmt.Errorf("invalid token subject"),
			Errors: jwt.ValidationErrorClaimsInvalid,
		}
	}

	return nil
}

func (c CustomClaims) IsAccessToken() bool {
	return c.Subject == AccessToken
}

func (c CustomClaims) IsRefreshToken() bool {
	return c.Subject == RefreshToken
}
