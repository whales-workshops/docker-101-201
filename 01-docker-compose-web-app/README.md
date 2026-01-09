# Docker 101
> This is a simple demo of using Docker Compose project to run a Node.js app with a Redis database.
## Docker Compose is useful for development

### Initialize the project

Start the project with Docker Compose:
```bash
docker compose up --build
```

Then go to Docker Desktop and:

- Show the running compose stack
- Then **Add data to the Redis database**

```bash
redis-cli

SET hamster 12
SET panda 12
SET tiger 68
SET fox 20
```

### Activate the watch mode

> TODO: explain the watch mode

Stop docker compose (CTRL+C) and restart it with the watch mode:

```bash
docker compose up --watch
```

Change something in `/web/templates/index.ejs`, for example the main title:

```html
<p class="title is-1">
All things 🐳 Compose
</p>
```

Then return to [http://localhost:7070/](http://localhost:7070/) and refresh the page to see the changes.

## Architecture Diagram

```mermaid
graph TD;
    subgraph "User's Browser";
        User[<fa:fa-user> User] -- "HTTP Request (Port 7070)" --> WebApp;
    end;

    subgraph "Docker Compose";
        subgraph "Frontend Network";
            WebApp[<fa:fa-window-maximize> webapp];
        end;

        subgraph "Backend Network";
            API[<fa:fa-server> api];
            Redis[<fa:fa-database> redis-server];
        end;
        
        WebApp -- "API Call (Port 6060)" --> API;
        API -- "Data Access (Port 6379)" --> Redis;
    end;

    style User fill:#f9f,stroke:#333,stroke-width:2px;
    style WebApp fill:#bbf,stroke:#333,stroke-width:2px;
    style API fill:#fb9,stroke:#333,stroke-width:2px;
    style Redis fill:#9f9,stroke:#333,stroke-width:2px;

```
