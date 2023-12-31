package controllers

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jbymy2000/reviewbot/config"
	"github.com/jbymy2000/reviewbot/internal/logger"
	"github.com/jbymy2000/reviewbot/internal/models"
	"net/http"
	"time"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	//time out control
	globalCtx, cancel1 := context.WithTimeout(context.Background(), time.Duration(config.Conf.GlobalTimeout)*time.Second)
	defer cancel1()
	logger.Infof("my app got request %v", req)
	mh, err := models.MethodFactory(req.HTTPMethod)
	if err != nil {
		return newResponse(http.StatusBadRequest, "get method handler err "+err.Error()), err
	}
	res, err = mh.Process(globalCtx, req)
	if err != nil {
		return newResponse(http.StatusBadRequest, "process "+err.Error()), err
	}

	return res, nil
}

func newResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}
