package mysql

// Token ...
type Token int

// Tokens ...
const (
	// Special
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT

	// Chars
	ASTERISK
	COMMA

	// Keywords
	SELECT
	FROM
	AUTOINC
)
