package lexer

type Lexer struct {
	name    string
	input   string
	start   int
	pos     int
	width   int
	lexemes chan lexeme
}

type lexerState func(*Lexer, lexerState) lexerState

type lexeme struct {
	typ lexemeType
	val string
}

const eof = -1

type lexemeType int

func (l *lexeme) String() string {
	switch l.typ {
	case ltEOF:
		return "EOF"
	case ltError:
		return fmt.Sprintf("error: %s", l.val)
	}
	return fmt.Sprintf("<%s, %s>", l.typ, l.val)
}

func (t *lexemeType) String() string {
	switch l.typ {
	case ltError:
		return "error"
	case ltEOF:
		return "eof"
	case ltIdentifier:
		return "identifier"
	case ltNumber:
		return "number"
	}
}

// lexeme types
const (
	ltError lexemeType = iota
	ltEOF
	ltIdentifier
	ltNumber
)

func New(name string) *Lexer {
	return &Lexer{name: name}
}

func (l *Lexer) Lex(input string) chan lexeme {
	l.input = input
	l.start = 0
	l.pos = 0
	l.width = 0
	go l.run()
	return l.lexemes
}

func (l *Lexer) run() {
	for state := lexIdentifier; state != nil; {
		state = state(l)
	}
	close(l.lexemes)
}

func (l *Lexer) emit(typ lexemeType) {
	l.lexemes <- lexeme{typ, l.input[l.start:l.pos]}
	l.start = l.pos
}

// errorf returns an error token and terminates the scan by passing back a nil
// pointer that will be the next state, terminating l.run.
func (l *Lexer) errorf(format string, args ...interface{}) lexerState {
	l.lexemes <- lexeme{ltError, fmt.Sprintf(format, args...)}
	return nil
}

// returns the next rune in the input
func (l *Lexer) next() (r int) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() int {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune
// if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// STATES
func lexUnit(l *Lexer, prevState lexerState) lexerState {
	//TODO
	return prevState
}

// EXAMPLE OF LEXING A NUMBER
func lexNumber(l *Lexer, prevState lexerState) lexerState {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if isAlphaNumeric(l.peek()) {
		l.next()
		return l.errorf("bad number syntax: %q",
			l.input[l.start:l.pos])
	}
	l.emit(ltNumber)
	return prevState
}
