package model

type Expenses struct {
	Id     string
	title  string
	amount float64
	note   string
	tags   []string
}
