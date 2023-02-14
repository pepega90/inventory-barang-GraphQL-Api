# inventory-barang-GraphQL-Api

## Technologies
Project is created with:
* Go
* PostgreSQL
* [gqlgen](https://gqlgen.com/)

## Query

- Mendapatkan semua barang masuk
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
