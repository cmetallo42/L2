package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Configuration struct {
		DSN string
	}

	Database struct {
		Pool *pgxpool.Pool
		DCTX context.Context
	}
)

const schema = `
CREATE TABLE IF NOT EXISTS event (
	event_id SERIAL PRIMARY KEY,
	user_id int4,
	name VARCHAR(32),
	event_date DATE
  );
`

func NewDatabase(c *Configuration) (d *Database, err error) {
	con := context.Background()
	p, err := pgxpool.Connect(con, c.DSN)
	if err != nil {
		return
	}
	_, err = p.Exec(con, schema)
	if err != nil {
		return
	}
	d = &Database{
		Pool: p,
		DCTX: con,
	}

	return
}

func (d *Database) Close() (err error) {
	d.Pool.Close()
	return
}
