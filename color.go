package debugo

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	hexFormat = "%02x%02x%02x"
)

// genColor gen a color from input byte
func genColor(in []byte) (uint8, uint8, uint8) {
	h := hex.EncodeToString(in)
	if len(h) < 6 {
		h += strings.Repeat("0", 6-len(h))
	}

	h = strings.ToUpper(h[len(h)-6:])

	var r, g, b uint8

	fmt.Sscanf(h, hexFormat, &r, &g, &b)

	return r, g, b
}
