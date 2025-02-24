package utils

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"log"
)

var (
	JwtPublicKey  ed25519.PublicKey
	JwtPrivateKey ed25519.PrivateKey
)

func init() {
	pubKeyB64 := GetEnvOrFallback("JWT_PUBLIC_KEY", "")
	privKeyB64 := GetEnvOrFallback("JWT_PRIVATE_KEY", "")
	if pubKeyB64 == "" || privKeyB64 == "" {
		log.Fatal("JWT keys are not set")
	}
	decodedPubKey, err := base64.StdEncoding.DecodeString(pubKeyB64)
	if err != nil {
		log.Fatalf("Failed to decode JWT_PUBLIC_KEY: %v", err)
	}
	decodedPrivKey, err := base64.StdEncoding.DecodeString(privKeyB64)
	if err != nil {
		log.Fatalf("Failed to decode JWT_PRIVATE_KEY: %v", err)
	}

	if len(pubKeyB64) != ed25519.PublicKeySize {
		log.Fatalf("Invalid public key size. Expected %d, got %d", ed25519.PublicKeySize, len(pubKeyB64))
	}
	if len(privKeyB64) != ed25519.PrivateKeySize {
		log.Fatalf("Invalid public key size. Expected %d, got %d", ed25519.PrivateKeySize, len(privKeyB64))
	}

	JwtPublicKey = ed25519.PublicKey(decodedPubKey)
	JwtPrivateKey = ed25519.PrivateKey(decodedPrivKey)

	fmt.Println("JWT keys loaded successfully")
}
