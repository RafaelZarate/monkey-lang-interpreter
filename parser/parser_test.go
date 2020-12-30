package parser

import (
	"testing"

	"github.com/rafaelzarate/monkey-lang-interpreter/ast"
	"github.com/rafaelzarate/monkey-lang-interpreter/lexer"
)

func TestLetStatemts(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statemets does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expctedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if stmt.TokenLiteral() != "let" {
			t.Fatalf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		}
		letStmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("stmt not *ast.LetStatement. got=%T", stmt)
		}
		if letStmt.Name.Value != tt.expctedIdentifier {
			t.Fatalf("letStmt.Name.Value not '%s'. got=%s", tt.expctedIdentifier, letStmt.Name.Value)
		}
		if letStmt.Name.TokenLiteral() != tt.expctedIdentifier {
			t.Fatalf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
				tt.expctedIdentifier, letStmt.Name.TokenLiteral())
		}
	}
}

func TestReturnStatemts(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statemets does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
