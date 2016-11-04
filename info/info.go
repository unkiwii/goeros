package info

import "fmt"

const (
	Name       = "eros"
	Author     = "Lucas Gabriel Sánchez"
	AuthorMail = "unkiwii@gmail.com"
)

// this are variables because are setted via compiler flags and can't be constants
var (
	Version       string = "NO-VERSION"
	YearCompiled  string = "NO-YEAR"
	MonthCompiled string = "NO-MONTH"
	DayCompiled   string = "NO-DAY"
)

func DateCompiled() string {
	return fmt.Sprintf("%s-%s-%s", YearCompiled, MonthCompiled, DayCompiled)
}
