package repository

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jbymy2000/reviewbot/config"
	"github.com/jbymy2000/reviewbot/internal/logger"
	"io"
	"net/http"
)

func SenderMessage(ctx context.Context, psID string, message string) (err error) {
	url := fmt.Sprintf(config.Conf.SenderMsgUrl, config.Conf.PageId, config.Conf.VerifyToken)
	data := fmt.Sprintf(`{
	  "recipient": {
		"id": "%s"
	  },
	  "messaging_type": "RESPONSE",
	  "message": {
		"text": "%s"
	  }
	}`, psID, message)
	logger.Logger.Println(url)
	logger.Logger.Println(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return nil
}

func SenderSurvy(ctx context.Context, psID string, message string) (err error) {
	url := fmt.Sprintf(config.Conf.SenderQuickUrl, config.Conf.VerifyToken)
	method := "POST"
	data := fmt.Sprintf(`{
		  "recipient":{
			"id":"%s"
		  },
		  "messaging_type": "RESPONSE",
		  "message":{
			"text": "rating:",
			"quick_replies":[
			  {
				"content_type":"text",
				"title":"1",
				"payload":"1",
				"image_url":""
			  },{
				"content_type":"text",
				"title":"2",
				"payload":"2",
				"image_url":""
			  },{
				"content_type":"text",
				"title":"3",
				"payload":"3",
				"image_url":""
			  },{
				"content_type":"text",
				"title":"4",
				"payload":"4",
				"image_url":""
			  },{
				"content_type":"text",
				"title":"5",
				"payload":"5",
				"image_url":""
			  }
			]
		  }
	}`, psID)

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return nil
}
