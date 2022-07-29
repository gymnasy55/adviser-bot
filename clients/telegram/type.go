package telegram

type TgResponse[T any] struct {
	Ok     bool `json:"ok"`
	Result T    `json:"result"`
}

type UpdatesResponse = TgResponse[[]Update]

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}
