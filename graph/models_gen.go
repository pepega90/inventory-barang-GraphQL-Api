// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"time"
)

type BarangKeluar struct {
	ID            string       `json:"id"`
	JumlahKeluar  int          `json:"jumlahKeluar"`
	BarangMasukID string       `json:"barangMasukId"`
	BarangMasuk   *BarangMasuk `json:"barangMasuk"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

type BarangKeluarResponse struct {
	BarangKeluar *BarangKeluar `json:"barangKeluar"`
}

type BarangMasuk struct {
	ID           string    `json:"id"`
	NamaBarang   string    `json:"namaBarang"`
	JumlahBarang int       `json:"jumlahBarang"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BarangMasukResponse struct {
	Barang *BarangMasuk `json:"barang"`
}

type CreateBarangKeluarInput struct {
	BarangMasukID string `json:"barangMasukId"`
	JumlahKeluar  int    `json:"jumlahKeluar"`
}

type CreateBarangmasukInput struct {
	NamaBarang   string `json:"namaBarang"`
	JumlahBarang int    `json:"jumlahBarang"`
}

type ResBody struct {
	Message    string `json:"message"`
	StatusCode string `json:"status_code"`
}

type UpdateBarangKeluarInput struct {
	JumlahKeluar  int    `json:"jumlahKeluar"`
	BarangMasukID string `json:"barangMasukId"`
}