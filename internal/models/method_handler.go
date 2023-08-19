package models

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jbymy2000/reviewbot/config"
	"github.com/jbymy2000/reviewbot/internal/logger"
	"github.com/jbymy2000/reviewbot/internal/repository"
	"net/http"
	"strconv"
)

type MethodHandler interface {
	Process(ctx context.Context, req events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error)
}

func MethodFactory(method string) (mh MethodHandler, err error) {
	if method == http.MethodPost {
		return &PostHandler{}, nil
	} else if method == http.MethodGet {
		return &GetHandler{}, nil
	}
	logger.Errorf("method not covered %s", method)
	return nil, errors.New("method not covered" + method)
}

type PostHandler struct{}

func (p *PostHandler) Process(ctx context.Context, req events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	var payload *WebhookPayload
	err = json.Unmarshal([]byte(req.Body), &payload)
	if err != nil {
		logger.Error("un marthal webhook payload error:" + err.Error())
		return newResponse(http.StatusBadRequest, err.Error()), nil
	}
	var text, psid string
	var quickReply *QuickReply
	if payload != nil && !(len(payload.Entry) >= 0 && len(payload.Entry[0].Messaging) > 0) {
		s, _ := json.Marshal(payload)
		logger.Errorf("un able to parse json %s", s)
		return newResponse(http.StatusBadRequest, err.Error()), nil
	}
	text = payload.Entry[0].Messaging[0].Message.Text
	psid = payload.Entry[0].Messaging[0].Sender.ID
	quickReply = payload.Entry[0].Messaging[0].Message.QuickReply
	//thanks letter
	if text == config.Conf.ThanksLetter {
		err = repository.SenderSurvy(context.TODO(), psid, text)
		if err != nil {
			logger.Errorf("send servy fail %s", err.Error())
			return newResponse(http.StatusBadRequest, err.Error()), nil
		}
	} else if quickReply != nil {
		//quick replay postsend message
		rating, _ := strconv.Atoi(quickReply.Payload)
		psidint, _ := strconv.Atoi(psid)
		logger.Infof("%v%", quickReply)
		err := repository.InsertRatings(ctx, 0, int64(rating), int64(psidint))
		if err != nil {
			logger.Errorf("Error:", err.Error())
			return newResponse(http.StatusBadRequest, err.Error()), nil
		}
	} else {
		//normal message
		err = repository.SenderMessage(context.TODO(), psid, text)
		if err != nil {
			return newResponse(http.StatusBadRequest, err.Error()), nil
		}
	}
	return newResponse(http.StatusOK, "Success"), nil
}

type GetHandler struct{}

func (p *GetHandler) Process(ctx context.Context, req events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	mode := req.QueryStringParameters["hub.mode"]
	token := req.QueryStringParameters["hub.verify_token"]
	challenge := req.QueryStringParameters["hub.challenge"]
	if !(mode == "subscribe") || !(token == config.Conf.VerifyToken) {
		return newResponse(http.StatusUnauthorized, challenge), nil

	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       challenge,
	}, nil
}

func newResponse(statusCode int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}
