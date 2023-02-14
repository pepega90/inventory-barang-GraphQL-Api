package entity

import "time"

type BarangKeluar struct {
	ID            string
	JumlahKeluar  int
	BarangMasukID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
