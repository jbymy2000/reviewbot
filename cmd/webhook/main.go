package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jbymy2000/reviewbot/config"
	"github.com/jbymy2000/reviewbot/internal/controllers"
	"github.com/jbymy2000/reviewbot/internal/logger"
	"github.com/jbymy2000/reviewbot/internal/repository"
	"time"
)

func main() {
	dbInitCtx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Conf.DbInitTimeout)*time.Second)
	defer cancel()
	//manage postgres conn pool
	_, err := repository.InitDBPool(dbInitCtx)
	defer repository.CloseDBPool()
	if err != nil {
		logger.Error("Data base init fail" + err.Error())
	}
	lambda.Start(controllers.Handler)
}
