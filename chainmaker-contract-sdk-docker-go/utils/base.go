package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GetSHA256String(content []string) (result string) {
	hash := sha256.New()
	hash.Write([]byte(strings.Join(content, "")))
	return hex.EncodeToString(hash.Sum(nil))
}
