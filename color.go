package debugo

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	hexFormat      = "#%02x%02x%02x"
)

// genColor gen a color from input byte
func genColor(in []byte) (uint8, uint8, uint8) {
	h := hex.EncodeToString(in)
	h = "#" + strings.ToUpper(h[:6])

	var r, g, b uint8

	fmt.Sscanf(h, hexFormat, &r, &g, &b)

	return r, g, b
}
