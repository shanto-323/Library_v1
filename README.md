## Project Layout

The starter kit uses the following project layout:
```
library/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── models/
│   │   ├── models.go
│   ├── repositories/
│   │   ├── student_repo.go
│   │   ├── book_repo.go
│   │   ├── borrowed_book_repo.go
│   │   ├── genre_repo.go
│   │   ├── author_repo.go
│   ├── services/
│   │   ├── student_service.go
│   │   ├── book_service.go
│   │   ├── borrowed_book_service.go
│   │   ├── genre_service.go
│   │   ├── author_service.go
│   └── handlers/
│       ├── api.go
│       ├── student_handler.go
│       ├── book_handler.go
│       ├── borrowed_book_handler.go
│       ├── genre_handler.go
│       ├── author_handler.go
├── migrations/
│   └── 001_create_tables.sql
├── seeds/
│   └── 001_seed.sql
├─  .gitignore
├─  Makefile
├─  go.sum
└── go.mod
```

### Change (api) with
- student
- book
- borrowed-books
- author
- genre

#### Get items
```http
  GET /api/items
```

#### Get item
```http
  GET /api/items/${id}
```

#### Create item
```http
  GET /api/new
```

#### Delete item
```http
  GET /api/delete/${id}
```

#### Update item
```http
  GET /api/update/${id}
```
