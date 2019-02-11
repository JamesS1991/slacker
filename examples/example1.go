package main

import (
	"context"
	"github.com/shomali11/slacker"
	"log"
)

func main() {
	bot := slacker.NewClient("xoxp-548463018663-547303896739-546679911856-81abd859899658d32760879779217461")

	definition := &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	}

	bot.Command("ping", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
