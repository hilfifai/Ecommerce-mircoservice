# Ecommerce-mircoservice
Prototype project - Backend : Golang - Redis - n8n - Appsmith - Grafana &amp; Frontend : Laravel - VUE.js - Tailwindcss

### **Arsitektur**

*   Frontend: Laravel (sebagai server untuk SPA) + Vue.js (sebagai SPA) + Directus (sebagai cache/sumber data produk untuk frontend) + tailwindcss (sebagai styling).
*   Backend: Golang (sebagai API utama) + Redis (sebagai message queue) + PostgreSQL (sebagai database utama).
*   Admin Panel: Appsmith (untuk manajemen produk dan data master lainnya).
*   Monitoring: Grafana (dengan data source dari Prometheus).
*   Automation: n8n (untuk sinkronisasi data dari Golang ke Directus).

### **Penjelasan Alur Kerja Detail**

1.  **Alur Manajemen Produk (Admin)**
    *   Admin membuka panel Appsmith yang sudah terintegrasi.
    *   Admin membuat atau mengubah data produk melalui form di Appsmith.
    *   Saat tombol "Simpan" ditekan, Appsmith memanggil endpoint `POST /api/v1/products` atau `PUT /api/v1/products/{id}` di **Golang API**.
    *   Golang API menyimpan produk ke database PostgreSQL dan kemudian memicu sebuah webhook yang ditujukan ke n8n.
    *   n8n menerima data produk dari webhook tersebut, lalu menggunakan API Directus untuk membuat atau memperbarui item produk di koleksi `products`. Dengan cara ini, Directus selalu memiliki data produk yang paling update.

2.  **Alur Pengguna (Melihat & Memesan Produk)**
    *   Pengguna membuka website. Laravel akan menyajikan file HTML dasar dan aset JavaScript (Vue.js SPA).
    *   Aplikasi Vue.js melakukan inisialisasi dan memanggil API Directus untuk mendapatkan daftar produk (`GET /items/products`). Ini sangat cepat karena Directus dioptimalkan untuk pengiriman konten.
    *   Pengguna melakukan Login/Register. Permintaan ini dikirim dari Vue.js langsung ke Golang API (`POST /api/v1/auth/login`). Golang akan memvalidasi dan mengembalikan token (misalnya JWT).
    *   Pengguna menambahkan produk ke keranjang (`cart`) dan melakukan checkout (`order`).
    *   Permintaan pembuatan order (`POST /api/v1/orders`) dikirim dari Vue.js ke Golang API dengan membawa token otentikasi.

3.  **Alur Pemrosesan Pesanan (Backend)**
    *   Golang API menerima permintaan order. Daripada langsung memproses, API hanya melakukan validasi data awal (misalnya, format benar) dan langsung memasukkan detail order ke dalam Redis Queue dengan topik `new_orders`. Ini membuat response ke user menjadi sangat cepat.
    *   Sebuah proses terpisah, Golang Worker, terus-menerus memantau queue `new_orders` di Redis.
    *   Ketika ada order baru, Worker mengambilnya dari queue.
    *   Worker melakukan logika bisnis yang lebih berat:
        1.  Memeriksa ketersediaan stok produk di PostgreSQL.
        2.  Jika stok cukup, kurangi stok produk dan simpan data order ke tabel `orders` di PostgreSQL.
        3.  Jika stok tidak cukup, tandai order sebagai gagal dan bisa mengirim notifikasi (misalnya, melalui email atau websocket).
    *   Seluruh proses ini terjadi secara asinkron, tidak membebani API utama dan lebih tangguh terhadap kegagalan.

## Instalasi & Menjalankan Proyek
### Inisialisasi Frontend

```bash
cd ecommerce-frontend
composer install
npm install
cd ..
```

### Layanan dengan Docker

```bash
docker-compose up -d
```
atau bisa pake powershell 
```bash
.\manage.ps1
```
### Jalankan Server Frontend Laravel & Vite

```bash
cd ecommerce-frontend
npm run dev
php artisan serve
```
### Konfigurasi Postgres

1. Buka docker-compose.yml.
2. Ubah nilai POSTGRES_USER, POSTGRES_PASSWORD, dan POSTGRES_DB di dalam blok environment untuk setiap layanan yang relevan.
```bash
services:
  postgres:
    environment:
      POSTGRES_USER: myuser           
      POSTGRES_PASSWORD: mysecretpass 
      POSTGRES_DB: myappdb            

  backend-api:
    environment:
      DB_USER: myuser                  
      DB_PASSWORD: mysecretpass        
      DB_NAME: myappdb                 
````

ubah point point diatas
