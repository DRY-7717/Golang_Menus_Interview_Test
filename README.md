# Golang Menu Management API

👋 Selamat datang di **Golang Menu Management API**, sebuah aplikasi backend untuk manajemen menu menggunakan **Golang 1.25**, **Fiber**, dan **GORM** dengan PostgreSQL. API ini dirancang untuk performa tinggi, modular, dan siap dijalankan di Docker.

---

## 🛠️ Teknologi yang Digunakan

- **Golang 1.25** - Bahasa utama backend
- **Fiber** - Web framework cepat dan ringan
- **GORM** - ORM untuk integrasi dengan PostgreSQL
- **PostgreSQL** - Database relasional
- **Golang-Migrate** - Untuk migrasi database
- **Air** - Hot reload untuk development
- **Docker & Docker Compose** - Untuk containerization
- **Swagger** - Dokumentasi API interaktif

**Arsitektur**:  
- Menggunakan **layered architecture**:  
  - `Handler` → menangani request HTTP  
  - `Service` → logika bisnis  
  - `Repository` → komunikasi dengan database 


---


## ⚡ Setup

1. Clone repository:

```bash
git clone <repository-url>
cd <repository-folder>
```
2. Install dependencies:

```bash
go mod tidy
```

3. Buat file .env berdasarkan .env.example dan sesuaikan konfigurasi database.

```bash
cp .env.example .env
```
Untuk Development Gunakan port 8001 dan APP_ENV gunakan development seperti ini:

```bash
APP_ENV="development"
APP_PORT=8001
```

Untuk Production Gunakan port 8000 dan APP_ENV gunakan production seperti ini:

```bash
APP_ENV="production"
APP_PORT=8000
```

4. Menjalakan migrasi database:

Setelah menyesuaikan konfigurasi database di file .env anda bisa menjalankan migrasinya.

  - Untuk menjalankan migrasinya terlebih dahulu harus menginstall golang migrate dengan cara:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
- Kemudian Kalo sudah terinstall, tinggal menjalankan migrasinya dengan menjalankan peritah ini diterminal:

```bash
migrate -database "postgres://<username>:<password>@<host>:<port>/<database name>?sslmode=disable" -path db/migrations up
```

