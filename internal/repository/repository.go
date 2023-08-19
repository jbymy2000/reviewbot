package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jbymy2000/reviewbot/config"
	_ "github.com/lib/pq"
)

var pool *pgxpool.Pool

func InitDBPool() (db *pgxpool.Pool, err error) {
	// Replace with your PostgreSQL connection information
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.Conf.MySQL.User,
		config.Conf.MySQL.Password,
		config.Conf.MySQL.IP,
		config.Conf.MySQL.Port,
		config.Conf.MySQL.Database,
	)
	// Open a database connection
	pool, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, errors.New("cann't get db connection " + err.Error())
	}
	return pool, nil
}

func CloseDBPool() {
	if pool != nil {
		pool.Close()
	}
}

func InsertRatings(ctx context.Context, userID, rating, psid int64) (err error) {
	query := `
        INSERT INTO public.ratings (user_id, rating, psid)
        VALUES ($1, $2, $3)
        RETURNING id;`

	var insertedID int
	err = pool.QueryRow(context.Background(), query, userID, rating, psid).Scan(&insertedID)
	if err != nil {
		return err
	}

	return nil
}
