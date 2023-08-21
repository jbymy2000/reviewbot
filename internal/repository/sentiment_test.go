package repository

import (
	"fmt"
	"testing"
)

func Test_sentiment(t *testing.T) {
	res, _ := sentimenter("Your mother is an awful lady")
	fmt.Println(fmt.Sprintf("0 negtive 1 positive result %d", res.Score))
	res, _ = sentimenter("Your mother is an great lady")
	fmt.Println(fmt.Sprintf("0 negtive 1 positive result %d", res.Score))
	res, _ = sentimenter("Your mother is an lady")
	fmt.Println(fmt.Sprintf("0 negtive 1 positive result %d", res.Score))
}
