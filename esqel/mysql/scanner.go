package mysql

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	uc "unicode"
)

var (
	eof = rune(0)
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return r
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) scanGlyphs() (tok Token, lit string) {
	buf := bytes.Buffer{}
	_, _ = buf.WriteRune(s.read())

	func() {
		for {
			r := s.read()
			switch {
			case r == eof:
				return
			case !uc.IsLetter(r) && !uc.IsDigit(r) && r != '_':
				s.unread()
				return
			default:
				_, _ = buf.WriteRune(r)
			}
		}
	}()

	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	}

	return IDENT, buf.String()
}

func (s *Scanner) scanSpace() (tok Token, lit string) {
	buf := bytes.Buffer{}
	_, _ = buf.WriteRune(s.read())

	func() {
		for {
			r := s.read()
			switch {
			case r == eof:
				return
			case !uc.IsSpace(r):
				s.unread()
				return
			default:
				_, _ = buf.WriteRune(r)
			}
		}
	}()

	return WS, buf.String()
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	r := s.read()

	if uc.IsSpace(r) {
		s.unread()
		return s.scanSpace()
	}

	if uc.IsLetter(r) {
		s.unread()
		return s.scanGlyphs()
	}

	switch r {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(r)
	case ',':
		return COMMA, string(r)
	}

	return ILLEGAL, string(r)
}
