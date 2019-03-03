package ua

//go:generate stringer -type Token token.go
type Token uint

const (
	Invalid Token = iota
	Literal
	Symbol
	EOF
)
