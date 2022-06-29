package core

type Message struct {
	Message string  `json:"message"`
	Tag     *string `json:"tag,omitempty"`
}
