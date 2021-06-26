package comma

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	buffer := bytes.Buffer{}
	i := 0
	if strings.HasPrefix(s, "-") {
		i++
	}
	for ; i < len(s); i += 3 {
		if i < len(s)-3 {
			buffer.WriteString(s[i : i+3])
			buffer.WriteByte(',')
		} else {
			buffer.WriteString(s[i:])
		}
	}
	return buffer.String()
}
