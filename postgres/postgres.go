package postgres

import (
	"context"
	"inventory_graphql_api/config"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
	conf *config.Config
}

func New(ctx context.Context, conf *config.Config) *DB {
	dbConf, err := pgxpool.ParseConfig(conf.DatabaseURL)
	if err != nil {
		log.Fatalf("cant pase postgres config: %v", err)
		return nil
	}

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("error connecting to potgres: %v", err)
		return nil
	}

	// pool.Exec(context.TODO(), `
	// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	// CREATE TABLE IF NOT EXISTS barang_masuks (
	//     id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
	// 	nama_barang VARCHAR(255) NOT NULL,
	//     jumlah INT NOT NULL,
	//     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	// );`)

	// pool.Exec(context.TODO(), `
	// 	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	// CREATE TABLE IF NOT EXISTS barang_keluars (
	//     id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
	//     jumlah_keluar INT NOT NULL,
	//     barang_masuk_id UUID NOT NULL REFERENCES barang_masuks (id),
	//     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	// );`)

	db := &DB{Pool: pool, conf: conf}

	db.Ping(ctx)

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("can't ping postgres: %v", err)
	}

	log.Println("postgres connected")
}
