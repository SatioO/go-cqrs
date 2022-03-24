package marshaler

import (
	"fmt"
	"strings"
)

func FullyQualifiedStructName(v any) string {
	s := fmt.Sprintf("%T", v)
	s = strings.TrimLeft(s, "*")
	return s
}
