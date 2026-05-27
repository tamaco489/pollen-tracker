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
├── backend/                  # Go Lambda (Echo v5)
│   ├── cmd/lambda/           # Lambda entrypoint
│   ├── internal/
│   │   └── gen/              # oapi-codegen generated code (do not edit manually)
│   ├── oas/                  # OpenAPI spec (split by resource)
│   │   ├── components/       # Schemas, security definitions
│   │   ├── config/           # oapi-codegen config files
│   │   ├── paths/            # Path definitions per resource
│   │   ├── oapi-base.yaml    # Root spec with $ref links
│   │   └── oapi.yaml         # Bundled spec (auto-generated)
│   ├── tools/
│   │   └── migrations/       # DB migrations (goose)
│   │       ├── cmd/          # Migration runner (go run)
│   │       └── sql/          # goose SQL files
│   └── Makefile
│
├── mobile/                   # React Native (Expo)
│   ├── assets/               # Icons, splash images
│   └── src/                  # Feature-based source code
│
├── infra/                    # AWS CDK (TypeScript)
│
├── .github/
│   ├── workflows/            # CI/CD workflows
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

Run DB migrations

```bash
make migrate-up
```

Create a new migration file

```bash
make migrate-create name=<migration_name>
```

Code generation (OpenAPI spec → Go)

```bash
make gen-api
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
