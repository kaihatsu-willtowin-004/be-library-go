package models

import "time"

type Loans struct {
	ID          int    `json:"id"`
	BOOK_ID     int `json:"book_id"`
	MEMBER_ID   int `json:"member_id"`
	STAFF_ID    int `json:"staff_id"`
	BORROW_DATE time.Time  `json:"borrow_date"`
	DUE_DATE     time.Time `json:"due_date"`
	RETURN_DATE  time.Time `json:"return_date"`
	STATUS   string `json:"status"`
}