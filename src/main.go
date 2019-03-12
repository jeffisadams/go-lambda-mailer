package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	gopher "github.com/jpoehls/gophermail"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := gopher.Message{}

	message.AddTo("Jeff Adams <jeff@tenmilesquare.com>")

	message.SetFrom("Jeff Adams <jeffisadams@gmail.com>")

	body, err := message.Bytes()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
