package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

// HmacSha1ToHex to hex num
func HmacSha1ToHex(message []byte, secret string) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha1 to byte slice
func HmacSha1(message []byte, secret string) []byte {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write(message)
	return h.Sum(nil)
}

// VerifyHmacSha1 verify hmacsha1 result
func VerifyHmacSha1(message []byte, secret string, signature string) bool {

	sign := HmacSha1(message, secret)

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature))

	return hmac.Equal(sign, actual)
}
