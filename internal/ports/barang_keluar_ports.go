package ports

import (
	"context"
	"inventory_graphql_api/internal/entity"
)

type BarangKeluarInput struct {
	JumlahKeluar  int
	BarangMasukID string
}

type IBarangMasukKeluarRepository interface {
	All(ctx context.Context) ([]entity.BarangKeluar, error)
	Create(ctx context.Context, barang BarangKeluarInput) (*entity.BarangKeluar, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, barang BarangKeluarInput) (*entity.BarangKeluar, error)
}
