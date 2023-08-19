package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Conf Config

func init() {
	lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT")
	var configFilePath string
	if lambdaTaskRoot != "" {
		configFilePath = os.Getenv("LAMBDA_TASK_ROOT") + "/config.json"
	} else {
		configFilePath = "/Users/wudirex/gopath/reviewbot/config/config.json"
	}
	// 读取配置文件
	configFile, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 解析 JSON 数据

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&Conf)
	if err != nil {
		log.Fatalf("Error:", err)
		return
	}
}

type Config struct {
	ThanksLetter   string `json:"thanks_letter_content"`
	SenderQuickUrl string `json:"sender_quick_url"`
	SenderMsgUrl   string `json:"sender_msg_url"`
	PageId         string `json:"page_id"`
	VerifyToken    string `json:"verify_token"`
	MySQL          struct {
		IP       string `json:"ip"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"postgres"`
}
