package pathsignature

import "strings"

type path_signature struct {
	//
	// **-**-**-**-**-**-**-**
	runes [16]rune
}

func (sig path_signature) toString() string {
	var sb strings.Builder

	for i := 0; i < 16; i += 2 {
		sb.WriteRune(sig.runes[i])
		sb.WriteRune(sig.runes[i+1])

		if i < 14 {
			sb.WriteRune('-')
		}
	}

	return sb.String()
}

func isSeperator(r rune) bool {
	if r == rune('/') {
		return true
	}

	if r == rune('\\') {
		return true
	}

	if r == rune(' ') {
		return true
	}

	return false
}

func isWildcard(r rune) bool {
	if r == '*' {
		return true
	}

	return false
}

func createPathSignature(str string) path_signature {
	signature := path_signature{}

	for i := 0; i < 16; i++ {
		signature.runes[i] = '*'
	}

	offset := 0
	lastWasSeperator := false
	for i := 0; i < len(str)-2 && offset < 16; i++ {
		r := rune(str[i])
		if isWildcard(r) {
			return signature
		}

		if isSeperator(r) {
			if lastWasSeperator {
				continue
			}
			lastWasSeperator = true

			a := rune(str[i+1])
			b := rune(str[i+2])

			if isWildcard(a) {
				signature.runes[0+offset] = rune('*')
				signature.runes[1+offset] = rune('*')
				return signature
			}

			if !isSeperator(a) {
				signature.runes[0+offset] = rune(a)
			}

			if isWildcard(b) {
				signature.runes[1+offset] = rune('*')
			}
			if !isSeperator(b) {
				signature.runes[1+offset] = rune(b)
			} else {
				signature.runes[1+offset] = rune(' ')
				offset += 2
				continue
			}

			i++
			offset += 2
			continue
		}

		lastWasSeperator = false
	}

	return signature
}

func createReversePathSignature(str string) path_signature {
	signature := path_signature{}

	for i := 0; i < 16; i++ {
		signature.runes[i] = '*'
	}

	offset := 0
	lastWasSeperator := false
	for i := len(str) - 1; i >= 2 && offset < 16; i-- {
		r := rune(str[i])
		if isWildcard(r) {
			return signature
		}

		if isSeperator(r) {
			if lastWasSeperator {
				continue
			}
			lastWasSeperator = true

			a := rune(str[i+1])
			b := rune(str[i+2])

			if isWildcard(a) {
				signature.runes[0+offset] = rune('*')
				signature.runes[1+offset] = rune('*')
				return signature
			}

			if !isSeperator(a) {
				signature.runes[0+offset] = rune(a)
			}

			if isWildcard(b) {
				signature.runes[1+offset] = rune('*')
			}
			if !isSeperator(b) {
				signature.runes[1+offset] = rune(b)
			} else {
				signature.runes[1+offset] = rune(' ')
				offset += 2
				continue
			}

			i--
			offset += 2
			continue
		}

		lastWasSeperator = false
	}

	return signature
}
