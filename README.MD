# Golang API Login-Logout-Transaction

### Instalasi
Repo ini menggunakan docker. Jadi untuk menjalankan program hanya memerlukan command
```bash
docker-compose up -d
```

### Contoh request
- Login `GET` ke alamat `http://localhost:8000/api/login`  
Pada postman, navigasi ke `Body > raw > JSON` dan masukkan
    ```json
    {
        "username": "hero",
        "password": "hero"
    }
    ```
    response  
    ```json
    {
        "type": "login",
        "message": "success"
    }
    ```
    jika login berlebihan, maka muncul response  
    ```json
    {
        "type": "login",
        "message": "user already login"
    }
    ```
- Logout `GET` ke alamat `http://localhost:8000/api/logout`  
Pada postman, navigasi ke `Body > raw > JSON` dan masukkan
    ```json
    {
        "username": "hero",
        "password": "hero"
    }
    ```
    response  
    ```json
    {
        "type": "logout",
        "message": "success"
    }
    ```
    jika logout berlebihan, maka muncul response  
    ```json
    {
        "type": "logout",
        "message": "user already logout"
    }
    ```
- Payment `POST` ke alamat `http://localhost:8000/api/payment`  
pada postman, navigasi ke `Body > raw > JSON` dan masukkan
    ```json
    {
        "from": "hero",
        "to": "zeus",
        "amount": 10000
    }
    ```
    pastikan untuk login terlebih dulu, lalu lanjutkan untuk melakukan transfer.  
    response jika sudah login
    ```json
    {
        "type": "payment",
        "message": "transaction success"
    }
    ```
    response jika sebelum login
    ```json
    {
        "type": "payment",
        "message": "user need to login"
    }
    ```

### Database
untuk melihat data apa saja yang sudah dibuat dari aktivitas API, 
bisa melihatnya dengan melihat isi dari database. 
Untuk mengakses database, bisa mengunjungi `localhost:8080` nanti akan 
menemukan tampilan form login adminer. Isi dengan data berikut:
```
sistem  : PostgreSQL
Server  : db
pengguna: admin
Sandi   : admin
```