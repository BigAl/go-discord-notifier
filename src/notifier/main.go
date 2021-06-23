package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}
    
//	message := map[string]interface{}{
//		"hello": "world",
//		"life":  42,
//		"embedded": map[string]string{
//			"yes": "of course!",
//		},
//	}

//	bytesRepresentation, err := json.Marshal(message)
//	if err != nil {
//		log.Fatalln(err)
//	}

//	discordresp, err := http.Post("***REMOVED***4", "application/json", bytes.NewBuffer(bytesRepresentation))
//	if err != nil {
//		log.Fatalln(err)
//	}

//	var result map[string]interface{}

//	json.NewDecoder(discordresp.Body).Decode(&result)

//	log.Println(result)
//	log.Println(result["data"])
    discordpost()

	return resp, nil
}

func discordpost() {

	// Authentication Token pulled from environment variable DISCORD_WEBHOOK_URL
	WEBHOOKURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if WEBHOOKURL == "" {
		return
	}

//Encode the data
postBody, _ := json.Marshal(map[string]string{
	"content":  "Alan Test hook via variable",
 })

 
 responseBody := bytes.NewBuffer(postBody)
//Leverage Go's HTTP Post function to make request
 resp, err := http.Post(WEBHOOKURL, "application/json", responseBody)
//Handle Error
 if err != nil {
	log.Fatalf("An Error Occured %v", err)
 }
 defer resp.Body.Close()
//Read the response body
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil {
	log.Fatalln(err)
 }
 sb := string(body)
 log.Printf(sb)

}

func main() {
	lambda.Start(Handler)
}
