package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aiomonitors/godiscord"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.CloudWatchEvent) {
	fmt.Printf("Detail = %s\n", event.Detail)

	// Authentication Token pulled from environment variable DISCORD_WEBHOOK_URL
	WEBHOOKURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if WEBHOOKURL == "" {
		log.Fatal("WEBHOOK URL not set")
		return
	}

	embed := godiscord.NewEmbed(string(event.DetailType), strings.Join(event.Resources, " "), "https://stackoverflow.com/questions/53935198/in-my-discord-webhook-i-am-getting-the-error-embeds-0")
	embed.AddField("Account ID", string(event.AccountID), true)

	//Creating the maps for JSON
	var m map[string]interface{}
	err := json.Unmarshal(event.Detail, &m)
	if err != nil {
		os.Exit(1)
	}

	// Start parseing
	for key, value := range m {
		switch vv := value.(type) {
		case string:
			//fmt.Println(key, ":", vv)
			embed.AddField(key, vv, false)
		case float64:
			//fmt.Println(key, ":", vv)
			embed.AddField(key, strconv.FormatFloat(vv, 'f', 1, 64), false)
		case []interface{}:
			//fmt.Println(key, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			//fmt.Println(key, ":")
			embed.AddField(key, ":", true)
			for i, u := range vv {
				//fmt.Println(i, ":", u)
				embed.AddField(i, u.(string), true)
			}
		default:
			fmt.Println(key, "is of a type I don't know how to handle")
		}
	}

	// Add raw message
	embed.AddField("Raw Message", string(event.Detail), false)

	// Send the embed
	err = embed.SendToWebhook(WEBHOOKURL)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lambda.Start(Handler)
}
