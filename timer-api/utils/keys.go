package utils

import (
	"crypto/ed25519"
	"encoding/base64"
	"log"
	"os"
	"strings"
)

var (
	JwtPublicKey  ed25519.PublicKey
	JwtPrivateKey ed25519.PrivateKey
)

func init() {
	pubkeyStr := strings.TrimSpace(os.Getenv("JWT_PUBLIC_KEY"))
	privkeyStr := strings.TrimSpace(os.Getenv("JWT_PRIVATE_KEY"))
	if pubkeyStr == "" {
		log.Fatal("JWT_PUBLIC_KEY must be set")
	}
	if privkeyStr == "" {
		log.Fatal("JWT_PRIVATE_KEY must be set")
	}

	pubBytes, err := base64.StdEncoding.DecodeString(pubkeyStr)
	if err != nil {
		log.Fatalf("Error decoding public key: %v", err)
	}
	privBytes, err := base64.StdEncoding.DecodeString(privkeyStr)
	if err != nil {
		log.Fatalf("Error decoding private: %v", err)
	}

	if len(pubBytes) != ed25519.PublicKeySize {
		log.Fatalf("Public key size is invalid: expected %d bytes, got %d bytes", ed25519.PublicKeySize, len(pubBytes))
	}
	if len(privBytes) != ed25519.PrivateKeySize {
		log.Fatalf("Public key size is invalid: expected %d bytes, got %d bytes", ed25519.PrivateKeySize, len(privBytes))
	}

	JwtPublicKey = ed25519.PublicKey(pubBytes)
	JwtPrivateKey = ed25519.PrivateKey(privBytes)
}
