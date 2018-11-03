package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + literal
	IDENT = "IDENT"
	INT   = "INT"

	// Operator
	ASSIGN = "="
	PLUS   = "+"

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LBRANCE = "{"
	RBRANCE = "}"

	// Keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
