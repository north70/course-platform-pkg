package pg

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/north70/course-platform-pkg/db"
)

type Client struct {
	master db.DB
}

func NewPgClient(master *pgxpool.Pool) *Client {
	return &Client{master: NewDB(master)}
}

func (c *Client) DB() db.DB {
	return c.master
}

func (c *Client) Close() error {
	if c.master != nil {
		c.master.Close()
	}

	return nil
}
