# Docker 101

## Docker Compose is useful for multi-container applications

### I need microservices, monitoring, and a reverse proxy

![stack](imgs/stack.png)

### Architecture diagram

```mermaid
graph TD
    User[User] -->|HTTP Request on port 8888| Nginx;

    subgraph "Application Services"
        Nginx -- "route /" --> HomePage[service-home-page:6000];
        Nginx -- "route /hello" --> HelloPage[service-hello:6000];
        Nginx -- "route /hey" --> HeyPage[service-hey:6000];
    end

    subgraph "Monitoring"
        Prometheus -- "Scrapes /metrics" --> HomePage;
        Grafana -- "Data Source" --> Prometheus;
        User --> |Port 4000| Grafana;
        User --> |Port 9090| Prometheus;
    end

    style Nginx fill:#f9f,stroke:#333,stroke-width:2px
    style HomePage fill:#ccf,stroke:#333,stroke-width:2px
    style HelloPage fill:#ccf,stroke:#333,stroke-width:2px
    style HeyPage fill:#ccf,stroke:#333,stroke-width:2px
    style Prometheus fill:#cfc,stroke:#333,stroke-width:2px
    style Grafana fill:#cfc,stroke:#333,stroke-width:2px
```


```bash
docker compose up
```

#### Open the Web application

- home: http://localhost:8888
- hello: http://localhost:8888/hello
- hey: http://localhost:8888/hey

✋ **The links to Prometheus and Grafana are on the home page**

**Dashboard**: http://localhost:4000/d/O4gR7H5Iz/k33g_dashboard?orgId=1&refresh=5s

#### Open the Web application pages with different ports

- home: http://localhost:6500
- hello: http://localhost:6100
- hey: http://localhost:6200