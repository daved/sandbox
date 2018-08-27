package mysql

import (
	"fmt"
	"io"
	"strings"
)

// SelectStmt ...
type SelectStmt struct {
	Table  string
	Fields []string
}

// String ...
func (s *SelectStmt) String() string {
	f := strings.Join(s.Fields, ", ")

	return fmt.Sprintf("SELECT %s FROM %s", f, s.Table)
}

type buffer struct {
	tok Token
	lit string
	n   int
}

// Parser ...
type Parser struct {
	s   *Scanner
	buf buffer
}

// NewParser ...
func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
	}
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	p.buf.tok, p.buf.lit = p.s.Scan()

	return p.buf.tok, p.buf.lit
}

func (p *Parser) scanWords() (tok Token, lit string) {
	t, l := p.scan()
	if t == WS {
		t, l = p.scanWords()
	}

	return t, l
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

// Parse ...
func (p *Parser) Parse() (*SelectStmt, error) {
	stmt := &SelectStmt{}

	if t, l := p.scanWords(); t != SELECT {
		return nil, fmt.Errorf("found %q, expected SELECT", l)
	}

	for {
		t, l := p.scanWords()
		if t != IDENT && t != ASTERISK {
			return nil, fmt.Errorf("found %q, expected field", l)
		}

		stmt.Fields = append(stmt.Fields, l)

		if t, _ = p.scanWords(); t != COMMA {
			p.unscan()
			break
		}
	}

	if t, l := p.scanWords(); t != FROM {
		return nil, fmt.Errorf("found %q, expected FROM", l)
	}

	t, l := p.scanWords()
	if t != IDENT {
		return nil, fmt.Errorf("found %q, expected table", l)
	}

	stmt.Table = l

	return stmt, nil
}
