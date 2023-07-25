package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK-BOT-TOKEN", "xoxb-5660323445328-5649346785889-QUlGSJUm8Eusjoa65gq5oGQS")
	os.Setenv("CHANNEL_ID", "C05K37EQGJV")
	api := slack.New(os.Getenv("SLACK-BOT-TOKEN"))
	channelArr := []string{os.Getenv("Channel_ID")}
	fileArr := []string{"GoDev_slack.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s, caused error", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
