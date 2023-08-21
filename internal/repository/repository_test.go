package repository

import (
	"context"
	"github.com/jbymy2000/reviewbot/internal/logger"
	"testing"
)

func Test_db(t *testing.T) {
	_, err := InitDBPool(context.TODO())
	if err != nil {
		logger.Error("Data base init fail" + err.Error())
	}
	defer CloseDBPool()
	println(InsertRatings(context.TODO(), 2, 2, 3))
}
