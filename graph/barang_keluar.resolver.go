package graph

import (
	"context"
	"fmt"
	"inventory_graphql_api/internal/ports"
	"log"
	"net/http"
)

func (q *queryResolver) Keluars(ctx context.Context) ([]*BarangKeluar, error) {
	res, err := q.BarangKeluarRepo.All(ctx)
	if err != nil {
		log.Fatalf("error (graph) barang keluar resolver: keluars: %v", err)
		return nil, err
	}

	return mapToListBarangKeluar(res), nil
}

func (m *mutationResolver) CreateBarangKeluar(ctx context.Context, input CreateBarangKeluarInput) (*BarangKeluarResponse, error) {
	res, err := m.BarangKeluarRepo.Create(ctx, ports.BarangKeluarInput{
		JumlahKeluar:  input.JumlahKeluar,
		BarangMasukID: input.BarangMasukID,
	})

	if err != nil {
		log.Fatalf("error (graph) barang keluar resolver: create barang keluar: %v", err)
		return nil, err
	}

	return mapToBarangKeluarResponse(res), nil

}

func (m *mutationResolver) UpdateBarangKeluar(ctx context.Context, id string, input UpdateBarangKeluarInput) (*BarangKeluarResponse, error) {
	res, err := m.BarangKeluarRepo.Update(ctx, id, ports.BarangKeluarInput{
		JumlahKeluar:  input.JumlahKeluar,
		BarangMasukID: input.BarangMasukID,
	})
	if err != nil {
		log.Fatalf("error (graph) update barang keluar: %v", err)
		return nil, buildError(ctx, err, http.StatusInternalServerError)
	}

	return mapToBarangKeluarResponse(res), nil

}

func (m *mutationResolver) HapusBarangKeluar(ctx context.Context, id string) (*ResBody, error) {
	err := m.BarangKeluarRepo.Delete(ctx, id)
	if err != nil {
		log.Fatalf("error (graph) gagal hapus barang keluar: %v", err)
		return nil, buildError(ctx, err, http.StatusInternalServerError)
	}

	return &ResBody{
		Message:    "success",
		StatusCode: fmt.Sprintf("%d", http.StatusOK),
	}, nil
}

func (b *barangResolver) BarangMasuk(ctx context.Context, obj *BarangKeluar) (*BarangMasuk, error) {
	barangMasuk, err := b.BarangMasukRepo.FindByID(ctx, obj.BarangMasukID)
	if err != nil {
		log.Fatalf("error (graph) barang masuk: %v", err)
		return nil, err
	}

	return mapToBarangMasuk(barangMasuk), nil
}
