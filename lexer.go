package main

type TokenType int

const (
	// Comandos de Movimento
	MoveUp    TokenType = iota // iota 0
	MoveDown                   // iota 1
	MoveLeft                   // iota 2
	MoveRight                  // iota 3

	// Comandos de Ação
	Jump
	Attack
	Defend

	// Operadores Aritméticos
	Plus
	Minus
	Multiply
	Divide

	// Números e Identificadores
	Number
	Identifier

	// Palavras-chave
	If
	Else
	While
	For

	// Símbolos
	LeftParen  // (
	RightParen // )
	LeftBrace  // {
	RightBrace // }
	Semicolon  // ;

	// Operadores Lógicos
	And
	Or
	Not

	// Outros
	EOF
	Illegal
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// Imagine como um leitor que vai caractere por caractere:
// input: "move_up"
//                 ↓
// posição inicial: m|o|v|e|_|u|p
// position: 0     ^
// readPosition: 1   ^
// ch: 'm'

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input} // cria novo Lexer
	l.readChar()              // lê o primeiro caractere
	return l                  // retorna o Lexer
}

// Quando criamos um novo lexer com o texto "move_up"
// lexer := NewLexer("move_up")

// O estado inicial será:
// lexer.input = "move_up"
// lexer.position = 0
// lexer.readPosition = 1
// lexer.ch = 'm'

// Quando chamamos NextToken(), ele vai ler "move_up" como um único token:
// token := lexer.NextToken()
// token será:
// Token {
//     Type: MoveUp,
//     Literal: "move_up"
// }

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token{Type: Illegal, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token{Type: Illegal, Literal: string(l.ch)}
		}
	case '+':
		token = Token{Type: Plus, Literal: string(l.ch)}
	case '-':
		token = Token{Type: Minus, Literal: string(l.ch)}
	case '*':
		token = Token{Type: Multiply, Literal: string(l.ch)}
	case '/':
		token = Token{Type: Divide, Literal: string(l.ch)}
	case '(':
		token = Token{Type: LeftParen, Literal: string(l.ch)}
	case ')':
		token = Token{Type: RightParen, Literal: string(l.ch)}
	case '{':
		token = Token{Type: LeftBrace, Literal: string(l.ch)}
	case '}':
		token = Token{Type: RightBrace, Literal: string(l.ch)}
	case ';':
		token = Token{Type: Semicolon, Literal: string(l.ch)}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			token = Token{Type: And, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token{Type: Illegal, Literal: string(l.ch)}
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			token = Token{Type: Or, Literal: string(ch) + string(l.ch)}
		} else {
			token = Token{Type: Illegal, Literal: string(l.ch)}
		}
	case '!':
		token = Token{Type: Not, Literal: string(l.ch)}
	case 0:
		token = Token{Type: EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			token.Literal = l.readIdentifier()
			token.Type = l.lookupIdent(token.Literal)
			return token
		} else if isDigit(l.ch) {
			token.Type = Number
			token.Literal = l.readNumber()
			return token
		} else {
			token = Token{Type: Illegal, Literal: string(l.ch)}
		}
	}

	l.readChar()
	return token
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) lookupIdent(ident string) TokenType {
	switch ident {
	case "move_up":
		return MoveUp
	case "move_down":
		return MoveDown
	case "move_left":
		return MoveLeft
	case "move_right":
		return MoveRight
	case "jump":
		return Jump
	case "attack":
		return Attack
	case "defend":
		return Defend
	case "if":
		return If
	case "else":
		return Else
	case "while":
		return While
	case "for":
		return For
	case "hero":
		return Identifier
	case "enemy":
		return Identifier
	case "treasure":
		return Identifier
	case "trap":
		return Identifier
	default:
		return Identifier
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
