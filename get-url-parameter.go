package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"encoding/json"
)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)


	tstring:=fmt.Sprintf("%s",request.QueryStringParameters["bob"])
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       "<html><body>"+tstring+"</body></html>",
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(Handler)
}




