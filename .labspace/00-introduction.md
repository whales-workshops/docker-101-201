<style>
        h1 {
            color: blue;
            font-family: Arial, sans-serif; /* Optionnel : pour une police plus moderne */
            font-size: 40px;
        }
        h2 {
            color: red;
            font-family: Arial, sans-serif; /* Optionnel : pour une police plus moderne */
        }        
</style>

# Docker 101-201 Workshop

This repository contains three progressive Docker Compose demonstrations, from basic web applications to complex microservices architectures with monitoring and testing.

## 01 - Docker Compose Web App

**Purpose**: Introduction to Docker Compose for development workflows

A simple Node.js web application demonstrating basic Docker Compose usage with:
- **Frontend**: Node.js webapp (port 7070)
- **Backend**: Node.js API service (port 6060)
- **Database**: Redis server (port 6379)

**Key Features**:
- Multi-service orchestration with Docker Compose
- Network isolation (frontend/backend networks)
- Volume persistence for Redis data
- **Watch mode** for live development: automatic sync of file changes without rebuilding

**Architecture**:
```
User → WebApp (7070) → API (6060) → Redis (6379)
```

**Usage**:
```bash
cd 01-docker-compose-web-app
docker compose up --build        # Initial start
docker compose up --watch        # Development mode with live reload
```

---

## 02 - Docker Compose Complex Architecture

**Purpose**: Multi-container microservices with reverse proxy and monitoring

A complete microservices stack demonstrating production-like architecture with:
- **Reverse Proxy**: Nginx routing traffic to services
- **Microservices**: Three Node.js services (home, hello, hey)
- **Monitoring**: Prometheus + Grafana dashboard
- **Metrics**: Prometheus client for application metrics

**Key Features**:
- Route-based load balancing with Nginx
- Centralized monitoring with Prometheus scraping
- Pre-configured Grafana dashboards
- Service health metrics and visualization
- Multiple access methods (via proxy or direct ports)

**Architecture**:
```
User (8888) → Nginx → Services (home/hello/hey)
                      ↓
                 Prometheus (9090) → Grafana (4000)
```

**Services**:
- Home: `http://localhost:8888/` or `:6500`
- Hello: `http://localhost:8888/hello` or `:6100`
- Hey: `http://localhost:8888/hey` or `:6200`
- Grafana: `http://localhost:4000`
- Prometheus: `http://localhost:9090`

**Usage**:
```bash
cd 02-docker-compose-complex-architecture
docker compose up
```

---

## 03 - Go Web App Demo | Testcontainers

**Purpose**: Modern Go application with Redis integration and comprehensive testing

A production-ready Go web application showcasing:
- **Application**: Go HTTP server with Redis backend
- **Database**: Redis for restaurant data storage
- **Multi-stage Build**: Optimized Dockerfile using `scratch` base image
- **Testing**: Integration tests with Testcontainers
- **AI Testing**: LLM integration testing with Docker Model Runner

**Key Features**:
- Static binary compilation (`CGO_ENABLED=0`) for minimal container size
- Restaurant CRUD API with Redis Hash and Set operations
- **Testcontainers**: Spin up real Redis instances for integration tests
- **Docker Model Runner**: Local LLM testing with Qwen 2.5 model
- Watch mode for development with file sync and auto-rebuild
- Environment-based configuration

**Architecture**:
```
WebApp (8080) → Redis (6379)
Test Suite → Testcontainers → Redis Container
Test Suite → Docker Model Runner → Qwen 2.5 LLM
```

**API Endpoints**:
- `/api/variables` - Application configuration
- `/api/restaurants` - List all restaurants
- `/api/info` - Application info

**Testing Documentation**:
- [TESTING.md](03-go-web-app-demo/TESTING.md) - Redis integration tests with Testcontainers
- [TESTING-LLM.md](03-go-web-app-demo/TESTING-LLM.md) - LLM testing with Docker Model Runner

**Usage**:
```bash
cd 03-go-web-app-demo
docker compose up --build        # Run application
go test -v                       # Run all tests
go test -v -run TestGetRestaurant          # Redis integration test
go test -v -run TestHawaiianPizzaExpert    # LLM test
```

---

## Workshop Progression

1. **Demo 01**: Learn Docker Compose basics and development workflows
2. **Demo 02**: Scale to microservices with monitoring and reverse proxy
3. **Demo 03**: Implement production-ready patterns with comprehensive testing

## Prerequisites

- Docker Desktop (with AI features for demo 03 LLM tests)
- Docker Compose V2
- Go 1.25+ (for demo 03)
- Node.js 18+ (for demos 01-02)

## Common Commands

```bash
# Start services
docker compose up

# Start with build
docker compose up --build

# Start in watch mode (development)
docker compose up --watch

# Stop services
docker compose down

# View logs
docker compose logs -f
```
