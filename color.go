package debugo

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const (
	hexFormat = "%02x%02x%02x"
)

func md5Hash(in []byte) string {
	h := md5.New()
	h.Write(in)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// genColor gen a color from input byte
func genColor(in []byte) (uint8, uint8, uint8) {
	h := md5Hash(in)
	h = strings.ToUpper(h[10:16])
	var r, g, b uint8

	fmt.Sscanf(h, hexFormat, &r, &g, &b)

	return r, g, b
}
