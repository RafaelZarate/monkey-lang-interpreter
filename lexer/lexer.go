package lexer

import "github.com/rafaelzarate/monkey-lang-interpreter/token"

// Lexer defines methods necessary to read an input and turn them into tokens.
type Lexer interface {
	NextToken() token.Token
}

// StringLexer implements Lexer with a string as the input
type StringLexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New returns a Lexer implementation.
func New(input string) *StringLexer {
	l := &StringLexer{input: input}
	l.readChar()
	return l
}

// NextToken parses the next character into a token.
func (l *StringLexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = token.NewToken(token.PLUS, string(l.ch))
	case '-':
		tok = token.NewToken(token.MINUS, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = token.NewToken(token.BANG, string(l.ch))
		}
	case '/':
		tok = token.NewToken(token.SLASH, string(l.ch))
	case '*':
		tok = token.NewToken(token.ASTERISK, string(l.ch))
	case '<':
		tok = token.NewToken(token.LT, string(l.ch))
	case '>':
		tok = token.NewToken(token.GT, string(l.ch))
	case ';':
		tok = token.NewToken(token.SEMICOLON, string(l.ch))
	case '(':
		tok = token.NewToken(token.LPAREN, string(l.ch))
	case ')':
		tok = token.NewToken(token.RPAREN, string(l.ch))
	case ',':
		tok = token.NewToken(token.COMMA, string(l.ch))
	case '{':
		tok = token.NewToken(token.LBRACE, string(l.ch))
	case '}':
		tok = token.NewToken(token.RBRACE, string(l.ch))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = token.NewToken(token.ILLEGAL, string(l.ch))
	}

	l.readChar()
	return tok
}

func (l *StringLexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *StringLexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *StringLexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *StringLexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *StringLexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
