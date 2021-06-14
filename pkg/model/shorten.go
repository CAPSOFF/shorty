package model

type ShortySpec struct {
	URL           string
	StartDate     string
	LastSeenDate  string
	RedirectCount int
}

var RegexFormula = `^[0-9a-zA-Z_]{6}$`
