package entity

import "time"

type BarangMasuk struct {
	ID         string
	NamaBarang string
	Jumlah     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
