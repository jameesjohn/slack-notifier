package main

import (
	"github.com/jameesjohn/slackNotifier"
	"log"
)

func main() {
	log.Println("Starting simple example")
	message := slack_notifier.NewMessage().Bold("Hi, My name is James.").
		NewLine().
		Italicize("And I'm checking out this guy")

	notifier := slack_notifier.NewNotifier("https://hooks.slack.com/services/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")

	err := notifier.Send(message)

	if err != nil {
		log.Println(err)
	}

	log.Println("Done with simple example")
}
