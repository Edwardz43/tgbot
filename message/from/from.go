package from

// Update is the Update message from specific channel
type Update struct {
	OK         bool     `json:"ok"`
	ResultList []Result `json:"result"`
}

// Result is the result list in update message
type Result struct {
	UpdateID int64   `json:"update_id"`
	Message  Message `json:"message"`
}

// Message represents a message data from update
type Message struct {
	MessageID int64  `json:"message_id"`
	From      From   `json:"from"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
}

// From represents message source info
type From struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"user_name"`
	LanguageCode string `json:"language_code"`
}

// Chat represents chat info from message
type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
}
