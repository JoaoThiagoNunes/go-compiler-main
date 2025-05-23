package main

import (
	"fmt" 
	"strings"
)

func main() {

	input_l := `
	move_u
	if (hero) {
		attack
	}
	while (enemy) {
		defend
	}
	`
	fmt.Println("Análise Léxica:")
	// Criar um lexer e consumir os tokens
	l := NewLexer(input_l)
	for {
		tok := l.NextToken()
		if tok.Type == EOF {
			break
		}
		fmt.Printf("Token: %v, Literal: %s\n", tok.Type, tok.Literal)
	}


	
	fmt.Println(strings.Repeat("//", 25))
	// Criar um novo lexer para o parser (zera a posição)
	l2 := NewLexer(input_l)

	// Criar e rodar o parser com o novo lexer
	parser := NewParser(l2)
	fmt.Println("Análise Sintática:")
	parser.Parse()
}




