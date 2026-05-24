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
├── pollen/
│   ├── backend/          # Go Lambda (Echo v5)
│   │   ├── cmd/lambda/   # Lambda entrypoint
│   │   ├── internal/     # Handlers, domain logic, DB
│   │   ├── migrations/   # goose migration SQL
│   │   ├── pkg/          # Shared packages (config, errors)
│   │   └── openapi.yaml  # OpenAPI 3.0 spec
│   │
│   ├── mobile/           # React Native (Expo)
│   │   ├── assets/       # Icons, splash images
│   │   └── src/          # Feature-based source code
│   │
│   └── infra/            # AWS CDK (TypeScript)
│
├── .github/
│   └── workflows/        # CI/CD workflows
│
└── docs/
    ├── README.md
    └── README.ja.md
```

## Getting Started

### Prerequisites

- Go 1.26.3 (managed via [asdf](https://asdf-vm.com/))
- Node.js 24.13.0 (managed via asdf)
- AWS CDK CLI (`npm install -g aws-cdk`)
- Expo CLI (`npx expo`)
- Docker (for local development)

### Backend (local)

```bash
cd pollen/backend
docker compose up
```

### Frontend (local)

```bash
cd pollen/mobile
npx expo start
```

### Infrastructure

```bash
cd pollen/infra
cdk bootstrap   # first time only
cdk diff        # preview changes
cdk deploy      # apply changes
```

## API Endpoints

| Method | Path               | Description                        |
| ------ | ------------------ | ---------------------------------- |
| GET    | `/pollen`          | Fetch pollen levels by coordinates |
| POST   | `/symptoms`        | Record daily symptoms              |
| GET    | `/symptoms`        | List recorded symptoms             |
| PUT    | `/symptoms/{date}` | Update symptoms for a date         |
