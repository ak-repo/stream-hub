stream-hub/
├── go.mod
├── go.sum
│
├── cmd/ # Entry points
│ ├── auth/
│ │ └── main.go
│ ├── file/
│ │ └── main.go
│ ├── chat/ ← ADD THIS
│ │ └── main.go
│ ├── notification/
│ │ └── main.go
│ └── gateway/
│ └── main.go
│
├── services/ ← Changed from services/
│ ├── auth/
│ │ ├── handler/
│ │ ├── service/
│ │ ├── repository/
│ │ ├── middleware/ ← Move JWT here
│ │ └── model/
│ │
│ ├── file/
│ │ ├── handler/
│ │ ├── service/
│ │ ├── repository/
│ │ ├── storage/ ← MinIO client
│ │ └── model/
│ │
│ ├── chat/ ← ADD THIS
│ │ ├── handler/
│ │ ├── service/
│ │ ├── repository/
│ │ └── model/
│ │
│ ├── notification/
│ │ ├── handler/
│ │ ├── service/
│ │ └── model/
│ │
│ └── gateway/
│ ├── router.go
│ ├── client/ ← gRPC clients
│ └── middleware/
│
├── pkg/ ← Moved from services/pkg/
│ ├── database/
│ │ └── postgres.go
│ ├── logger/
│ │ └── logger.go
│ ├── config/
│ │ └── config.go
│ ├── validator/
│ │ └── validator.go
│ └── errors/
│ └── errors.go
│
├── api/
│ └── proto/
│ ├── auth.proto
│ ├── file.proto
│ ├── chat.proto ← ADD THIS
│ ├── notification.proto
│ └── common.proto
│
├── gen/ ← Generated code
│ └── go/
│ ├── auth/
│ ├── file/
│ ├── chat/ ← ADD THIS
│ ├── notification/
│ └── common/
│
├── migrations/
│ ├── 001_create_users.sql
│ ├── 002_create_files.sql
│ ├── 003_create_messages.sql ← ADD THIS
│ └── 004_create_notifications.sql
│
├── docker/ ← Better name than deployments/
│ ├── docker-compose.yml
│ ├── auth.Dockerfile
│ ├── file.Dockerfile
│ ├── chat.Dockerfile ← ADD THIS
│ ├── notification.Dockerfile
│ └── gateway.Dockerfile
│
├── configs/
│ ├── .env.example ← Single file better for MVP
│ └── docker.env
│
├── scripts/
│ ├── proto-gen.sh
│ ├── migrate.sh
│ └── seed.sh
│
└── docs/
├── architecture.md
├── api-design.md
└── sequence-diagrams.md

// DB manual check
docker exec -it streamhub-postgres psql -U streamhub_user -d streamhub
docker exec -t streamhub-postgres pg_dump -U streamhub_user streamhub > backup.sql


//after clone
docker compose up -d
docker exec -i streamhub-db psql -U streamhub_user streamhub < backup.sql
