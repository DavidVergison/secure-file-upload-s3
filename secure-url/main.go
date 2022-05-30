package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//TODO: use a base path like "/audio"
	//TODO: use a directory like "/[campaign id]"
	//TODO: check if the directory isn't full (more than X audio file already uploaded)
	filename := request.QueryStringParameters["filename"]

	url, err := getUrl(filename)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       url,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
