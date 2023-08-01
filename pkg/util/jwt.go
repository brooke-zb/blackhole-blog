package util

import (
	"blackhole-blog/pkg/setting"
	"crypto/ed25519"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AccountClaims struct {
	jwt.RegisteredClaims
	Uid uint64 `json:"uid"`
}

var privateKey ed25519.PrivateKey

func initJwt() {
	privateKey = ed25519.NewKeyFromSeed([]byte(setting.Config.Server.Jwt.Secret))
}

func GenerateToken(uid uint64, expire time.Duration) (string, error) {
	ts := time.Now()
	claims := AccountClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(ts),
			ExpiresAt: jwt.NewNumericDate(ts.Add(expire)),
		},
	}

	// use ed25519 to sign token
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(privateKey)
}

func VerifyToken(rawToken string) (claims AccountClaims, err error) {
	token, err := jwt.ParseWithClaims(rawToken, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check sign method
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey.Public(), nil
	})
	if err != nil {
		return
	}
	return *token.Claims.(*AccountClaims), err
}
