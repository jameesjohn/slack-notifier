package slack_notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Notifier struct {
	webhookUrl string
}

func NewNotifier(webhookURL string) Notifier {
	return Notifier{
		webhookUrl: webhookURL,
	}
}

func (n Notifier) Send(message *Message) error {
	payload := map[string]string{
		"text": message.GetText(),
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	response, err := http.Post(n.webhookUrl, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Message send failed"))
	}

	return nil
}
