package to

type Send struct {
	OK     bool       `json:"ok"`
	Result SendResult `json:"result"`
}

type SendResult struct {
	/*
		"message_id":63,
		"from":{},
		"chat":{},
		"date":1584115665,
		"text"
	*/
	From      SendFrom `json:"from"`
	Chat      Chat     `json:"chat"`
	MessageID int64    `json:"message_id"`
	Date      int64    `json:"date"`
	Text      string   `json:"text"`
}

type SendFrom struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName  string `json:"user_name"`
}

// Chat represents chat info from message
type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
}
