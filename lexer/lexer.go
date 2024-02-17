package lexer

import (
	"mlang/token"
)

type LEXER struct {
	input string
	curr  int
	peek  int
	char  byte
}

func New(input string) *LEXER {
	lexer := &LEXER{
		input: input,
	}
	lexer.read()
	return lexer
}

func (lexer *LEXER) read() {
	if lexer.peek >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.peek]
	}
	lexer.curr = lexer.peek
	lexer.peek = lexer.peek + 1
}

func (lexer *LEXER) pk() byte {
	if lexer.peek >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.peek]
	}
}

func (lexer *LEXER) parse_ident() string {
	curr := lexer.curr
	for ascii(lexer.char) {
		lexer.read()
	}
	return lexer.input[curr:lexer.curr]
}

func (lexer *LEXER) parse_digit() string {
	curr := lexer.curr
	for digit(lexer.char) {
		lexer.read()
	}
	return lexer.input[curr:lexer.curr]
}

func (lexer *LEXER) ignore() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\r' || lexer.char == '\n' {
		lexer.read()
	}
}

func (lexer *LEXER) Next() token.TOKEN {
	var TOKEN token.TOKEN
	lexer.ignore()
	switch lexer.char {
	case '=':
		if lexer.pk() == '=' {
			TOKEN = token.TOKEN{TYPE: token.EQ, LITERAL: "=="}
			lexer.read()
		} else {
			TOKEN = newToken(token.ASSIGN, lexer.char)
		}
	case '!':
		if lexer.pk() == '=' {
			TOKEN = token.TOKEN{TYPE: token.NOT_EQ, LITERAL: "!="}
			lexer.read()
		} else {
			TOKEN = newToken(token.BANG, lexer.char)
		}
	case ';':
		TOKEN = newToken(token.SEMICOLON, lexer.char)
	case '(':
		TOKEN = newToken(token.LPAREN, lexer.char)
	case ')':
		TOKEN = newToken(token.RPAREN, lexer.char)
	case ',':
		TOKEN = newToken(token.COMMA, lexer.char)
	case '+':
		TOKEN = newToken(token.PLUS, lexer.char)
	case '{':
		TOKEN = newToken(token.LBRACE, lexer.char)
	case '}':
		TOKEN = newToken(token.RBRACE, lexer.char)
	case '-':
		TOKEN = newToken(token.MINUS, lexer.char)

	case '/':
		TOKEN = newToken(token.SLASH, lexer.char)
	case '*':
		TOKEN = newToken(token.ASTERISK, lexer.char)
	case '<':
		TOKEN = newToken(token.LT, lexer.char)
	case '>':
		TOKEN = newToken(token.GT, lexer.char)
	case 0:
		TOKEN.LITERAL = ""
		TOKEN.TYPE = token.EOF
	default:
		if ascii(lexer.char) {
			TOKEN.LITERAL = lexer.parse_ident()
			TOKEN.TYPE = token.Lookup(TOKEN.LITERAL)
			return TOKEN
		} else if digit(lexer.char) {
			TOKEN.LITERAL = lexer.parse_digit()
			TOKEN.TYPE = token.INT
			return TOKEN
		} else {
			TOKEN = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.read()
	return TOKEN
}

func newToken(token_type token.TOKEN_TYPE, char byte) token.TOKEN {
	return token.TOKEN{TYPE: token_type, LITERAL: string(char)}
}

func ascii(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func digit(char byte) bool {
	return '0' <= char && char <= '9'
}