Untuk melihat cara penggunaan golang migrate bisa dilihat langsung didokumentasi resminya (https://github.com/golang-migrate/migrate)

## 🚀 Menjalankan Aplikasi

### Development

Menjalankan aplikasi pada tahap developement ada 2 cara menggunakan **golang air** untuk hot reload atau tidak dengan menggunakan golang air.

⚠️ Pastikan database sudah hidup dan konfigurasi database di .env sudah disesuaikan dan migrasi sudah dijalankan

**Tanpa Golang Air**

```bash
go run .
```

**Menggunakan Golang Air**

⚠️ Pastikan untuk menggunakan golang air harus terinstall golang air nya terlebih dahulu dengan cara:

```bash
go install github.com/air-verse/air@latest
```

Jika sudah terinstall atau sudah ada golang air tinggal jalankan perintah:

```bash
air
```

### Production

Menjalankan aplikasi untuk tahap production.

⚠️ Pastikan database sudah hidup dan konfigurasi database di .env sudah disesuaikan dan migrasi sudah dijalankan

Untuk menjalankan golang ditahap production, anda bisa menjalankan perintah build ini diterminal anda

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o go_menus .
```
nama `go_menus` bisa diganti dengan nama yang lain sesuai keinginan anda.

Kemudian kalau filenya sudah dibuild, bisa jalankan perintah ini diterminal untuk menjalankan aplikasi yang sudah dibuild:

```bash
./go_menus
```


## 🐳 Menjalankan Aplikasi Menggunakan Docker

Aplikasi dijalankan menggunakan docker compose. Pada docker compose terdapat beberapa configurasi environment dari masing masing Service


**Service App**

Konfigurasi env bisa disesuaikan dengan konfigurasi anda. Untuk tahap development APP_ENV diisikan `development` untuk tahap production diisikan `production`  dan APP_PORT tahap development diisikan `8001` dan tahap production `8000`, seperti ini:

```bash
  environment:
      - APP_ENV=development
      - APP_PORT=8001
      - DATABASE_PORT=5432
      - DATABASE_HOST=db
      - DATABASE_USER=admindev
      - DATABASE_PASSWORD=4dm1n123
      - DATABASE_NAME=db_gomenus
      - DATABASE_MAX_OPEN_CONNECTION=10
      - DATABASE_MAX_IDLE_CONNECTION=20
```

untuk konfigurasi seperti env `DATABASE_HOST` itu harus diisikan dengan nama dari service database, kalau disini nama service database saya yaitu `db` dan untuk environment seperti `DATABASE_USER`, `DATABASE_PASSWORD` dan `DATABASE_NAME` itu mengikuti atau harus sama dengan environtment yang ada di service database.


**Service Database**

Untuk environtment database bisa disesuaikan dengan konfigurasi anda, contoh:

```bash
 environment:
      - POSTGRES_USER=admindev
      - POSTGRES_PASSWORD=4dm1n123
      - POSTGRES_DB=db_gomenus
```

**Service Migrate**

Service migrate ini berfungsi untuk melakukan migrasi otomatis saat container dijalankan semua, pada configurasi migrate ini terdapat entrypoint untuk menjalankan migrasi secara otomatis dan ini bisa disesuaikan dengan konfigurasi anda, contoh:

```bash
 entrypoint: [
        "sh",
        "-c",
        "until nc -z db 5432; do echo 'Waiting for Postgres...'; sleep 2; done; \
        echo 'Postgres is ready!'; \
        migrate -path=/migrations -database postgres://<username>:<password>@<host>:<port>/<database name>?sslmode=disable up",
      ]
```

**Service PGAdmin**

Untuk environtment pgadmin bisa disesuaikan dengan konfigurasi anda, contoh:

```bash
 environment:
      - PGADMIN_DEFAULT_EMAIL=admin@gmail.com
      - PGADMIN_DEFAULT_PASSWORD=okeokeoke
```

⚠️ Untuk melihat seluruh konfigurasi docker compose baik development atau production saya bisa dilihat didalam file dev.docker-comppose.yaml dan prod.docker-compose.yaml

### Menjalankan Aplikasi

Setelah mengkonfigurasi docker compose sesuai kebutuhan anda atau bisa juga memakai konfigurasi yang sudah saya buat. Bisa langsung menjalankan perintah ini:

**Development**

Jalankan build image terlebih dahulu dengan cara menjalankan perintah ini di terminal:

```bash
docker compose -f dev.docker-compose.yaml build --no-cache
```

Kemudian baru jalankan semua service yang ada di docker compose dengan cara menjalankan perintah ini di terminal:

```bash
docker compose -f dev.docker-compose.yaml up -d
```

**Production**

Jalankan build image terlebih dahulu dengan cara menjalankan perintah ini di terminal:

```bash
docker compose -f prod.docker-compose.yaml build --no-cache
```

Kemudian baru jalankan semua service yang ada di docker compose dengan cara menjalankan perintah ini di terminal:

```bash
docker compose -f prod.docker-compose.yaml up -d
```


## 🔗 API Endpoints

Berikut daftar endpoint yang tersedia untuk mengelola menu items, mengecek status server, dan dokumentasi API:

| Method | Endpoint                | Description                                                     |
|--------|-------------------------|-----------------------------------------------------------------|
| GET    | `/api/check`            | ✅ Check if the server is running                               |
| GET    | `/api/swagger`          | 📄 Open Swagger API documentation                               |
| GET    | `/api/menus`            | 📝 Get all menu items (tree structure)                          |
| GET    | `/api/menus/:id`        | 📝 Get single menu item                                         |
| POST   | `/api/menus`            | 📝 Create new menu item                                         |
| PUT    | `/api/menus/:id`        | 📝 Update menu item                                             |
| DELETE | `/api/menus/:id`        | 📝 Delete menu item (and children)                              |
| PATCH  | `/api/menus/:id/move`   | 📝 Move menu item to different parent                           |
| PATCH  | `/api/menus/:id/reorder`| 📝 Reorder menu item within same level                          |
