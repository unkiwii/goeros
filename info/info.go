package info

import "fmt"

const (
	// Name is the name of the executable
	Name = "eros"

	// Author is the name of the author of the executable
	Author = "Lucas Gabriel Sánchez"

	// AuthorMail is the email of the author of the executable
	AuthorMail = "unkiwii@gmail.com"
)

var (
	// Version is the version code
	Version = "NO-VERSION"

	// YearCompiled is the year in which this was compiled
	YearCompiled = "NO-YEAR"

	// MonthCompiled is the month in which this was compiled
	MonthCompiled = "NO-MONTH"

	// DayCompiled is the day in which this was compiled
	DayCompiled = "NO-DAY"
)

// DateCompiled returns the compiled date formatted as YEAR-MONTH-DAY
func DateCompiled() string {
	return fmt.Sprintf("%s-%s-%s", YearCompiled, MonthCompiled, DayCompiled)
}
