package graph

import (
	"inventory_graphql_api/internal/entity"
)

func mapToBarangMasuk(data *entity.BarangMasuk) *BarangMasuk {
	return &BarangMasuk{
		ID:           data.ID,
		NamaBarang:   data.NamaBarang,
		JumlahBarang: data.Jumlah,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func mapToBarangMasukResponse(input *entity.BarangMasuk) *BarangMasukResponse {
	return &BarangMasukResponse{
		Barang: &BarangMasuk{
			ID:           input.ID,
			NamaBarang:   input.NamaBarang,
			JumlahBarang: input.Jumlah,
			CreatedAt:    input.CreatedAt,
			UpdatedAt:    input.UpdatedAt,
		},
	}
}

func mapToListBarangMasuk(arr []entity.BarangMasuk) []*BarangMasuk {
	var listResponse []*BarangMasuk
	for _, val := range arr {
		listResponse = append(listResponse, &BarangMasuk{
			ID:           val.ID,
			NamaBarang:   val.NamaBarang,
			JumlahBarang: val.Jumlah,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		})
	}

	return listResponse
}

func mapToListBarangKeluar(arr []entity.BarangKeluar) []*BarangKeluar {
	var listResponse []*BarangKeluar
	for _, val := range arr {
		listResponse = append(listResponse, &BarangKeluar{
			ID:            val.ID,
			JumlahKeluar:  val.JumlahKeluar,
			BarangMasukID: val.BarangMasukID,
			CreatedAt:     val.CreatedAt,
			UpdatedAt:     val.UpdatedAt,
		})
	}

	return listResponse
}

func mapToBarangKeluarResponse(input *entity.BarangKeluar) *BarangKeluarResponse {
	return &BarangKeluarResponse{
		BarangKeluar: &BarangKeluar{
			ID:            input.ID,
			JumlahKeluar:  input.JumlahKeluar,
			BarangMasukID: input.BarangMasukID,
			CreatedAt:     input.CreatedAt,
			UpdatedAt:     input.UpdatedAt,
		},
	}
}
