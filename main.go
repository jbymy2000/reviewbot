package reviewbot

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

type MyEvent struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"how old are you"`
}
type MyResponse struct {
	Message string `json:"Answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}
