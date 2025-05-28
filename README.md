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
- Memory-based persistence
- Full HTTP API (Go `net/http`)
- Responsive frontend using Quasar
- Unit tests (backend + frontend)
- Edge-case: `500000 items` with `{23, 31, 53}` handled efficiently

---

## 🧰 Technologies

- **Backend:** Go, net/http, UUID, dynamic programming
- **Frontend:** Quasar (Vue 3, Composition API)
- **Containerization:** Docker-ready
- **Persistence:** In-memory with UUID-based retrieval

---

## 🚀 Running Locally
Server runs at http://localhost:8080
Frontend runs at http://localhost:9000

```bash
# Backend
cd backend/cmd/server
go run . serve

# Frontend
cd frontend/quasar-project
quasar dev