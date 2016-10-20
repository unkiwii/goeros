package lexer

type lexer struct {
	name    string
	input   string
	start   int
	pos     int
	width   int
	lexemes chan lexeme
}

type lexerState func(*lexer) lexerState

var startState = nil //TBD

func lex(name, input string) (*lexer, chan lexeme) {
	l := &lexer{
		name:    name,
		input:   input,
		lexemes: make(chan lexeme),
	}
	go l.run()
	return l, l.lexemes
}

func (l *lexer) run() {
	for state := startState; state != nil; {
		state = state(l)
	}
	close(l.lexemes)
}

func (l *lexer) emit(cls lexemeClass) {
	l.lexemes <- lexeme{cls, l.input[l.start:l.pos]}
	l.start = l.pos
}
