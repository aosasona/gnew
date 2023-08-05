package lexer

import (
	"fmt"
	"gnew/token"
)

type Lexer struct {
	input        string
	position     int // current position
	readPosition int // next position
	char         byte

	lineNumber int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, lineNumber: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// peekChar returns the next character in the input without advancing the lexer to the next character
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) printPosition() {
	fmt.Printf("line %d\n", l.lineNumber)
}

func (l *Lexer) makeTokenWithDoubleLiteral(tokenType token.TokenType) token.Token {
	ch := l.char
	l.readChar()
	return l.makeToken(tokenType, string(ch)+string(l.char))
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			tk = l.makeTokenWithDoubleLiteral(token.EQ)
		} else {
			tk = l.makeToken(token.ASSIGN)
		}
	case '+':
		if l.peekChar() == '+' {
			tk = l.makeTokenWithDoubleLiteral(token.INCREMENT)
		} else {
			tk = l.makeToken(token.PLUS)
		}
	case '-':
		if l.peekChar() == '-' {
			tk = l.makeTokenWithDoubleLiteral(token.DECREMENT)
		} else {
			tk = l.makeToken(token.MINUS)
		}
	case '!':
		if l.peekChar() == '=' {
			tk = l.makeTokenWithDoubleLiteral(token.NOT_EQ)
		} else {
			tk = l.makeToken(token.BANG)
		}
	case '*':
		tk = l.makeToken(token.ASTERISK)
	case '/':
		tk = l.makeToken(token.FORWARD_SLASH)
	case '<':
		tk = l.makeToken(token.LT)
	case '>':
		tk = l.makeToken(token.GT)

	case ',':
		tk = l.makeToken(token.COMMA)
	case ';':
		tk = l.makeToken(token.SEMICOLON)
	case '(':
		tk = l.makeToken(token.LPAREN)
	case ')':
		tk = l.makeToken(token.RPAREN)
	case '{':
		tk = l.makeToken(token.LBRACE)
	case '}':
		tk = l.makeToken(token.RBRACE)
	case '[':
		tk = l.makeToken(token.LBRACKET)
	case ']':
		tk = l.makeToken(token.RBRACKET)

	case 0:
		tk = token.Token{Type: token.EOF, Literal: ""}

	default:
		if isLetter(l.char) {
			literal := l.readIdentifier()
			return l.makeToken(token.LookupIdentifier(literal), literal)
		} else if isDigit(l.char) {
			return l.makeToken(token.INT, l.readNumber())
		} else {
			tk = l.makeToken(token.ILLEGAL)
		}
	}

	l.readChar()
	return tk
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.lineNumber += 1
		}
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || char == '_'
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position

	// keep walking the input as long as it is still a letter
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readNumber() string {
	startPosition := l.position

	// keep walking the input as long as it is still a digit
	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) makeToken(tokenType token.TokenType, charOpt ...string) token.Token {
	char := string(l.char)
	if len(charOpt) > 0 {
		char = charOpt[0]
	}
	return token.Token{Type: tokenType, Literal: char}
}
