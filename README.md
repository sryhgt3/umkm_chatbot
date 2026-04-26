# UMKM Chatbot - Implementasi Webhook Telegram

Ini adalah implementasi webhook Telegram siap produksi menggunakan bahasa pemrograman Go dan framework Gin, dengan mengikuti prinsip *Clean Architecture*.

## Fitur
- POST `/webhook/telegram` untuk menerima update dari Telegram.
- Arsitektur modular (Handler, Service, Model).
- Konfigurasi menggunakan variabel lingkungan (environment variables).
- Integrasi Database PostgreSQL.
- Dockerized environment untuk kemudahan deployment.

## Prasyarat
Sebelum memulai, pastikan Anda telah menginstal:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Telegram Bot Token (dapatkan dari [@BotFather](https://t.me/botfather)).

## Instalasi & Setup (Docker)

1. **Clone repositori** ini ke direktori lokal Anda.
2. **Konfigurasi Environment**:
   Salin file `.env.example` menjadi `.env`:
   ```bash
   cp .env.example .env
   ```
   Buka file `.env` dan sesuaikan nilainya, terutama `BOT_TOKEN`:
   ```env
   PORT=8080
   BOT_TOKEN=your_telegram_bot_token_here
   DB_HOST=db
   DB_PORT=5432
   DB_USER=umkm
   DB_PASSWORD=Umkmajah13
   DB_NAME=umkm_db
   JWT_SECRET=UmkmChatbotSecret2024
   ```
3. **Jalankan Environment**:
   Gunakan Docker Compose untuk membangun dan menjalankan semua layanan (backend, database, pgAdmin):
   ```bash
   docker compose up --build -d
   ```

## Akses Layanan

- **Backend API**: `http://localhost:8080`
- **pgAdmin (Database Management)**: `http://localhost:5050`
  - **Email**: `AdminUmkm@Vory.com`
  - **Password**: `AdminVory`

### Cara Menghubungkan pgAdmin ke Database:
1. Buka `http://localhost:5050` dan login.
2. Klik kanan pada 'Servers' > 'Register' > 'Server...'.
3. Di tab **General**, beri nama (misal: `UMKM DB`).
4. Di tab **Connection**:
   - **Host name/address**: `db` (ini adalah nama service di docker-compose)
   - **Port**: `5432`
   - **Maintenance database**: `umkm_db`
   - **Username**: `umkm`
   - **Password**: `Umkmajah13`
5. Klik **Save**.

## Verifikasi
Untuk memastikan backend berjalan dengan baik, Anda dapat memeriksa health check (jika sudah diimplementasikan) atau mencoba mengakses root URL:
```bash
curl http://localhost:8080/health
```
*(Catatan: Pastikan endpoint `/health` sudah terdaftar di router Anda).*

## Menyiapkan Webhook
Agar dapat menerima pesan dari Telegram, Anda harus mengekspos server lokal ke internet (contohnya menggunakan `ngrok`) dan mendaftarkan URL webhook tersebut ke Telegram:
```bash
curl -X POST https://api.telegram.org/bot<TOKEN_BOT_ANDA>/setWebhook?url=<URL_PUBLIK_ANDA>/webhook/telegram
```

## Struktur Proyek
```bash
├── agent/                  # Dokumentasi & spesifikasi agen AI
├── cmd/
│   └── server/
│       └── main.go         # Entry point aplikasi & Setup Router
├── internal/
│   ├── client/             # Client pihak ketiga (Telegram API)
│   ├── config/             # Pengaturan konfigurasi & Env loader
│   ├── database/           # Koneksi database & Migrasi
│   ├── handler/            # HTTP Handlers (Webhook, Auth)
│   ├── middleware/         # Gin Middleware (Auth, Role check)
│   ├── model/              # Domain Models & Data structures
│   ├── repository/         # Data access layer (PostgreSQL)
│   ├── service/            # Business Logic layer
│   └── utils/              # Helper functions (Hashing, JWT)
├── migrations/             # File migrasi SQL
├── docs/                   # Dokumentasi API (Baru)
├── .env                    # Variabel lingkungan (Private)
├── .env.example            # Template variabel lingkungan
├── Dockerfile              # Konfigurasi Docker backend
├── docker-compose.yml      # Orchestration (App, DB, pgAdmin)
└── go.mod                  # Definisi modul Go
```
