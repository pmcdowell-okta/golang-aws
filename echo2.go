package main


import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	datas := make(map[string]string)

	datas["Dude where is my on echo2"]=request.Body

	returnString, err := json.Marshal(datas)
	_=returnString
	_=err

	//return events.APIGatewayProxyResponse{Body: string("<html><title>wowwowwow</title></html>"), StatusCode: 200}, nil
	//return events.APIGatewayProxyResponse{ StatusCode: 200,
	//  IsBase64Encoded: false,Headers: map[string]string{
	//		"Content-Type": "html/text",
	//	},
	//	Body: "<html><title>wowwowwow</title></html>",
	//	}, nil
	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       "<html><body><h1>First Post!</h1></body></html>",
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(Handler)
}
