package repository

import (
	"context"
	"inventory_graphql_api/internal/entity"
	"inventory_graphql_api/internal/ports"
	"inventory_graphql_api/postgres"
	"log"
)

type BarangKeluarRepo struct {
	DB *postgres.DB
}

func NewBarangKeluarRepo(db *postgres.DB) ports.IBarangMasukKeluarRepository {
	return &BarangKeluarRepo{db}
}

func (br *BarangKeluarRepo) All(ctx context.Context) ([]entity.BarangKeluar, error) {
	query := `SELECT * FROM barang_keluars;`

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) all: begin: %v", err)
		return nil, err
	}

	rows, err := tx.Query(ctx, query)
	if err != nil {
		log.Fatalf("error (barang keluar repo) get all barang keluar: %v", err)
		return nil, err
	}

	var listKeluar []entity.BarangKeluar
	for rows.Next() {
		barangKeluar := entity.BarangKeluar{}
		rows.Scan(&barangKeluar.ID, &barangKeluar.JumlahKeluar, &barangKeluar.BarangMasukID, &barangKeluar.CreatedAt, &barangKeluar.UpdatedAt)
		listKeluar = append(listKeluar, barangKeluar)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) all: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (barang keluar repo) all: roolback transaction: %v", err)
			return nil, err
		}
		return nil, err
	}

	return listKeluar, nil
}

func (br *BarangKeluarRepo) Create(ctx context.Context, barang ports.BarangKeluarInput) (*entity.BarangKeluar, error) {
	query := `
	INSERT INTO barang_keluars 
	(jumlah_keluar, barang_masuk_id) 
	VALUES ($1, $2) 
	RETURNING *;`

	updateJumlahBarangMasukQuery := `UPDATE barang_masuks SET jumlah = jumlah - $1  WHERE id = $2;`
	br.DB.Pool.Exec(ctx, updateJumlahBarangMasukQuery, barang.JumlahKeluar, barang.BarangMasukID)

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) create: begin: %v", err)
		return nil, err
	}

	var createdBarangKeluar entity.BarangKeluar

	row := tx.QueryRow(ctx, query, barang.JumlahKeluar, barang.BarangMasukID)

	err = row.Scan(&createdBarangKeluar.ID, &createdBarangKeluar.JumlahKeluar, &createdBarangKeluar.BarangMasukID, &createdBarangKeluar.CreatedAt, &createdBarangKeluar.UpdatedAt)

	if err != nil {
		log.Fatalf("error (barang keluar repo) create: row scan: %v", err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) all: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (barang keluar repo) all: roolback transaction: %v", err)
			return nil, err
		}
		return nil, err
	}

	return &createdBarangKeluar, nil
}

func (br *BarangKeluarRepo) Update(ctx context.Context, id string, barang ports.BarangKeluarInput) (*entity.BarangKeluar, error) {
	query := `UPDATE barang_keluars SET jumlah_keluar = $1 WHERE id = $2 RETURNING *;`

	// barangMasukRow := br.DB.Pool.QueryRow(ctx, `SELECT jumlah FROM barang_masuks WHERE id = $1;`, barang.BarangMasukID)
	// barangKeluarRow := br.DB.Pool.QueryRow(ctx, `SELECT jumlah_keluar FROM barang_keluars WHERE id = $1;`, id)

	// var jumlahMasuk int
	// var jumlahKeluar int
	// barangMasukRow.Scan(&jumlahMasuk)
	// barangKeluarRow.Scan(&jumlahKeluar)
	// queryUpdateBarangMasuk := `UPDATE barang_masuks SET jumlah = $1 + ($2 - $3) WHERE id = $4;`
	// jumlah = jumlah + (jumlah_keluar - update_jumlah_keluar)
	// 15 = 15 + (5 - 2)
	// 18
	// br.DB.Pool.Exec(ctx, queryUpdateBarangMasuk, jumlahMasuk, jumlahKeluar, barang.JumlahKeluar, barang.BarangMasukID)

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) update: begin: %v", err)
		return nil, err
	}

	var updatedKeluar entity.BarangKeluar
	row := tx.QueryRow(ctx, query, barang.JumlahKeluar, id)

	row.Scan(&updatedKeluar.ID, &updatedKeluar.JumlahKeluar, &updatedKeluar.BarangMasukID, &updatedKeluar.CreatedAt, &updatedKeluar.UpdatedAt)

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) update: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (barang keluar repo) delete: roolback transaction: %v", err)
			return nil, err
		}
		return nil, err
	}

	return &updatedKeluar, nil
}

func (br *BarangKeluarRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM barang_keluars WHERE id = $1;`

	tx, err := br.DB.Pool.Begin(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) delete: begin: %v", err)
		return err
	}

	_, err = tx.Exec(ctx, query, id)

	if err != nil {
		log.Fatalf("error (barang keluar repo) delete: gagal menghapus barang: %v", err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("error (barang keluar repo) delete: commit: %v", err)
		err = tx.Rollback(ctx)
		if err != nil {
			log.Fatalf("error (barang keluar repo) delete: roolback transaction: %v", err)
			return err
		}
		return err
	}

	return nil
}
