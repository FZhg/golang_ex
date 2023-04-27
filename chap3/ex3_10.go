package chap3

import "bytes"

func commaIterative(s string) string {
	commaPeriod := 3
	if len(s) <= commaPeriod {
		return s
	}

	var buf bytes.Buffer
	i := len(s) % 3
	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
