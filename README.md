# Broker Micro Servicer

Aplikasi ini merupakan gateway utama untuk mengakses micro service lainnya. Server ini menyediakan REST API endpoint untuk setiap micro service.

## Auth
### Login
Endpoint login digunakan untuk mendapatkan access_token untuk mengakses endpoint lain.

Endpoint: `[POST] /auth/login`

Params:
- username
- password

Response:
- userid
- name
- session:
    - id
    - user_id
    - access_token
    - refresh_token
    - access_token_expires_at
    - refresh_token_expires_at

### Register
Endpoint register digunakan untuk mendaftarkan user baru.

Endpoint: `[POST] /auth/register`

Params:
- name
- username
- password

Respons:
- userid
- username
- name

## Payment
### Inquiry
Endpoint untuk melihat tagihan dari `bill_number` yang dikirim.

Auth:
- authorization: Bearer <access_token>

Endpoint: `[POST] /biller/inquiry`

Params:
- bill_number
- product_code

Response:
- inq_id
- bill_number
- product_code
- name
- total_amount

### Payment
Endpoint untuk melakukan pelunasan tagihan.

Auth:
- authorization: Bearer <access_token>

Endpoint: `[POST] /biller/payment`

Params:
- inq_id

Response:
- bill_number
- product_code
- name
- total_amount
- refference_number
- transaction_datetime