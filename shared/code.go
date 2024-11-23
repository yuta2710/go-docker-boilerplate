package shared

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type AuthIdProvider interface {
	Encode(userId int) string
	Decode(authKey string) (int, error)
}

// Strategy Pattern for encode/decode auth keys
type Base64AuthIdProvider struct {
}

func (p *Base64AuthIdProvider) Encode(userId int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("user-%d", userId)))
}

func (p *Base64AuthIdProvider) Decode(authId string) (int, error) {
	fmt.Printf("Decoding authId: %s\n", authId)

	// Step 1: Base64 decode
	decoded, err := base64.StdEncoding.DecodeString(authId)
	if err != nil {
		return 0, fmt.Errorf("failed to decode authId: %v", err)
	}
	fmt.Printf("Decoded string: %s\n", string(decoded))

	// Step 2: Split the string into parts
	parts := strings.Split(string(decoded), "-")
	if len(parts) != 2 || parts[0] != "user" {
		return 0, fmt.Errorf("invalid authId format: %s", string(decoded))
	}
	fmt.Printf("Parts after split: %v\n", parts)

	// Step 3: Convert the ID part to an integer
	userId, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("failed to parse userId from authId: %v", err)
	}
	fmt.Printf("Parsed userId: %d\n", userId)

	return userId, nil
}
