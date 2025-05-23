package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `
    move_up
    move_down
    move_left
    move_right
    jump
    attack
    defend
    if (hero) {
        move_up
    }
    while (enemy) {
        attack
    }
    for (treasure; trap; hero) {
        defend
    }
    123
    + - * /
    && || !
    `

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{MoveUp, "move_up"},
		{MoveDown, "move_down"},
		{MoveLeft, "move_left"},
		{MoveRight, "move_right"},
		{Jump, "jump"},
		{Attack, "attack"},
		{Defend, "defend"},
		{If, "if"},
		{LeftParen, "("},
		{Identifier, "hero"},
		{RightParen, ")"},
		{LeftBrace, "{"},
		{MoveUp, "move_up"},
		{RightBrace, "}"},
		{While, "while"},
		{LeftParen, "("},
		{Identifier, "enemy"},
		{RightParen, ")"},
		{LeftBrace, "{"},
		{Attack, "attack"},
		{RightBrace, "}"},
		{For, "for"},
		{LeftParen, "("},
		{Identifier, "treasure"},
		{Semicolon, ";"},
		{Identifier, "trap"},
		{Semicolon, ";"},
		{Identifier, "hero"},
		{RightParen, ")"},
		{LeftBrace, "{"},
		{Defend, "defend"},
		{RightBrace, "}"},
		{Number, "123"},
		{Plus, "+"},
		{Minus, "-"},
		{Multiply, "*"},
		{Divide, "/"},
		{And, "&&"},
		{Or, "||"},
		{Not, "!"},
		{EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
