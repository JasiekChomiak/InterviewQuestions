// Link: https://techdevguide.withgoogle.com/resources/former-interview-question-compression-and-decompression/#!

package Google

import (
"strings"
)

func isDigit(r uint8) bool {
	return '0' <= r && r <= '9'
}

func parseNumber(s string, index int) (int, int) {
	value := 0
	bytesRead := 0

	for i := index; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}

		value *= 10
		value += (int)(s[i] - '0')
		bytesRead++
	}

	return value, bytesRead
}

func decompressAux(s string, startIndex int) (*strings.Builder, int) {
	sb := &strings.Builder{}
	i := startIndex

	for i < len(s) {
		if isDigit(s[i]) { // read the number, parse the inside and
			multiplier, multiplierLen := parseNumber(s, i)
			inside, insideLen := decompressAux(s, i+multiplierLen+1) // omit the '['
			sb.WriteString(strings.Repeat(inside.String(), multiplier))
			i += multiplierLen + insideLen + 2 // omit the '[' and ']'
		} else if s[i] == '[' { // in theory we should never find that, but just to make sure
			i++
		} else if s[i] == ']' {
			break
		} else { // a letter
			sb.WriteByte(s[i])
			i++
		}
	}

	return sb, i - startIndex
}

func decompress(s string) string {
	sb, _ := decompressAux(s, 0)
	return sb.String()
}


