package repository

import (
	"errors"
	"fmt"
	"github.com/cdipaolo/sentiment"
	"time"
)

func sentimenter(text string) (res *sentiment.Analysis, err error) {
	if text == "" {
		return nil, errors.New("no text")
	}
	start := time.Now()
	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	duration := time.Since(start) / time.Millisecond
	println(duration)
	return model.SentimentAnalysis(text, sentiment.English), nil // 0
}
