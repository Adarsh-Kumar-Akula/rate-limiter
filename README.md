# Rate-Limited API Service (Golang)

A concurrency-safe, in-memory rate-limited API built in Go using a clean layered architecture.  
Supports per-user request limiting with accurate enforcement under parallel load.

---

## Live Demo

Base URL: [https://rate-limiter-vcbt.onrender.com](https://rate-limiter-vcbt.onrender.com)

---

## Features

- Rate limit: 5 requests per user per minute
- Concurrency-safe implementation
- Sliding Window Log algorithm (accurate rate limiting)
- Per-user statistics tracking
- Clean architecture (controller → service → store)
- Auto-deployed via Render
- Proper HTTP status handling (429 Too Many Requests)

---

## Architecture

```
controller → service → store
```

### Components

| Layer | Description |
|---|---|
| `controller/` | Handles HTTP requests and responses |
| `service/` | Core rate limiting logic |
| `store/` | In-memory data storage |
| `model/` | Data structures |

---

## API Endpoints

### `POST /request`

Accepts a request and enforces rate limiting.

#### Request Body

```json
{
  "user_id": "user1",
  "payload": "any-data"
}
```

#### Responses

| Status | Meaning |
|---|---|
| `200 OK` | Request accepted |
| `429 Too Many Requests` | Rate limit exceeded |

---

### `GET /stats`

Returns request statistics for a user.

#### Query Params

```
/stats?user_id=user1
```

#### Response

```json
{
  "user_id": "user1",
  "total_requests": 10,
  "allowed": 5,
  "blocked": 5
}
```

---

## Rate Limiting Design

### Algorithm Used: Sliding Window Log

- Stores timestamps of recent requests per user
- Removes timestamps older than 60 seconds
- Allows request only if count < 5

### Why This Approach

- Precise enforcement without burst loopholes
- Simple and easy to reason about
- Suitable for moderate traffic systems

---

## Concurrency Handling

- Per-user mutex locks ensure correctness under parallel requests
- Global `RWMutex` used only for safe user map access

This prevents:

- Race conditions
- Incorrect request counts
- Over-accepting requests under load

---

## Running Locally

### Prerequisites

- Go 1.20+

### Steps

```bash
git clone https://github.com/your-username/rate-limiter.git
cd rate-limiter

go mod tidy
go run main.go
```

Server runs on:

```
http://localhost:8080
```

---

## Testing the API

### Examples

```bash
curl -X POST http://localhost:8080/request \
  -H "Content-Type: application/json" \
  -d '{"user_id":"user1","payload":"test"}'
```

```bash
curl "http://localhost:8080/stats?user_id=user1"
```

---

## Deployment

Deployed using [Render](https://render.com) (free tier).

### Deployment Details

- Auto-deploy on every push to `main` branch
- Uses `PORT` environment variable
- Instance may sleep after inactivity (cold start possible)

---

## Limitations

- **In-memory storage:**
  - Data resets on server restart or redeploy
  - Not suitable for distributed systems
- No persistence layer
- No horizontal scaling support
- Time Dependency
  -  Relies on system clock (time.Now())
  -  Clock drift or skew can affect accuracy

---

## Future Improvements

- Replace in-memory store with **Redis** for distributed rate limiting
- Implement **Token Bucket** algorithm for burst handling
- Add middleware-based rate limiting
- Add monitoring and logging
- Add unit and load tests
- Containerize and deploy with orchestration

---

## Design Decisions

| Decision | Reason |
|---|---|
| Sliding Window Log | Strict accuracy over simplicity |
| Per-user locking | Better concurrency vs. global locking |
| Service layer rate limiter | Separation of concerns |
| Dependency injection | Flexibility and testability |

---

## Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Web Framework
- **Storage:** In-memory data structures
- **Deployment:** Render
