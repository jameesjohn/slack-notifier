package slack_notifier

import "fmt"

type Message struct {
	text string
}

func NewMessage() *Message {
	return &Message{text: ""}
}

func (m *Message) SetMessage(text string) *Message {
	m.text = text

	return m
}

func (m *Message) NewLine() *Message {
	m.text += "\n"

	return m
}

func (m *Message) Bold(text string) *Message {
	m.text += fmt.Sprintf("*%s*", text)

	return m
}

func (m *Message) Italicize(text string) *Message {
	m.text += fmt.Sprintf("_%s_", text)

	return m
}

func (m *Message) Link(text, href string) *Message {
	m.text += fmt.Sprintf("<%s|%s>", href, text)

	return m
}

func (m *Message) GetText() string {
	return m.text
}
