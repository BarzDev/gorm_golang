# Library API

RESTful API sederhana untuk mengelola buku, penulis, dan kategori menggunakan **Go**, **Gin**, dan **GORM (PostgreSQL)**.

---

## 📚 Fitur

- CRUD **Authors** (Penulis)
- CRUD **Categories** (Kategori buku)
- CRUD **Books** (Buku) dengan relasi Author & Category
- Endpoint **Books** mendukung:
  - List semua buku
  - Filter berdasarkan `author_id` dan `category_id`
  - Get buku berdasarkan `id`
- JSON response **nested** untuk Author & Category
- Validasi input dan error handling

---

## 🛠️ Teknologi

- [Go](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM ORM](https://gorm.io/)
- PostgreSQL
- Git

---

## ⚡ Struktur Project

```
library-api/
├── config/            # konfigurasi database & environment
├── delivery/          # HTTP layer (controller + server)
│   └── controller/    # Gin controller untuk setiap resource
├── model/             # Struct model (Book, Author, Category)
├── repository/        # Repository layer (DB access)
├── usecase/           # Business logic / usecase layer
├── shared/             # Reusable utilities
│   └── common/         # Helper & fungsi global (response formatter, error wrapper, middleware)
│   └── shared_model/   # Struct umum (response, paging, status)
└── main.go            # entrypoint server
```

---

## 🚀 Jalankan Server

1. **Clone repository**  
```bash
git clone https://github.com/BarzDev/gorm_golang.git
cd gorm_golang
```

2. **Install dependencies**  
```bash
go mod tidy
```

3. **Jalankan server**  
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

## 📌 Endpoint API

### Authors
- `GET /authors` → List semua penulis  
- `GET /authors/:id` → Detail penulis berdasarkan ID  
- `POST /authors` → Tambah penulis baru  
- `PUT /authors/:id` → Update penulis berdasarkan ID  
- `DELETE /authors/:id` → Hapus penulis berdasarkan ID  

---

### Categories
- `GET /categories` → List semua kategori  
- `GET /categories/:id` → Detail kategori berdasarkan ID  
- `POST /categories` → Tambah kategori baru  
- `PUT /categories/:id` → Update kategori berdasarkan ID  
- `DELETE /categories/:id` → Hapus kategori berdasarkan ID   

---

### Books
- `GET /books` → List semua buku  
  - Optional query params:  
    - `author_id` → filter berdasarkan penulis  
    - `category_id` → filter berdasarkan kategori  
  - Contoh:  
    - `/books?author_id=1` → filter by author  
    - `/books?category_id=2` → filter by category  
    - `/books?author_id=1&category_id=2` → filter keduanya  
- `GET /books/:id` → Detail buku berdasarkan ID  

**Contoh response JSON:**

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "title": "Harry Potter and the Philosopher's Stone",
      "author": {
        "id": 1,
        "name": "J.K. Rowling",
        "bio": "Penulis asal Inggris, terkenal dengan seri Harry Potter."
      },
      "category": {
        "id": 1,
        "name": "Fantasy"
      },
      "published_year": 1997,
      "price": 150000.00,
      "stock": 25
    }
  ]
}
```

---

## 🧩 License

Barz License © 2025

