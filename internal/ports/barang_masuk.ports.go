package ports

import (
	"context"
	"inventory_graphql_api/internal/entity"
)

type BarangInput struct {
	NamaBarang string
	Jumlah     int
}

type IBarangMasukRepository interface {
	All(ctx context.Context) ([]entity.BarangMasuk, error)
	FindByID(ctx context.Context, id string) (*entity.BarangMasuk, error)
	Create(ctx context.Context, barang BarangInput) (*entity.BarangMasuk, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, barang BarangInput) (*entity.BarangMasuk, error)
}
