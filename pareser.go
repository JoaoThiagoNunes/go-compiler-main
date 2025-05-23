package main

import "fmt"

// Estrutura do Parser!
type Parser struct {
	lexer        *Lexer
	currentToken Token
	peekToken    Token
	errors       []string
}

// Criar um novo parser
func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l, errors: []string{}}

	// Ler dois tokens para inicializar currentToken e peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// Avança para o próximo token
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// Adiciona um erro sintático
func (p *Parser) addError(message string) {
	p.errors = append(p.errors, message)
}

// Análise sintática! = Principal, será chamado no main
// Válidar a gramática da QuestLang = Parseia os comandos da linguagem
func (p *Parser) Parse() {
	for p.currentToken.Type != EOF {
		p.parseStatement()
		p.nextToken()
	}

	// Exibir erros sintáticos, se existirem
	if len(p.errors) > 0 {
		fmt.Println("Erros encontrados na análise sintática:")
		for _, err := range p.errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Análise sintática concluída sem erros!")
	}
}

// Regras da gramática!!
// Identifica e processa comandos
func (p *Parser) parseStatement() {
	switch p.currentToken.Type {
	case MoveUp, MoveDown, MoveLeft, MoveRight, Jump, Attack, Defend:
		// Comandos válidos, apenas avançamos
		fmt.Printf("Comando válido: %s\n", p.currentToken.Literal)

	case If:
		p.parseIfStatement()

	case While:
		p.parseWhileStatement()

	case For:
		p.parseForStatement()

	default:
		if p.currentToken.Type != RightBrace { // Evita erro ao fechar blocos corretamente
			p.addError(fmt.Sprintf("Token inesperado: %s", p.currentToken.Literal))
		}
	}
}

func (p *Parser) parseIfStatement() {
	fmt.Println("Analisando comando if")

	// Esperamos um parêntese esquerdo depois do 'if'
	if !p.expectPeek(LeftParen) {
		return
	}
	
	// Avançamos para o conteúdo dentro dos parênteses
	p.nextToken()

	// Validamos que a condição é um dos identificadores permitidos
	if p.currentToken.Type != Identifier || 
		(p.currentToken.Literal != "hero" && p.currentToken.Literal != "enemy" &&
		p.currentToken.Literal != "treasure" && p.currentToken.Literal != "trap") {
		p.addError(fmt.Sprintf("Erro: Esperado 'hero', 'enemy', 'treasure' ou 'trap', mas encontrado '%s'", p.currentToken.Literal))
	}

	// Esperamos um parêntese direito após a condição
	if !p.expectPeek(RightParen) {
		return
	}
	
	// Esperamos uma chave esquerda para iniciar o bloco
	if !p.expectPeek(LeftBrace) {
		return
	}
	
	// Avançamos para o primeiro comando dentro do bloco
	p.nextToken()

	// Processa os comandos dentro do bloco
	for p.currentToken.Type != RightBrace && p.currentToken.Type != EOF {
		p.parseStatement()
		p.nextToken()
	}

	// Verifica se o bloco foi devidamente fechado
	if p.currentToken.Type != RightBrace {
		p.addError("Erro: Esperado '}' ao final do bloco 'if'")
	}
}

func (p *Parser) parseWhileStatement() {
	fmt.Println("Analisando comando while")

	// Esperamos um parêntese esquerdo depois do 'while'
	if !p.expectPeek(LeftParen) {
		return
	}
	
	// Avançamos para o conteúdo dentro dos parênteses
	p.nextToken()

	// Validamos que a condição é um dos identificadores permitidos
	if p.currentToken.Type != Identifier || 
		(p.currentToken.Literal != "hero" && p.currentToken.Literal != "enemy" &&
		p.currentToken.Literal != "treasure" && p.currentToken.Literal != "trap") {
		p.addError(fmt.Sprintf("Erro: Esperado 'hero', 'enemy', 'treasure' ou 'trap', mas encontrado '%s'", p.currentToken.Literal))
	}

	// Esperamos um parêntese direito após a condição
	if !p.expectPeek(RightParen) {
		return
	}
	
	// Esperamos uma chave esquerda para iniciar o bloco
	if !p.expectPeek(LeftBrace) {
		return
	}
	
	// Avançamos para o primeiro comando dentro do bloco
	p.nextToken()

	// Processa os comandos dentro do bloco
	for p.currentToken.Type != RightBrace && p.currentToken.Type != EOF {
		p.parseStatement()
		p.nextToken()
	}

	// Verifica se o bloco foi devidamente fechado
	if p.currentToken.Type != RightBrace {
		p.addError("Erro: Esperado '}' ao final do bloco 'while'")
	}
}

func (p *Parser) parseForStatement() {
	fmt.Println("Analisando comando for")

	// Esperamos um parêntese esquerdo depois do 'for'
	if !p.expectPeek(LeftParen) {
		return
	}
	
	// Inicialização
	p.nextToken()
	if p.currentToken.Type != Identifier {
		p.addError("Erro: Esperado um identificador na inicialização do 'for'")
	}
	
	if !p.expectPeek(Semicolon) {
		p.addError("Erro: Esperado ';' após a inicialização do 'for'")
		return
	}
	
	// Condição
	p.nextToken()
	if p.currentToken.Type != Identifier {
		p.addError("Erro: Esperado um identificador na condição do 'for'")
	}
	
	if !p.expectPeek(Semicolon) {
		p.addError("Erro: Esperado ';' após a condição do 'for'")
		return
	}
	
	// Incremento
	p.nextToken()
	if p.currentToken.Type != Identifier {
		p.addError("Erro: Esperado um identificador no incremento do 'for'")
	}
	
	// Esperamos um parêntese direito após os três componentes
	if !p.expectPeek(RightParen) {
		return
	}
	
	// Esperamos uma chave esquerda para iniciar o bloco
	if !p.expectPeek(LeftBrace) {
		return
	}
	
	// Avançamos para o primeiro comando dentro do bloco
	p.nextToken()

	// Processa os comandos dentro do bloco
	for p.currentToken.Type != RightBrace && p.currentToken.Type != EOF {
		p.parseStatement()
		p.nextToken()
	}

	// Verifica se o bloco foi devidamente fechado
	if p.currentToken.Type != RightBrace {
		p.addError("Erro: Esperado '}' ao final do bloco 'for'")
	}
}

func (p *Parser) expectPeek(expected TokenType) bool {
	if p.peekToken.Type == expected {
		p.nextToken()
		return true
	}
	
	// Convertendo o token type para string para exibição mais amigável
	var expectedStr string
	switch expected {
	case LeftParen:
		expectedStr = "("
	case RightParen:
		expectedStr = ")"
	case LeftBrace:
		expectedStr = "{"
	case RightBrace:
		expectedStr = "}"
	case Semicolon:
		expectedStr = ";"
	default:
		expectedStr = fmt.Sprintf("%v", expected)
	}
	
	p.addError(fmt.Sprintf("Erro: Esperado '%s', mas encontrado '%s'", expectedStr, p.peekToken.Literal))
	return false
}