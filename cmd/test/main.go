package main

import (
	"context"
	"log"

	"github.com/rleszilm/genms/sql"
	postgres_sql "github.com/rleszilm/genms/sql/postgres"
	"github.com/rleszilm/genms/sql/sqlx"
)

func main() {
	cfg := &postgres_sql.Config{
		EnvConfig: sql.EnvConfig{
			User:     "gameday",
			Password: "gameday",
			Host:     "localhost",
			Port:     5432,
			Database: "gameday",
		},
	}

	db := sqlx.NewDB(cfg)

	ctx := context.Background()
	if err := db.Initialize(ctx); err != nil {
		log.Fatal(err)
	}
	defer db.Shutdown(ctx)

	rows, err := db.Query(ctx, "select pos from pos_test;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj := Wrapper{}
		if err := rows.StructScan(&obj); err != nil {
			log.Fatal(err)
		}
	}
}

type Wrapper struct {
	Geo postgres_sql.NullPoint `db:"pos"`
}
