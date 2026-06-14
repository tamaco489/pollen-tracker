# pollen-tracker

> 日本語版: [README.ja.md](README.ja.md)

## Overview

iOS app to track pollen levels and allergy symptoms.
Records daily symptom intensity and correlates it with real-time pollen data from the Google Pollen API.

## Tech Stack

| Layer     | Technology                                    |
| --------- | --------------------------------------------- |
| Mobile    | React Native (Expo)                           |
| Backend   | Go / Echo v5 / AWS Lambda                     |
| Database  | Turso (libSQL)                                |
| Infra     | AWS CDK (TypeScript) / API Gateway (HTTP API) |
| Auth      | API Key (`x-api-key`)                         |
| CI/CD     | GitHub Actions / EAS Build                    |
| Pollen    | Google Pollen API                             |
| Geocoding | Google Places API                             |

## Directory Structure

```text
pollen-tracker/
├── backend/                    # Go Lambda (Echo v5)
│   ├── cmd/lambda/             # Lambda entrypoint
│   ├── internal/
│   │   ├── di/                 # Dependency injection container
│   │   ├── domain/             # Domain entities & constants (PollenType etc.)
│   │   ├── dto/                # Data transfer objects between layers
│   │   ├── gen/                # oapi-codegen generated — do not edit manually
│   │   ├── handler/            # HTTP handlers (one file per endpoint)
│   │   ├── infrastructure/     # Repository implementations & sqlc generated code
│   │   ├── server/             # Echo server setup & error handler
│   │   └── usecase/            # Business logic
│   ├── oas/                    # OpenAPI spec (split by resource)
│   │   ├── components/         # Schemas, security definitions
│   │   ├── config/             # oapi-codegen config files
│   │   └── paths/              # Path definitions per resource
│   ├── pkg/
│   │   ├── config/             # Environment variable parsing (caarlos0/env)
│   │   ├── errors/             # Sentinel & HTTP error definitions
│   │   ├── infrastructure/     # Turso DB connection
│   │   ├── library/            # External API clients (Google Pollen API etc.)
│   │   ├── logger/             # Structured logger (slog)
│   │   └── utils/              # Utility functions
│   ├── tools/
│   │   ├── httprequest/        # VS Code REST Client test files
│   │   ├── migrations/         # DB migrations (goose SQL files & runner)
│   │   └── sqlc/               # sqlc configuration (sqlc.yaml)
│   └── Makefile
│
├── mobile/                     # React Native (Expo)
│   ├── assets/                 # Icons, splash images
│   └── src/                    # Feature-based source code
│
├── infra/                      # AWS CDK (TypeScript)
│   ├── bin/                    # CDK App entrypoint
│   ├── config/                 # Environment config (EnvConfig)
│   ├── lib/
│   │   ├── stacks/             # Stack definitions
│   │   └── constructs/         # Custom L3 constructs
│   └── test/                   # Jest unit tests
│
├── .github/
│   ├── workflows/              # CI/CD workflows
│   └── PULL_REQUEST_TEMPLATE.md
│
└── docs/
    ├── README.md
    └── README.ja.md
```

## Getting Started

### Prerequisites

- Go 1.26.3 (managed via [asdf](https://asdf-vm.com/))
- Node.js 24.x (managed via asdf)
- AWS CDK CLI (`npm install -g aws-cdk`)
- Expo CLI (`npx expo`)
- Docker (for local development)

### Backend

Install development tools

```bash
cd backend
make setup-tools
```

Start API server

```bash
make up
```

DB migrations

```bash
# Run migrations
make migrate-up

# Create a new migration file
make migrate-create name=<migration_name>
```

Code generation

```bash
# OpenAPI spec → Go (run when oas/ files change)
make gen-api

# SQL queries → Go (run when tools/migrations/sql/ or datastore/queries/ change)
make gen-sqlc

# Validate SQL queries (sqlc-diff + sqlc-vet)
make sqlc-lint
```

### Frontend

```bash
cd mobile
npx expo start
```

### Infrastructure

```bash
cd infra
cdk bootstrap   # first time only
cdk diff        # preview changes
cdk deploy      # apply changes
```
