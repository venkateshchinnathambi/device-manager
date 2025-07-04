package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var publicKey *rsa.PublicKey

func InitPublicKey(path string) error {
	keyData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return errors.New("failed to parse PEM block")
	}

	pub, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	var ok bool
	publicKey, ok = pub.PublicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("not RSA public key")
	}

	log.Println("Loaded public key for RS256 JWT validation")
	return nil
}

type DeviceClaims struct {
	DeviceID string `json:"device_id"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenStr string) (*DeviceClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &DeviceClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*DeviceClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or claims")
	}

	return claims, nil
}
