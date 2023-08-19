package models

type MessagingEvent struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
}

type Sender struct {
	ID string `json:"id"`
}

type Recipient struct {
	ID string `json:"id"`
}
type QuickReply struct {
	Payload string `json:"payload"`
}

type Message struct {
	MID        string      `json:"mid"`
	Text       string      `json:"text"`
	QuickReply *QuickReply `json:"quick_reply"`
}

type Entry struct {
	ID        string           `json:"id"`
	Time      int64            `json:"time"`
	Messaging []MessagingEvent `json:"messaging"`
}

type WebhookPayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}
