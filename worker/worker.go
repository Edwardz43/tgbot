package worker

import (
	"Edwardz43/tgbot/message/from"
	"time"
)

type Worker interface {
	Do(func(args ...interface{}) error)
}

type Job struct {
	ID          int64        `json:"id"`
	DeliverDate *time.Time   `json:"deliver_date"`
	FinishDate  *time.Time   `json:"finish_date"`
	Done        bool         `json:"done"`
	Result      *from.Result `json:"result"`
}
