package token

type TOKEN_TYPE string
type TOKEN struct {
	TYPE    TOKEN_TYPE
	LITERAL string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	ASSIGN     = "="
	PLUS       = "+"
	COMMA      = ","
	SEMICOLON  = ";"
	LPAREN     = "("
	RPAREN     = ")"
	LBRACE     = "{"
	RBRACE     = "}"
	FUNCTION   = "FUNCTION"
	LET        = "LET"
	CONST      = "CONST"
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"
	LT         = "<"
	GT         = ">"
	TRUE       = "TRUE"
	FALSE      = "FALSE"
	IF         = "IF"
	ELSE       = "ELSE"
	RETURN     = "RETURN"
	EQ         = "=="
	NOT_EQ     = "!="
)

var Keywords = map[string]TOKEN_TYPE{
	"fn":     FUNCTION,
	"let":    LET,
	"const":  CONST,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func Lookup(identifier string) TOKEN_TYPE {
	if token, ok := Keywords[identifier]; ok {
		return token
	}

	return IDENTIFIER
}
