package types

import (
	"encoding/hex"
	"strings"
)

func StringToBytes(hexString string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(hexString, "0x"))
}


