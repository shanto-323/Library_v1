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
│   ├── services/
│   │   ├── student_service.go
│   └── handlers/
│       ├── api.go
│       ├── student_handler.go
├── migrations/
│   └── 001_create_tables.sql
├─  .gitignore
├─  Makefile
├─  go.sum
└── go.mod
```


# Create Schema

Run an postgres containter on your local matchine 

**Make a migration folder in your postgres container**

```
docker exec -it <containter_id> mkdir -p /migrations
```
**Make schema**
```
docker cp /path/to/your/go/project/migrations/001_create_tables.sql 9a3d46eb3b8d:/migrations/001_create_tables.sql
````
