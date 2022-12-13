package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	//"github.com/slack-go/slack"
	"github.com/shomali11/slacker"
)

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
func main() {
	
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4509054344323-4521856622833-iA5mfzZx07OXFNHw8VtJgejR")
	//os.Setenv("CHANNEL_ID","C04ESGP8B70") //for upload file
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04ESGXHURL-4509175888706-4f6a91541e60dd585f103e12e45f684ded124c74783a384d2fb32ba5f356297c")

	//for upload file
	/*api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"token.txt"}
	for i := 0; i < len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)

	}*/

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))


	  go printCommandEvents(bot.CommandEvents())

	bot.Command("my birthday is <year>", &slacker.CommandDefinition{
		Description: "year calculator",
		//Example:  	"my birthday is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year") //year is in upper "<year>"
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("Your age is %d", age)
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
