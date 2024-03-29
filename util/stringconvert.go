package util

import (
	"crypto/sha1"
	"encoding/base64"
	"strconv"
)

// UNUSED
// func sha1Bytes(pStrData string) []byte {
// 	hasher := sha1.New()
// 	hasher.Write([]byte(pStrData))
// 	return hasher.Sum(nil)
// }

// FromBase64e gets base-64 StdEncoding (with error)
func FromBase64e(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

// FromBase64 gets base-64 StdEncoding; ignores error.
func FromBase64(input string) []byte {
	result, _ := base64.StdEncoding.DecodeString(input)
	return result
}

// ToBase64 gets base-64 StdEncoding
func ToBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// ToUBase64 gets base-64 URLEncoding
func ToUBase64(input string) string {
	return base64.URLEncoding.EncodeToString([]byte(input))
}

// FromUBase64 gets base-64 URLEncoding
func FromUBase64(input string) string {
	result, _ := base64.URLEncoding.DecodeString(input)
	return string(result)
}

// BytesToBase64 gets base-64 StdEncoding
func BytesToBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Sha1String just gets SHA1 StdEncoding.
func Sha1String(pStrData string) string {
	hasher := sha1.New()
	hasher.Write([]byte(pStrData))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

// StrInt64 string to int helper
func StrInt64(pStrInput string) int64 {
	var err error
	var fpoop int64
	if fpoop, err = strconv.ParseInt(pStrInput, 10, 32); err != nil {
		return 0
	}
	return int64(fpoop)
}
