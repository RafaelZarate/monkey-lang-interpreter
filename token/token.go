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
	EQ
	NOT_EQ

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
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

func (tt TokenType) String() string {
	switch tt {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case BANG:
		return "BANG"
	case ASTERISK:
		return "ASTERISK"
	case SLASH:
		return "SLASH"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case RETURN:
		return "RETURN"
	case EQ:
		return "EQ"
	case NOT_EQ:
		return "NOT_EQ"
	default:
		return "UNKNOWN"
	}
}

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
	case "true":
		tok = TRUE
	case "false":
		tok = FALSE
	case "if":
		tok = IF
	case "else":
		tok = ELSE
	case "return":
		tok = RETURN
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
	case TRUE:
		s = "TRUE"
	case FALSE:
		s = "FALSE"
	case IF:
		s = "IF"
	case ELSE:
		s = "ELSE"
	case RETURN:
		s = "RETURN"
	case EQ:
		s = "EQ"
	case NOT_EQ:
		s = "NOT_EQ"
	}
	return s
}
