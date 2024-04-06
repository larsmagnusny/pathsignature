package pathsignature

import "strings"

type PathSignature struct {
	//
	// **-**-**-**-**-**-**-**
	Runes [16]rune
}

func (sig PathSignature) ToString() string {
	var sb strings.Builder

	for i := 0; i < 16; i += 2 {
		sb.WriteRune(sig.Runes[i])
		sb.WriteRune(sig.Runes[i+1])

		if i < 14 {
			sb.WriteRune('-')
		}
	}

	return sb.String()
}

func IsSeperator(r rune) bool {
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

func IsWildcard(r rune) bool {
	if r == '*' {
		return true
	}

	return false
}

func CreatePathSignature(str string) PathSignature {
	signature := PathSignature{}

	for i := 0; i < 16; i++ {
		signature.Runes[i] = '*'
	}

	offset := 0
	lastWasSeperator := false
	for i := 0; i < len(str)-2 && offset < 16; i++ {
		r := rune(str[i])
		if IsWildcard(r) {
			return signature
		}

		if IsSeperator(r) {
			if lastWasSeperator {
				continue
			}
			lastWasSeperator = true

			a := rune(str[i+1])
			b := rune(str[i+2])

			if IsWildcard(a) {
				signature.Runes[0+offset] = rune('*')
				signature.Runes[1+offset] = rune('*')
				return signature
			}

			if !IsSeperator(a) {
				signature.Runes[0+offset] = rune(a)
			}

			if IsWildcard(b) {
				signature.Runes[1+offset] = rune('*')
			}
			if !IsSeperator(b) {
				signature.Runes[1+offset] = rune(b)
			} else {
				signature.Runes[1+offset] = rune(' ')
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

func CreateReversePathSignature(str string) PathSignature {
	signature := PathSignature{}

	for i := 0; i < 16; i++ {
		signature.Runes[i] = '*'
	}

	offset := 0
	lastWasSeperator := false
	for i := len(str) - 1; i >= 2 && offset < 16; i-- {
		r := rune(str[i])
		if IsWildcard(r) {
			return signature
		}

		if IsSeperator(r) {
			if lastWasSeperator {
				continue
			}
			lastWasSeperator = true

			a := rune(str[i+1])
			b := rune(str[i+2])

			if IsWildcard(a) {
				signature.Runes[0+offset] = rune('*')
				signature.Runes[1+offset] = rune('*')
				return signature
			}

			if !IsSeperator(a) {
				signature.Runes[0+offset] = rune(a)
			}

			if IsWildcard(b) {
				signature.Runes[1+offset] = rune('*')
			}
			if !IsSeperator(b) {
				signature.Runes[1+offset] = rune(b)
			} else {
				signature.Runes[1+offset] = rune(' ')
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
