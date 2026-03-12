package util

import (
	"encoding/hex"
)

func Hex2Bin(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
