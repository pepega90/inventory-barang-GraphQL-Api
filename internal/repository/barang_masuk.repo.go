package repository

import (
	"context"
	"inventory_graphql_api/internal/entity"
	"inventory_graphql_api/internal/ports"
	"inventory_graphql_api/postgres"
	"log"
)

type BarangMasukRepo struct {
	DB *postgres.DB
}

func NewBarangMasukRepo(db *postgres.DB) ports.IBarangMasukRepository {
	return &BarangMasukRepo{db}
}

func (br *BarangMasukRepo) All(ctx context.Context) ([]entity.BarangMasuk, error) {
	query := `SELECT * FROM barang_masuks;`

	rows, err := br.DB.Pool.Query(ctx, query)
	if err != nil {
		log.Fatalf("error (repo) get all barang masuk: %v", err)
		return nil, err
	}

	var listBarangMasuk []entity.BarangMasuk
	for rows.Next() {
		barang := entity.BarangMasuk{}
		rows.Scan(&barang.ID, &barang.NamaBarang, &barang.Jumlah, &barang.CreatedAt, &barang.UpdatedAt)
		listBarangMasuk = append(listBarangMasuk, barang)
	}

	return listBarangMasuk, nil
}

func (br *BarangMasukRepo) Create(ctx context.Context, barang ports.BarangInput) (*entity.BarangMasuk, error) {
	log.Printf("create: %+v", barang)
	query := `INSERT INTO barang_masuks (nama_barang, jumlah) VALUES ($1, $2) RETURNING *;`

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (repo) create barang: begin: %v", err)
		return nil, err
	}

	row := tx.QueryRow(ctx, query, barang.NamaBarang, barang.Jumlah)

	createdBarang := entity.BarangMasuk{}

	if err := row.Scan(&createdBarang.ID, &createdBarang.NamaBarang, &createdBarang.Jumlah, &createdBarang.CreatedAt, &createdBarang.UpdatedAt); err != nil {
		log.Fatalf("error (repo) insert new barang: %v", err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (repo) create: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (repo) create: roolback transaction: %v", err)
			return nil, err
		}
		return nil, err
	}

	return &createdBarang, nil
}

func (br *BarangMasukRepo) FindByID(ctx context.Context, id string) (*entity.BarangMasuk, error) {
	query := `SELECT * FROM barang_masuks WHERE id = $1;`

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (repo) findbyid barang: begin: %v", err)
		return nil, err
	}

	var findedBarang entity.BarangMasuk

	row := tx.QueryRow(ctx, query, id)

	err = row.Scan(&findedBarang.ID, &findedBarang.NamaBarang, &findedBarang.Jumlah, &findedBarang.CreatedAt, &findedBarang.UpdatedAt)
	if err != nil {
		log.Fatalf("error (repo) gagal find barang by id: %v", err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (repo) findbyid: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (repo) findbyid: roolback transaction: %v", err)
			return nil, err
		}
		return nil, err
	}

	return &findedBarang, nil
}

func (br *BarangMasukRepo) Update(ctx context.Context, id string, barang ports.BarangInput) (*entity.BarangMasuk, error) {
	query := `UPDATE barang_masuks SET nama_barang=$1, jumlah=$2 WHERE id=$3 RETURNING *;`

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (repo) update barang: begin: %v", err)
		return nil, err
	}

	var updatedBarang entity.BarangMasuk

	row := tx.QueryRow(ctx, query, barang.NamaBarang, barang.Jumlah, id)

	err = row.Scan(&updatedBarang.ID, &updatedBarang.NamaBarang, &updatedBarang.Jumlah, &updatedBarang.CreatedAt, &updatedBarang.UpdatedAt)
	if err != nil {
		log.Fatalf("error (repo) error scan updated barang: %v", err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (repo) update: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (repo) update: roolback transaction: %v", err)
			return nil, err
		}
		return nil, err
	}

	return &updatedBarang, nil
}

func (br *BarangMasukRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM barang_masuks WHERE id = $1;`

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (repo) delete: begin: %v", err)
		return err
	}

	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		log.Fatalf("error (repo) delete: gagal menghapus barang: %v", err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (repo) delete: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (repo) delete: roolback transaction: %v", err)
			return err
		}
		return err
	}

	return nil
}
