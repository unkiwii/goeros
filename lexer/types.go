package lexer

type lexeme struct {
	class lexemeClass
	value string
}

type lexemeClass int

const (
	clsError lexemeClass = iota
	clsEOF

	clsKeyword

	// any TitleCase word
	clsType

	// any camelCase word
	clsName

	// :=
	clsAssign

	// [ ]
	clsLeftBracket
	clsRightBracket

	// ( )
	clsLeftParen
	clsRightParen

	// { }
	clsLeftBrace
	clsRightBrace

	// : , .
	clsColon
	clsComma
	clsDot

	// a single space (32)
	clsSpace

	// any line starting with #
	clsComment

	// =>
	clsArrow

	// 1 45 4.34
	clsNumber

	// "any text"
	clsString
)
