package graph

import (
	"context"
	"fmt"
	"inventory_graphql_api/internal/ports"
	"log"
	"net/http"
)

func (q *queryResolver) Masuks(ctx context.Context) ([]*BarangMasuk, error) {
	res, err := q.BarangMasukRepo.All(ctx)
	if err != nil {
		log.Fatalf("error (graph) get all barang masuk: %v", err)
		return nil, err
	}

	return mapToListBarangMasuk(res), nil
}

func (m *mutationResolver) GetBarangMasukByID(ctx context.Context, id string) (*BarangMasukResponse, error) {
	res, err := m.BarangMasukRepo.FindByID(ctx, id)
	if err != nil {
		log.Fatalf("error (graph) find barang masuk by id: %v", err)
		return nil, err
	}

	return mapToBarangMasukResponse(res), nil
}

func (m *mutationResolver) CreateBarangMasuk(ctx context.Context, input CreateBarangmasukInput) (*BarangMasukResponse, error) {
	res, err := m.BarangMasukRepo.Create(ctx, ports.BarangInput{
		NamaBarang: input.NamaBarang,
		Jumlah:     input.JumlahBarang,
	})
	if err != nil {
		log.Fatalf("error (graph) create barang masuk: %v", err)
		return nil, err
	}

	return mapToBarangMasukResponse(res), nil
}

func (m *mutationResolver) UpdateBarangMasuk(ctx context.Context, id string, input CreateBarangmasukInput) (*BarangMasukResponse, error) {
	res, err := m.BarangMasukRepo.Update(ctx, id, ports.BarangInput{
		NamaBarang: input.NamaBarang,
		Jumlah:     input.JumlahBarang,
	})
	if err != nil {
		log.Fatalf("error (graph) update barang masuk: %v", err)
		return nil, err
	}

	return mapToBarangMasukResponse(res), nil
}

func (m *mutationResolver) HapusBarangMasuk(ctx context.Context, id string) (*ResBody, error) {
	err := m.BarangMasukRepo.Delete(ctx, id)
	if err != nil {
		log.Fatalf("error (graph) hapus barang masuk: %v", err)
		return nil, err
	}
	return &ResBody{
		Message:    "success",
		StatusCode: fmt.Sprintf("%d", http.StatusOK),
	}, nil
}
