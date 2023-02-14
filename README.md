# inventory-barang-GraphQL-Api

## Technologies
Project is created with:
* Go
* PostgreSQL
* [gqlgen](https://gqlgen.com/)

## Query

1. Mendapatkan semua barang masuk
```
query {
  masuks {
    id
    namaBarang
    jumlahBarang
    createdAt
    updatedAt
  }
}
```
- Response
```
{
  "data": {
    "masuks": [
      {
        "id": "546946ba-ab65-11ed-9845-0242ac130002",
        "namaBarang": "cireng",
        "jumlahBarang": 5,
        "createdAt": "2023-02-13T13:12:01.458095+07:00",
        "updatedAt": "2023-02-13T13:12:01.458095+07:00"
      },
      {
        "id": "38379974-ab6f-11ed-8440-0242ac130002",
        "namaBarang": "nasi uduk",
        "jumlahBarang": 55,
        "createdAt": "2023-02-13T14:22:49.122274+07:00",
        "updatedAt": "2023-02-13T14:22:49.122274+07:00"
      }
    ]
  }
}
```
2. Mendapatkan semua barang keluar
```
query {
  keluars {
    id
    barangMasukId
    jumlahKeluar
    barangMasuk {
      namaBarang
      jumlahBarang
    }
  	createdAt
    updatedAt
  }
}
```
- Response
```
{
  "data": {
    "keluars": [
      {
        "id": "f190d50c-ab6f-11ed-8f18-0242ac130002",
        "barangMasukId": "38379974-ab6f-11ed-8440-0242ac130002",
        "jumlahKeluar": 20,
        "barangMasuk": {
          "namaBarang": "nasi uduk",
          "jumlahBarang": 55
        },
        "createdAt": "2023-02-13T14:28:00.085853+07:00",
        "updatedAt": "2023-02-13T14:28:00.085853+07:00"
      }
    ]
  }
}
```

