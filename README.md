<div align="center">

![Kantoor](https://user-images.githubusercontent.com/62880102/177792650-f33f4a7d-48f3-4145-a6b9-4799bfd3ac33.png)

# Kantoor — Office Booking System API

**Backend REST API untuk platform penyewaan ruang kantor di Jakarta, dibangun dengan Go.**

![Go](https://img.shields.io/badge/Go-1.18-00ADD8?logo=go&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-v4-1f6feb)
![MySQL](https://img.shields.io/badge/MySQL-GORM-4479A1?logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-GPL--3.0-blue)

[Dokumentasi API (Swagger)](https://app.swaggerhub.com/apis/45-OfficeBooking/Office-Booking/1.0.0/) ·
[Demo Frontend](https://front-end-vue-office-booking-system-ku6sny2of-didiroyadi123.vercel.app) ·
[Demo Live Chat](https://gauri-golang-chat.herokuapp.com/)

</div>

---

## Daftar Isi

- [Tentang Proyek](#tentang-proyek)
- [Fitur](#fitur)
- [Tech Stack](#tech-stack)
- [Arsitektur](#arsitektur)
- [Struktur Folder](#struktur-folder)
- [Skema Database](#skema-database)
- [Memulai](#memulai)
  - [Prasyarat](#prasyarat)
  - [Instalasi](#instalasi)
  - [Konfigurasi Environment](#konfigurasi-environment)
  - [Menjalankan Aplikasi](#menjalankan-aplikasi)
  - [Menjalankan dengan Docker](#menjalankan-dengan-docker)
- [Dokumentasi API](#dokumentasi-api)
  - [Autentikasi](#autentikasi)
  - [Daftar Endpoint](#daftar-endpoint)
  - [Contoh Request](#contoh-request)
- [CI/CD & Deployment](#cicd--deployment)
- [Roadmap](#roadmap)
- [Kontribusi](#kontribusi)
- [Lisensi](#lisensi)

---

## Tentang Proyek

**Kantoor** adalah platform berbasis web untuk menyewa ruang kantor di kawasan bisnis Jakarta. Tujuannya mempermudah pengguna mencari ruang kerja yang sesuai, mulai dari melihat daftar gedung, memfilter berdasarkan harga dan lokasi, membaca ulasan, melihat fasilitas di sekitar gedung, hingga melakukan pemesanan.

Repositori ini berisi **layanan backend (REST API)**-nya. Antarmuka pengguna dibangun terpisah menggunakan Vue.js (lihat [Demo Frontend](https://front-end-vue-office-booking-system-ku6sny2of-didiroyadi123.vercel.app)).

## Fitur

**Customer**
- Registrasi dan login dengan JWT
- Melihat daftar gedung, detail gedung, dan filter berdasarkan harga
- Melihat jenis gedung dan fasilitas terdekat (*nearby*)
- Membuat dan melihat ulasan (*review*)
- Melihat riwayat dan detail booking
- Melihat dan memperbarui profil

**Admin**
- CRUD data gedung, jenis gedung, dan fasilitas terdekat
- Manajemen booking (buat, lihat, ubah, hapus)
- Manajemen pengguna dan profil admin
- Moderasi ulasan

## Tech Stack

| Kategori | Teknologi |
|---|---|
| Bahasa | [Go](https://go.dev/doc/) 1.18 |
| Web framework | [Echo v4](https://echo.labstack.com/) |
| ORM & migrasi | [GORM](https://gorm.io/docs/) (auto-migrate) |
| Database | [MySQL](https://dev.mysql.com/doc/) (driver PostgreSQL & SQL Server juga tersedia di `go.mod`) |
| Autentikasi | JWT ([golang-jwt/jwt](https://github.com/golang-jwt/jwt)) |
| Konfigurasi | [godotenv](https://github.com/joho/godotenv) |
| Kontainerisasi | [Docker](https://docs.docker.com/) |
| CI/CD | GitHub Actions |
| Hosting | [DigitalOcean](https://www.digitalocean.com/) |
| Dokumentasi API | Swagger / OpenAPI (SwaggerHub) |

## Arsitektur

Proyek ini menerapkan pola **Clean Architecture** dengan pemisahan tanggung jawab yang jelas antar lapisan:

```
HTTP Request
    │
    ▼
┌─────────────┐    ┌───────────┐    ┌──────────────┐    ┌──────────┐
│ Controllers │───▶│  Usecase  │───▶│  Repository  │───▶│ Database │
│ (Echo/HTTP) │    │ (bisnis)  │    │ (GORM/query) │    │  (MySQL) │
└─────────────┘    └───────────┘    └──────────────┘    └──────────┘
        ▲                 ▲                  ▲
        └─────────────────┴──────────────────┘
                     Domain
       (entity, interface, request & response DTO)
```

| Lapisan | Tanggung jawab |
|---|---|
| `domain` | Entity, interface repository/usecase, serta DTO request & response |
| `controllers` | Handler HTTP dan registrasi route per modul |
| `usecase` | Logika bisnis aplikasi |
| `repository` | Akses data ke database menggunakan GORM |
| `delivery/http` | Middleware (logger, JWT) dan helper pembuat token |
| `app` | Inisialisasi dependensi, koneksi database, dan server |

## Struktur Folder

```
Office-Booking/
├── .github/workflows/
│   └── ci.yml                  # Pipeline build & deploy
├── app/
│   ├── config/config.go        # Koneksi database & auto-migrate
│   └── routes.go               # Dependency injection & start server
├── controllers/                # Handler HTTP per modul
│   ├── booking/
│   ├── gedung/
│   ├── jenisgedung/
│   ├── nearby/
│   ├── review/
│   └── users/
├── delivery/http/
│   ├── helper/jwt.go           # Pembuatan token JWT
│   └── middleware/             # Logger & auth middleware
├── domain/                     # Entity, interface, request/response DTO
│   ├── booking/
│   ├── gedung/
│   ├── jenisgedung/
│   ├── nearby/
│   ├── review/
│   └── users/
├── repository/                 # Implementasi akses data (GORM)
├── usecase/                    # Logika bisnis
├── dockerfile
├── go.mod
├── go.sum
├── main.go                     # Entry point
└── LICENSE
```

## Skema Database

Tabel dibuat otomatis melalui `AutoMigrate` GORM saat aplikasi pertama kali dijalankan.

| Tabel | Kolom utama |
|---|---|
| `users` | `id`, `email`, `name`, `fullname`, `alamat`, `phone`, `password`, `id_booking` |
| `gedungs` | `id`, `name`, `location`, `price`, `latitude`, `longitude`, `description`, `id_booking` |
| `jenisgedungs` | `id`, `jenis`, `id_gedung` |
| `nearbies` | `id`, `namefacilities`, `jenis`, `jarak`, `latitude`, `longtitude`, `id_gedung` |
| `reviews` | `id`, `img`, `rating`, `description`, `id_gedung` |
| `bookings` | `id`, `status`, `bookingcode`, `totalbooking`, `orderdate`, `checkin`, `checkout`, `fullname`, `phone` |

Seluruh tabel (kecuali `reviews`) juga memiliki `created_at`, `updated_at`, dan `deleted_at` (*soft delete*).

**Relasi**
- Satu `gedung` memiliki banyak `review`, `nearby`, dan `jenisgedung` (melalui `id_gedung`)
- Satu `booking` terhubung ke `user`, `gedung`, dan `jenis` (melalui `id_booking`)

## Memulai

### Prasyarat

- [Go](https://go.dev/dl/) 1.18 atau lebih baru
- [MySQL](https://dev.mysql.com/downloads/) 8.x (atau MariaDB yang kompatibel)
- [Git](https://git-scm.com/)
- [Docker](https://docs.docker.com/get-docker/) *(opsional)*

### Instalasi

```bash
# 1. Clone repositori
git clone https://github.com/45-Office-Booking-System/Office-Booking.git
cd Office-Booking

# 2. Unduh dependensi
go mod download
```

### Konfigurasi Environment

Buat file `.env` di root proyek:

```env
# Database
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=your_password
DB_NAME=office_booking

# JWT
JWT_SECRET=ganti_dengan_secret_yang_kuat
```

> [!IMPORTANT]
> Pada versi kode saat ini, kredensial database dan JWT secret masih ditulis langsung (*hardcoded*) di `app/config/config.go`, `delivery/http/middleware/middlewares.go`, dan `delivery/http/helper/jwt.go`. Sebelum menjalankan atau men-deploy sendiri, pindahkan nilai tersebut ke environment variable seperti di atas (paket `godotenv` sudah tersedia di `go.mod`). Lihat [Roadmap](#roadmap).

Buat database kosong terlebih dahulu; tabel akan dibuat otomatis:

```sql
CREATE DATABASE office_booking CHARACTER SET utf8mb4;
```

### Menjalankan Aplikasi

```bash
go run main.go
```

Server berjalan di `http://localhost:8080`.

Untuk membuat binary:

```bash
go build -o main .
./main
```

### Menjalankan dengan Docker

```bash
# Build image
docker build -t office-booking .

# Jalankan kontainer
docker run -d -p 8080:8080 --name office-booking office-booking
```

Atau gunakan image yang sudah dipublikasikan:

```bash
docker pull officebooking/officebooking-go
docker run -d -p 8080:8080 --name office-booking officebooking/officebooking-go
```

## Dokumentasi API

Dokumentasi interaktif lengkap tersedia di [SwaggerHub](https://app.swaggerhub.com/apis/45-OfficeBooking/Office-Booking/1.0.0/).

**Base URL (lokal):** `http://localhost:8080`

### Autentikasi

Endpoint yang dilindungi memerlukan token JWT (algoritma HS256, berlaku 48 jam) pada header:

```
Authorization: Bearer <token>
```

Token diperoleh dari endpoint `POST /login`. Payload token berisi `userID`, `email`, `name`, dan `phone`.

### Daftar Endpoint

Kolom **Auth** menunjukkan endpoint yang saat ini diproteksi middleware JWT di kode.

#### Autentikasi & Pengguna

| Method | Endpoint | Auth | Deskripsi |
|---|---|:---:|---|
| POST | `/register` | – | Registrasi akun baru |
| POST | `/login` | – | Login dan mendapatkan token JWT |
| POST | `/customer` | – | Membuat data customer |
| GET | `/customer/profile/:id` | ✅ | Lihat profil customer |
| PUT | `/customer/profile/:id` | – | Ubah profil customer |
| GET | `/admin/users` | – | Daftar seluruh pengguna |
| GET | `/admin/user/:id` | – | Detail pengguna |
| PUT | `/admin/user/:id` | – | Ubah data pengguna |
| PUT | `/admin/profile/:id` | – | Ubah profil admin |
| DELETE | `/admin/user/:id` | – | Hapus pengguna |
| POST | `/admin/booking/user` | – | Kaitkan pengguna ke booking |

#### Gedung

| Method | Endpoint | Auth | Deskripsi |
|---|---|:---:|---|
| GET | `/customer/gedungs` | – | Daftar gedung |
| GET | `/customer/gedung/price` | – | Filter gedung berdasarkan harga |
| GET | `/customer/gedung/:id` | – | Detail gedung |
| GET | `/admin/gedungs` | – | Daftar gedung (admin) |
| GET | `/admin/gedung/:id` | – | Detail gedung (admin) |
| POST | `/admin/gedung` | – | Tambah gedung |
| PUT | `/admin/gedung/:id` | – | Ubah gedung |
| DELETE | `/admin/gedung/:id` | – | Hapus gedung |

#### Jenis Gedung

| Method | Endpoint | Auth | Deskripsi |
|---|---|:---:|---|
| GET | `/jenisgedung` | – | Daftar jenis gedung |
| GET | `/jenisgedung/:id` | ✅ | Detail jenis gedung |
| GET | `/admin/jenisgedung` | – | Daftar jenis gedung (admin) |
| GET | `/admin/jenisgedung/:id` | – | Detail jenis gedung (admin) |
| POST | `/admin/jenisgedung` | – | Tambah jenis gedung |
| PUT | `/admin/jenisgedung/:id` | – | Ubah jenis gedung |
| DELETE | `/admin/jenisgedung/:id` | – | Hapus jenis gedung |

#### Fasilitas Terdekat (Nearby)

| Method | Endpoint | Auth | Deskripsi |
|---|---|:---:|---|
| GET | `/customer/nearby` | – | Daftar fasilitas terdekat |
| GET | `/customer/nearby/:id` | – | Detail fasilitas terdekat |
| GET | `/admin/nearby` | – | Daftar fasilitas (admin) |
| GET | `/admin/nearby/:id` | – | Detail fasilitas (admin) |
| POST | `/admin/nearby/` | – | Tambah fasilitas |
| PUT | `/admin/nearby/:id` | – | Ubah fasilitas |
| DELETE | `/admin/nearby/:id` | – | Hapus fasilitas |

#### Review

| Method | Endpoint | Auth | Deskripsi |
|---|---|:---:|---|
| POST | `/customer/review/` | – | Tambah ulasan |
| GET | `/customer/review/` | – | Daftar ulasan |
| GET | `/review/:id` | – | Detail ulasan |
| GET | `/admin/review` | – | Daftar ulasan (admin) |
| GET | `/admin/review/:id` | – | Detail ulasan (admin) |
| DELETE | `/admin/review/:id` | – | Hapus ulasan |

#### Booking

| Method | Endpoint | Auth | Deskripsi |
|---|---|:---:|---|
| GET | `/customer/bookings` | – | Daftar booking customer |
| GET | `/customer/booking/:id` | – | Detail booking customer |
| GET | `/admin/bookings` | – | Daftar seluruh booking |
| GET | `/admin/booking/:id` | – | Detail booking |
| POST | `/admin/booking` | – | Buat booking |
| PUT | `/admin/booking/:id` | – | Ubah booking |
| DELETE | `/admin/booking/:id` | – | Hapus booking |

### Contoh Request

**Registrasi**

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "budi@example.com",
    "name": "budi",
    "fullname": "Budi Santoso",
    "alamat": "Jakarta Selatan",
    "phone": "081234567890",
    "password": "rahasia123",
    "konfirmpassword": "rahasia123"
  }'
```

**Login**

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email": "budi@example.com", "password": "rahasia123"}'
```

**Request ke endpoint terproteksi**

```bash
curl http://localhost:8080/customer/profile/1 \
  -H "Authorization: Bearer <token>"
```

**Tambah gedung**

```bash
curl -X POST http://localhost:8080/admin/gedung \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Menara Sudirman",
    "location": "Jl. Jend. Sudirman, Jakarta Pusat",
    "price": "5000000",
    "latitude": "-6.2088",
    "longitude": "106.8456",
    "description": "Ruang kantor modern di kawasan SCBD"
  }'
```

## CI/CD & Deployment

Pipeline GitHub Actions (`.github/workflows/ci.yml`) berjalan setiap ada *push* ke branch `master`:

1. **Build** — build image Docker dan push ke Docker Hub (`officebooking/officebooking-go`)
2. **Deploy** — koneksi SSH ke server, lalu menarik image terbaru dan menjalankan ulang kontainer pada port `8080`

**Secrets yang diperlukan** (Settings → Secrets and variables → Actions):

| Secret | Keterangan |
|---|---|
| `DOCKERHUB_USERNAME` | Username Docker Hub |
| `DOCKERHUB_PASSWORD` | Password / access token Docker Hub |
| `HOST` | Alamat server tujuan deploy |
| `USERNAME` | User SSH server |
| `KEY` | Private key SSH |
| `PORT` | Port SSH |

## Roadmap

- [ ] Pindahkan kredensial database dan JWT secret ke environment variable
- [ ] Terapkan JWT middleware dan pembatasan role (admin/customer) pada seluruh route `/admin/*` dan `/customer/*`
- [ ] Hash password pengguna (misalnya bcrypt)
- [ ] Tambahkan unit test untuk usecase dan repository (`go-sqlmock` sudah tersedia)
- [ ] Tambahkan validasi input pada request
- [ ] Tambahkan pagination dan filter lokasi pada daftar gedung
- [ ] Konfigurasi CORS yang lebih ketat untuk production
- [ ] Sediakan `docker-compose.yml` (API + MySQL) untuk setup lokal satu perintah
- [ ] Tambahkan file `.env.example`

## Kontribusi

Kontribusi sangat diterima.

1. Fork repositori ini
2. Buat branch fitur: `git checkout -b feature/nama-fitur`
3. Commit perubahan: `git commit -m "feat: deskripsi singkat"`
4. Push ke branch: `git push origin feature/nama-fitur`
5. Buka Pull Request

## Lisensi

Didistribusikan di bawah lisensi **GNU General Public License v3.0**. Lihat berkas [LICENSE](LICENSE) untuk detail selengkapnya.

## Tim

Dikembangkan oleh **Tim 45 — Office Booking System**.
