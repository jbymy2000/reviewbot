package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jbymy2000/reviewbot/internal/controllers"
	"github.com/jbymy2000/reviewbot/internal/logger"
	"github.com/jbymy2000/reviewbot/internal/repository"
)

func main() {
	//manage postgres conn pool
	_, err := repository.InitDBPool()
	defer repository.CloseDBPool()
	if err != nil {
		logger.Error("Data base init fail" + err.Error())
	}
	lambda.Start(controllers.Handler)
}
