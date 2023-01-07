package model

import "github.com/lib/pq"

type Expenses struct {
	Id     int            `json:"id"`
	Title  string         `json:"title"`
	Amount float64        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `json:"tags"`
}
