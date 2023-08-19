package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_sender(t *testing.T) {
	Payload := map[string]interface{}{
		"recipient": map[string]string{
			"id": "6190237481105605", // Replace with the recipient's Facebook ID
		},
		"messaging_type": "RESPONSE",
		"message": map[string]interface{}{
			"text": "Rating:",
			"quick_replies": []map[string]interface{}{
				{
					"content_type": "text",
					"title":        "1",
					"payload":      "1",
				},
				{
					"content_type": "text",
					"title":        "2",
					"payload":      "2",
				},
				{
					"content_type": "text",
					"title":        "3",
					"payload":      "3",
				},
				{
					"content_type": "text",
					"title":        "4",
					"payload":      "4",
				},
				{
					"content_type": "text",
					"title":        "5",
					"payload":      "5",
				},
			},
		},
	}
	println(Payload)
}

func Test_1(t *testing.T) {
	url := "https://graph.facebook.com/v17.0/me/messages?access_token=EAADr4dWYN1kBO5Tqh6tZB6LmtZC99aqObtltsiQyqHP51UpZCeEUVZCKQ092QH31EMCGGhKGI168Im35oyKWhgQ40JLQ7viyAZBObgPiBRu4z7lTrMLC6dHUG40gIaxHGWZB2zLAMyZCm7eW8rdEbEzZAInBnZC8vnd38S2ahIabZBpddZCapc22JeRbsTv6OueMS0NENwZD"
	method := "POST"

	payload := strings.NewReader(`{
  "recipient":{
    "id":"6190237481105605"
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
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
