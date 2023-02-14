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
## Mutation

1. Membuat barang masuk
```
mutation {
  createBarangMasuk(input: {namaBarang: "nasi uduk", jumlahBarang: 100}) {
    barang {
      id
      namaBarang
      jumlahBarang
      createdAt
      updatedAt
    }
  }
}
```
- Response
```
{
  "data": {
    "createBarangMasuk": {
      "barang": {
        "id": "a690bd70-ac0e-11ed-8703-0242ac120002",
        "namaBarang": "nasi uduk",
        "jumlahBarang": 100,
        "createdAt": "2023-02-14T09:24:04.226925+07:00",
        "updatedAt": "2023-02-14T09:24:04.226925+07:00"
      }
    }
  }
}
```
2. Hapus barang masuk berdasarkan id
```
mutation {
  hapusBarangMasuk(id: "4a6e8a9e-ab65-11ed-9845-0242ac130002") {
    message
    status_code
  }
}
```
- Response
```
{
  "data": {
    "hapusBarangMasuk": {
      "message": "success",
      "status_code": "200"
    }
  }
}
```
3. Update barang masuk
```
mutation {
  updateBarangMasuk(
    id: "4a6e8a9e-ab65-11ed-9845-0242ac130002"
    input: {namaBarang: "cireng", jumlahBarang: 25}
  ) {
    barang {
      id
      namaBarang
      jumlahBarang
      createdAt
      updatedAt
    }
  }
}
```
- Response
```
{
  "data": {
    "updateBarangMasuk": {
      "barang": {
        "id": "546946ba-ab65-11ed-9845-0242ac130002",
        "namaBarang": "ayam geprek",
        "jumlahBarang": 25,
        "createdAt": "2023-02-13T13:12:01.458095+07:00",
        "updatedAt": "2023-02-13T13:12:01.458095+07:00"
      }
    }
  }
}
```
4. Mendapatkan barang masuk by id
```
mutation {
  getBarangMasukByID(id:"546946ba-ab65-11ed-9845-0242ac130002") {
    barang {
      id
      namaBarang
      jumlahBarang
    }
  }
}
```
- Response
```
{
  "data": {
    "getBarangMasukByID": {
      "barang": {
        "id": "546946ba-ab65-11ed-9845-0242ac130002",
        "namaBarang": "ayam geprek",
        "jumlahBarang": 25
      }
    }
  }
}
```
5. Membuat barang keluar
```
mutation {
  createBarangKeluar(
    input: {barangMasukId: "38379974-ab6f-11ed-8440-0242ac130002", jumlahKeluar: 5}
  ) {
    barangKeluar {
      id
      barangMasukId
      barangMasuk {
        namaBarang
        jumlahBarang
      }
      jumlahKeluar
    }
  }
}
```
- Response
```
{
  "data": {
    "createBarangKeluar": {
      "barangKeluar": {
        "id": "35666f40-ac0f-11ed-9238-0242ac120002",
        "barangMasukId": "38379974-ab6f-11ed-8440-0242ac130002",
        "barangMasuk": {
          "namaBarang": "nasi uduk",
          "jumlahBarang": 50
        },
        "jumlahKeluar": 5
      }
    }
  }
}
```
6. Hapus barang keluar
```
mutation {
  hapusBarangKeluar(id:"35666f40-ac0f-11ed-9238-0242ac120002") {
    message
    status_code
  }
}
```
- Response
```
{
  "data": {
    "hapusBarangKeluar": {
      "message": "success",
      "status_code": "200"
    }
  }
}
```
7. Update barang keluar
```
mutation {
  updateBarangKeluar(
    id: "f190d50c-ab6f-11ed-8f18-0242ac130002"
    input: {jumlahKeluar: 20, barangMasukId: "38379974-ab6f-11ed-8440-0242ac130002"}
  ) {
    barangKeluar {
      id
      jumlahKeluar
      barangMasuk {
        namaBarang
        jumlahBarang
      }
      createdAt
      updatedAt
    }
  }
}
```
- Response
```
{
  "data": {
    "updateBarangKeluar": {
      "barangKeluar": {
        "id": "f190d50c-ab6f-11ed-8f18-0242ac130002",
        "jumlahKeluar": 20,
        "barangMasuk": {
          "namaBarang": "nasi uduk",
          "jumlahBarang": 50
        },
        "createdAt": "2023-02-13T14:28:00.085853+07:00",
        "updatedAt": "2023-02-13T14:28:00.085853+07:00"
      }
    }
  }
}
```
