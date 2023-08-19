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
	//time out control
	globalCtx, cancel1 := context.WithTimeout(context.Background(), time.Duration(config.Conf.GlobalTimeout)*time.Second)
	defer cancel1()
	dbInitCtx, cancel2 := context.WithTimeout(context.Background(), time.Duration(config.Conf.DbInitTimeout)*time.Second)
	defer cancel2()
	//manage postgres conn pool
	_, err := repository.InitDBPool(dbInitCtx)
	defer repository.CloseDBPool()
	if err != nil {
		logger.Error("Data base init fail" + err.Error())
	}
	lambda.StartWithOptions(controllers.Handler, lambda.WithContext(globalCtx))
}
