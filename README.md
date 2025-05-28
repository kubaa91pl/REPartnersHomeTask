# RE Partners Shipment Optimization

A Golang-based backend service and Vue.js (Quasar) frontend application for calculating optimal shipment packaging.

## 📦 Problem

Given a number of items ordered and available pack sizes, determine:
1. **Only full packs** are allowed.
2. **Minimize total items sent** (avoid overpacking).
3. **Minimize number of packs** (secondary priority).

Example:
- 501 items → Optimal: `1x500 + 1x250` (least extra items and fewer packs)

---

## 🧪 Features

- Dynamic programming-based calculation
- Extensible pack size configuration
- Possibility of using default pack sizes
- Memory-based persistence
- Full HTTP API (Go `net/http`)
- Responsive frontend using Quasar
- Unit tests (backend + frontend)
- Edge-case: large amount of items `e.g. 500000 items` with low capacity packs `e.g. {23, 31, 53}` handled efficiently
---

## 🧰 Technologies

- **Backend:** Go, net/http, UUID, dynamic programming
- **Frontend:** Quasar (Vue 3, Composition API)
- **Containerization:** Docker-ready
- **Persistence:** In-memory with UUID-based retrieval

---

## Running Unit Tests

- Backend uses testing and testify packages.
- Frontend uses vitest for testing.

```bash
# Backend
cd backend
go test ./... -v

# Frontend
cd frontend/quasar-project
npm test
``` 

## 🚀 Running Application Locally

- Server runs at `http://localhost:8080`
- UI runs at `http://localhost:9000`

```bash
# Backend
cd backend/cmd/server
go run . serve

# Frontend
cd frontend/quasar-project
quasar dev
```

## 🐳 Running Application with Docker Compose

To build and start both the backend and frontend using Docker Compose, 
execute following from the main project directory:
```bash
docker-compose up --build
```

This will:
- Start the backend server at `http://localhost:8080`
- Start the frontend UI at `http://localhost:9000`

### 🔁 Hot Reloading

Thanks to mounted volumes, any changes you make in local frontend/quasar-project or backend will be reflected without restarting the containers.

### 🔎 Verifying in Browser

After running the above command:
Navigate to `http://localhost:9000` in your browser to use the application UI
The backend API will be accessible at `http://localhost:8080`

## 🚀 Live Demo

- Frontend (UI): `https://repartnershometask-ui.onrender.com`
- Backend API: `https://repartnershometask.onrender.com/shipment` (POST endpoint)