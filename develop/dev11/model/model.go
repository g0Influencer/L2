package model

import (
	"fmt"
	"strings"
	"time"
)

//формат даты
const Layout = "2006-01-02"

// структура с событиями
type Event struct {
	EventID int    `json:"event_id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Description   string `json:"description"`
	Date    Date   `json:"date"`
}
type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "" || string(b) == `""` {
		*t = Date{time.Now()}
		return nil
	}

	timeStr := strings.ReplaceAll(string(b), `"`, "")
	parsedTime, err := time.Parse("2006-01-02T15:04", timeStr)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02T15:04:00Z",timeStr)
		if err != nil {
			parsedTime, err = time.Parse("2006-01-02",timeStr)
			if err != nil {
				return fmt.Errorf("date format: e.g. 2022-05-10T14:10 error: %v", err)
			}
		}
	}
	*t = Date{parsedTime}
	return nil
}
