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
	// "const": CONST
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators
	ASSIGN        = "="
	PLUS          = "+"
	MINUS         = "-"
	BANG          = "!"
	ASTERISK      = "*"
	FORWARD_SLASH = "/"
	LT            = "<"
	GT            = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"

	CONST = "CONST" // for later
)

func LookupIdentifier(identifier string) TokenType {
	if tk, ok := keywords[identifier]; ok {
		return tk
	}

	return IDENTIFIER
}
