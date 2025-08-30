```
social-api/
│
├── cmd/                    # Main application entry point(s)
│   └── server/             # `main.go` lives here
│       └── main.go
│
├── internal/               # Application internal code
│   ├── config/             # Config loading (env, yaml, etc.)
│   ├── db/                 
│   │   ├── migrations/     # SQL migration files (used by migrate)
│   │   └── postgres.go     # DB connection logic
│   │
│   ├── handler/            # HTTP handlers
│   │   ├── auth/           # Login, signup handlers
│   │   ├── user/           # Profile routes
│   │   ├── post/           # Post creation, listing
│   │   └── comment/        # Comments logic
│   │
│   ├── middleware/         # JWT auth, logging, CORS, etc.
│   │
│   ├── model/              # Structs for DB & API (User, Post, Comment)
│   │
│   ├── repository/         # DB layer (CRUDs)
│   │   ├── user_repo.go
│   │   ├── post_repo.go
│   │   └── comment_repo.go
│   │
│   ├── service/            # Business logic (e.g., AuthService)
│   │   ├── auth_service.go
│   │   ├── user_service.go
│   │   └── post_service.go
│   │
│   └── router/             # Router setup (e.g., gorilla/mux, chi)
│       └── routes.go
│
├── migrations/             # Optional: alias to internal/db/migrations for clarity
│
├── .env                    # Environment variables
```
├── Dockerfile              # Docker build config
├── docker-compose.yml      # Docker services (e.g., app + PostgreSQL)
├── go.mod                  # Go module file
└── README.md
