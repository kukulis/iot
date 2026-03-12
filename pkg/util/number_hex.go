package util

func NumberHex(number int) string {
	return StrPadLeft(Dechex(int64(number)), 4, "0")
}
