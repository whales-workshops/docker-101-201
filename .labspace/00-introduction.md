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

