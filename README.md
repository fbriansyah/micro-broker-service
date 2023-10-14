# Broker Micro Servicer

Aplikasi ini merupakan gateway utama untuk mengakses payment micro service. Service ini menyediakan REST API endpoint untuk setiap micro service. Terdapat beberapa micro service yang berada di belakang gateway, sepert:
- [Session Micro Service](https://github.com/fbriansyah/micro-session-service)
- [Auth Micro Service](https://github.com/fbriansyah/micro-auth-service)
- [Payment Micro Service](https://github.com/fbriansyah/micro-payment-service)
- [Biller Micro Service](https://github.com/fbriansyah/micro-biller-service)

Aplikasi ini menggunkan [gRpc module](https://github.com/fbriansyah/micro-payment-proto) untuk berkomunikasi dengan micro service lainnya.

## Run Application
Untuk menjalankan aplikasi ini di local membutuhkan docker dan docker compose. Kita bisa menggunakan command `docker compose up -d` atau `make up` untuk memulai aplikasi. Untuk menghentikan aplikasi bisa menggunakan `docker compose down` atau `make down`.

# Fitures

## Auth
### Login
Endpoint login digunakan untuk mendapatkan access_token untuk mengakses endpoint lain.

Endpoint: `[POST] /auth/login`

Params:
```json
{
    "username": "test",
    "password": "S3cr3t"
}
```

Response:
```json
{
    "error": false,
    "message": "",
    "data": {
        "session": {
            "id": "<session_id>",
            "user_id": "<user_id>",
            "access_token": "<access_token>",
            "refresh_token": "<refresh_token>",
            "access_token_expires_at": "<date time>",
            "refresh_token_expires_at": "<date time>"
        },
        "user": {
            "id": "<user_id>",
            "username": "<username>",
            "name": "<user full name>"
        }
    }
}
```

## Biller
### Inquiry
Endpoint untuk melihat tagihan dari `bill_number` yang dikirim.

Auth:
- authorization: Bearer <access_token>

Endpoint: `[POST] /biller/inquiry`

Params:
```json
{
    "bill_number": "6310233333331",
    "product_code": "P01"
}
```

Response:
```json
{
    "error": false,
    "message": "",
    "data": {
        "inquiry_id": "<inquiry_id>",
        "name": "<biller_name>",
        "bill_number": "6310233333331",
        "amount": <amount>
    }
}
```

### Payment
Endpoint untuk melakukan pelunasan tagihan.

Auth:
- authorization: Bearer <access_token>

Endpoint: `[POST] /biller/payment`

Params:
```json
{
    "inquiry_id": "<inquiry_id>",
    "amount": <amount>
}
```

Response:
```json
{
    "error": false,
    "message": "",
    "data": {
        "bill_number": "6310233333331",
        "product_code": "P01",
        "name": "Dummy 1",
        "total_amount": 10000,
        "refference_number": "679AFAA67817EF88F6E2DE8F6D05AF771311F49F",
        "transaction_date": "2023-10-11T02:30:50.820057Z"
    }
}
```

## Info
### Balance
Endpoint untuk cek user balance.

Auth:
- authorization: Bearer <access_token>

Endpoint: `[POST] /info/balance`

Response:
```json
{
    "error": false,
    "message": "",
    "data": <int64>
}
```
### Product
Endpoint untuk melihat seluruh product yang terdaftar.

Auth:
- authorization: Bearer <access_token>

Endpoint: `[POST] /info/products`

Response:
```json
{
    "error": false,
    "message": "",
    "data": [
        {
            "product_code": "P01",
            "product_name": "Biller 1"
        }
    ]
}
```

# Note

Aplikasi ini merupakan PoC (proof of concept) atau prototype, jadi masih banyak ruang untuk pengembangan. Saya berharap ada masukan yang bisa membuat project ini menjadi lebih baik lagi.
