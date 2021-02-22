package diffmatchpatch

import (
	"regexp"
	"unicode"
)

// TextClassifier is an interface with methods for classifying runes and strings.
// Providing a custom implementation allows the application to use a different encoding than Unicode.
type TextClassifier interface {
	// IsAlphanumeric returns true if the given rune is a letter or digit.
	IsAlphanumeric(r rune) bool
	// IsWhitespace returns true if the given rune is a space, tab or similar whitespace character.
	IsWhitespace(r rune) bool
	// IsLinebreak returns true if the given rune is a linebreak character such as \r or \n.
	IsLinebreak(r rune) bool
	// BeginsWithBlankLine returns true if the given string begins with two consecutive linebreaks.
	BeginsWithBlankLine(s string) bool
	// EndsWithBlankLine returns true if the given string ends with two consecutive linebreaks.
	EndsWithBlankLine(s string) bool
}

// UnicodeTextClassifier is the default implementation of TextClassifier, meant for ordinary Unicode strings.
type UnicodeTextClassifier struct {
}

func NewUnicodeTextClassifier() UnicodeTextClassifier {
	return UnicodeTextClassifier{}
}

func (utc UnicodeTextClassifier) IsAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func (utc UnicodeTextClassifier) IsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func (utc UnicodeTextClassifier) IsLinebreak(r rune) bool {
	return r == '\r' || r == '\n'
}

var (
	blanklineStartRegex = regexp.MustCompile(`^\r?\n\r?\n`)
	blanklineEndRegex   = regexp.MustCompile(`\n\r?\n$`)
)

func (utc UnicodeTextClassifier) BeginsWithBlankLine(s string) bool {
	return blanklineStartRegex.MatchString(s)
}

func (utc UnicodeTextClassifier) EndsWithBlankLine(s string) bool {
	return blanklineEndRegex.MatchString(s)
}
