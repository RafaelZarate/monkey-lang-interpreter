package token

// TokenType is the uint representation of a type of token.
type TokenType uint8

const (
	// ILLEGAL signifies a token that we don't know about.
	ILLEGAL TokenType = iota
	// EOF stands for an end of file character.
	EOF

	// IDENT represents an indentation.
	IDENT

	// INT is an integer literal.
	INT

	// ASSIGN is the token that represents assignation.
	ASSIGN
	// PLUS is the addition operator token type.
	PLUS
	// MINUS is the substraction operator token type.
	MINUS
	// BANG is an exclamation mark operator.
	BANG
	// ASTERISK is literally an asterisk.
	ASTERISK
	// SLASH dfasdfasdf. TODO: fix this
	SLASH

	// LT represents the less than operator.
	LT
	// GT represents the greater than operator.
	GT

	// COMMA represents the literal comma symbol ,.
	COMMA
	// SEMICOLON represents the literal semicolon symbol ;.
	SEMICOLON

	// LPAREN is an opening parenthesis.
	LPAREN
	// RPAREN is an closing parenthesis.
	RPAREN
	// LBRACE is an opening brace.
	LBRACE
	// RBRACE is an closing brace.
	RBRACE

	// FUNCTION is the keyword for function declaration.
	FUNCTION
	// LET is the keyword for a variable declaration.
	LET
)

// Token represents a parsed token.
type Token struct {
	Type    TokenType
	Literal string
}

// NewToken creates a new instance of Token.
func NewToken(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

// LookupIdentifier returns a type for the given token literal.
func LookupIdentifier(ident string) TokenType {
	var tok TokenType
	switch ident {
	case "fn":
		tok = FUNCTION
	case "let":
		tok = LET
	default:
		tok = IDENT
	}

	return tok
}

func TokenTypeString(tokenType TokenType) string {
	var s string
	switch tokenType {
	case ILLEGAL:
		s = "ILLEGAL"
	case EOF:
		s = "EOF"
	case IDENT:
		s = "IDENT"
	case INT:
		s = "INT"
	case ASSIGN:
		s = "ASSIGN"
	case PLUS:
		s = "PLUS"
	case MINUS:
		s = "MINUS"
	case BANG:
		s = "BANG"
	case ASTERISK:
		s = "ASTERISK"
	case SLASH:
		s = "SLASH"
	case LT:
		s = "LT"
	case GT:
		s = "GT"
	case COMMA:
		s = "COMMA"
	case SEMICOLON:
		s = "SEMICOLON"
	case LPAREN:
		s = "LPAREN"
	case RPAREN:
		s = "RPAREN"
	case LBRACE:
		s = "LBRACE"
	case RBRACE:
		s = "RBRACE"
	case FUNCTION:
		s = "FUNCTION"
	case LET:
		s = "LET"
	}
	return s
}
