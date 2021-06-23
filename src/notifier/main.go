package main

import (

	"context"
	"encoding/json"


    "log"
	"os"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aiomonitors/godiscord"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type DiscordMessage struct {
	username    string          `json:"username"`
	avatar_url  string          `json:"avatar_url"`
	content 	string          `json:"content"`
	embeds     	json.RawMessage `json:"detail"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event events.CloudWatchEvent) {
		fmt.Printf("Detail = %s\n", event.Detail)


	// Authentication Token pulled from environment variable DISCORD_WEBHOOK_URL
	WEBHOOKURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if WEBHOOKURL == "" {
		log.Fatal("WEBHOOK URL not set")
		return
	}

	embed := godiscord.NewEmbed("AWS Event Type", string(event.DetailType), "https://stackoverflow.com/questions/53935198/in-my-discord-webhook-i-am-getting-the-error-embeds-0")
	embed.AddField("Account ID", string(event.AccountID), true)
	embed.AddField("Raw Message Detail", string(event.Detail), true)
	err := embed.SendToWebhook(WEBHOOKURL)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	lambda.Start(Handler)
}
