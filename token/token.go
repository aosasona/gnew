package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"

	// Operators
	ASSIGN        TokenType = "="
	PLUS          TokenType = "+"
	MINUS         TokenType = "-"
	BANG          TokenType = "!"
	ASTERISK      TokenType = "*"
	FORWARD_SLASH TokenType = "/"
	INCREMENT     TokenType = "++"
	DECREMENT     TokenType = "--"

	// Comparison
	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	LBRACKET  TokenType = "["
	RBRACKET  TokenType = "]"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"

	// Control flow
	RETURN TokenType = "RETURN"
	IF     TokenType = "IF"
	ELSE   TokenType = "ELSE"

	// Booleans
	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
)

func LookupIdentifier(identifier string) TokenType {
	if tk, ok := keywords[identifier]; ok {
		return tk
	}

	return IDENTIFIER
}
