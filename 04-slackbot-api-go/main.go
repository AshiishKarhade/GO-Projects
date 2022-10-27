package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-811591088449-3741496313510-91UXuZTKD7C9LCu08ne11vQw")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03NCMBQYQ1-3750493479844-d79322297369cb030060790284c951e2bd446666275cd7e4a3051382d0986641")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Example:     "my year of birth is 1998",
		Handler: func(botCtcx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age s is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
