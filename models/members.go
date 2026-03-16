package models

import (
	"time"

)

type Members struct {
	ID              int    `json:"id"`
	NAME            string `json:"name"`
	PHONE           string `json:"phone"`
	ADDRESS         string    `json:"address"`
	MEMBERSHIP_DATE time.Time   `json:"membership_date"`
	STATUS          string `json:"status"`
	MEMBERSCOL      string `json:"memberscol"`
}