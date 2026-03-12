package util

import (
	"fmt"
	"strconv"
)

func Hex2Bin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%016b", ui), nil
}
