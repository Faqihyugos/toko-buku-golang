# Dokumentasi API Toko Buku

## API dapat menyimpan buku
API dapat menyimpan buku melalui rute :
- Method : POST
- URL : /api/v1/book/add
- Body Request :
```json
{
    "title": "Atomic Habbits",
    "description": "buku inspirasi untuk melakukan habit yang positif",
    "year": 2010,
    "pages": 100,
    "language": "Indonesia",
    "publisher": "PT Gramedia",
    "price": 100000,
    "stock" : 50
}
```

Response to be returned :
- Status Code : 201
- Response Body :

```json
{
  "Status": 201,
  "Message": "Book create successfully",
  "Data": {
    "id": 1,
    "title": "Atomic Habbits",
    "description": "buku inspirasi untuk melakukan habit yang positif",
    "year": 2010,
    "pages": 100,
    "language": "Indonesia",
    "publisher": "PT Gramedia",
    "price": 100000,
    "stock": 50,
    "purchase_amount": 0
  }
}
```

## API dapat menampilkan seluruh buku
API dapat menampilkan semua buku yang disimpan melalui rute :
- Method : GET
- URL : /api/v1/book/list

Response to be returned :
- Status Code : 200
- Response Body :

```json
{
    "Status": 200,
    "Message": "All Get Data Success",
    "Data": [
        {
            "id": 1,
            "title": "Bahasa golang",
            "description": "Golang google",
            "year": 2021,
            "pages": 10,
            "language": "Indonesia",
            "publisher": "enigma",
            "price": 50000,
            "stock": 95,
            "purchase_amount": 5
        },
        {
            "id": 2,
            "title": "Dongeng kancil",
            "description": "",
            "year": 2007,
            "pages": 100,
            "language": "Indonesia",
            "publisher": "Bobo",
            "price": 20000,
            "stock": 50,
            "purchase_amount": 0
        }
    ]
}
```

## API dapat mencari berdasarkan Id
API dapat mencari book berdasarkan id buku melalui rute :
- Method : GET
- URL : /api/v1/book/:id

Response to be returned :
- Status Code : 200
- Response Body :
```json
{
  "Status": 200,
  "Message": "OK",
  "Data": {
    "id": 1,
    "title": "Atomic Habbits",
    "description": "buku inspirasi untuk melakukan habit yang positif",
    "year": 2010,
    "pages": 100,
    "language": "Indonesia",
    "publisher": "PT Gramedia",
    "price": 100000,
    "stock": 50,
    "purchase_amount": 0
  }
}
```

## API dapat mengubah data buku
API dapat mengubah book berdasarkan id buku melalui rute :
- Method : PUT
- URL : /api/v1/book/:id
- Body Request :
```json
{
    "title": "Atomic Habbits",
    "description": "buku inspirasi untuk melakukan habit yang positif",
    "year": 2010,
    "pages": 100,
    "language": "Indonesia",
    "publisher": "PT Gramedia",
    "price": 120000,
    "stock" : 50
}
```

Response to be returned :
- Status Code : 200
- Response Body :
```json
{
    "Status": 200,
    "Message": "Book updated successfully",
    "Data": {
        "id": 5,
        "title": "Atomic Habbits",
        "description": "buku inspirasi untuk melakukan habit yang positif",
        "year": 2010,
        "pages": 100,
        "language": "Indonesia",
        "publisher": "PT Gramedia",
        "price": 120000,
        "stock": 50,
        "purchase_amount": 0
    }
}
```

Selanjutnya bisa di test langsung mengggunakan postman donwload dan import dengan namafile : Toko-buku.postman_collection.json

Terimakasih.

