package pathsignature

import "strings"

// Fragments of the path based on the seperators contained within
// Example:
// path: 	  C:\example\path\to\something.exe
// signature: c:-ex-pa-to-so-ex-**-**
type PathSignature struct {
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
	switch r {
	case '/':
		return true
	case '\\':
		return true
	case ' ':
		return true
	case '.':
		return true
	}

	return false
}

func IsWildcard(r rune) bool {
	return r == '*'
}

func Create(str string) PathSignature {
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

func CreateReverse(str string) PathSignature {
	signature := PathSignature{}

	for i := 0; i < 16; i++ {
		signature.Runes[i] = '*'
	}

	offset := 0
	lastWasSeperator := false
	for i := len(str) - 1; i >= 0 && offset < 16; i-- {
		r := rune(str[i])

		runeIsWildcard := IsWildcard(r)
		runeIsSeperator := IsSeperator(r)
		if runeIsWildcard {
			return signature
		}

		if runeIsSeperator || i == 0 {
			if lastWasSeperator {
				continue
			}
			lastWasSeperator = true
			var a rune = ' '
			var b rune = ' '

			if i == 0 && !runeIsSeperator {
				if i < len(str) {
					a = rune(str[i])
				}

				if i+1 < len(str) {
					b = rune(str[i+1])
				}
			} else {
				if i+1 < len(str) {
					a = rune(str[i+1])
				}

				if i+2 < len(str) {
					b = rune(str[i+2])
				}
			}

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

			offset += 2
			continue
		}

		lastWasSeperator = false
	}

	return signature
}
