package debugo

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	hexFormat      = "#%02x%02x%02x"
	hexShortFormat = "#%1x%1x%1x"
	hexToRGBFactor = 17
)

// genColor gen a color from input byte
func genColor(in []byte) (uint8, uint8, uint8) {
	h := hex.EncodeToString(in)
	h = "#" + strings.ToUpper(h[:6])

	var r, g, b uint8

	if len(h) == 4 {
		fmt.Sscanf(h, hexShortFormat, &r, &g, &b)
		r *= hexToRGBFactor
		g *= hexToRGBFactor
		b *= hexToRGBFactor
	} else {
		fmt.Sscanf(h, hexFormat, &r, &g, &b)
	}

	return r, g, b
}
