# Testing with Testcontainers

This document explains how the integration tests work using Testcontainers to test Redis interactions.

## Overview

The tests use [Testcontainers](https://golang.testcontainers.org/) to spin up real Redis containers during testing. This ensures tests run against actual Redis instances rather than mocks, providing more reliable integration testing.

## Test Architecture

```mermaid
graph LR
    A[Go Test] --> B[Testcontainers]
    B --> C[Docker Engine]
    C --> D[Redis Container]
    A --> D
    D --> E[Test Data]
```

## Test 1: TestGetRestaurant

This test validates retrieving a single restaurant by ID from Redis.

### Flow

```mermaid
sequenceDiagram
    participant Test
    participant Testcontainers
    participant Docker
    participant Redis

    Test->>Testcontainers: Start Redis container
    Testcontainers->>Docker: Create redis:7.2.4 container
    Docker->>Testcontainers: Container ready
    Testcontainers->>Test: Return endpoint

    Test->>Redis: Insert Restaurant A (ID: 1)
    Test->>Redis: Insert Restaurant B (ID: 2)

    Test->>Redis: getRestaurant("2")
    Redis->>Test: Return Restaurant B

    Test->>Test: Assert Name == "Restaurant B"
```

### Steps

1. **Container Setup**: Testcontainers creates a Redis 7.2.4 container
2. **Data Population**: Two restaurants are inserted into Redis using `HSet`
   - Restaurant A (ID: 1)
   - Restaurant B (ID: 2)
3. **Retrieval**: The `getRestaurant()` function fetches restaurant with ID "2"
4. **Validation**: Assert the returned restaurant name is "Restaurant B"

### Redis Commands Used

```
HSET restaurant:1 name "Restaurant A" address "123 Main St" ...
HSET restaurant:2 name "Restaurant B" address "456 Elm St" ...
HGETALL restaurant:2
```

## Test 2: TestGetAllRestaurants

This test validates retrieving all restaurants from Redis.

### Flow

```mermaid
sequenceDiagram
    participant Test
    participant Testcontainers
    participant Docker
    participant Redis

    Test->>Testcontainers: Start Redis container
    Testcontainers->>Docker: Create redis:7.2.4 container
    Docker->>Testcontainers: Container ready
    Testcontainers->>Test: Return endpoint

    Test->>Redis: Insert Restaurant A (ID: 1)
    Test->>Redis: SADD restaurants "1"
    Test->>Redis: Insert Restaurant B (ID: 2)
    Test->>Redis: SADD restaurants "2"

    Test->>Redis: getAllRestaurants()
    Redis->>Test: Return all restaurant IDs
    Test->>Redis: Get details for each ID
    Redis->>Test: Return restaurant details

    Test->>Test: Assert length == 2
    Test->>Testcontainers: Terminate container
    Testcontainers->>Docker: Stop & remove container
```

### Steps

1. **Container Setup**: Testcontainers creates a Redis 7.2.4 container
2. **Data Population**: Two restaurants are inserted with an additional step:
   - Insert restaurant data using `HSet`
   - Add restaurant ID to a set using `SAdd`
3. **Retrieval**: The `getAllRestaurants()` function:
   - Fetches all IDs from the "restaurants" set
   - Retrieves details for each ID
4. **Validation**: Assert exactly 2 restaurants are returned
5. **Cleanup**: Container is explicitly terminated

### Redis Commands Used

```
HSET restaurant:1 name "Restaurant A" address "123 Main St" ...
SADD restaurants "1"
HSET restaurant:2 name "Restaurant B" address "456 Elm St" ...
SADD restaurants "2"
SMEMBERS restaurants
HGETALL restaurant:1
HGETALL restaurant:2
```

## Testcontainers Lifecycle

```mermaid
stateDiagram-v2
    [*] --> Creating: RunContainer()
    Creating --> Starting: Container Created
    Starting --> Waiting: Container Started
    Waiting --> Ready: Health Check Passed
    Ready --> Testing: Run Tests
    Testing --> Terminating: Tests Complete
    Terminating --> [*]: Container Removed
```

### Key Components

- **Ryuk Container**: Testcontainers automatically starts a Ryuk container (testcontainers/ryuk:0.11.0) for cleanup
- **Wait Strategy**: Tests wait for the log message "Ready to accept connections" before proceeding
- **Auto-cleanup**: Containers are automatically removed after tests (except TestGetRestaurant which doesn't explicitly terminate)

## Running the Tests

```bash
# Run all tests
go test -v

# Run specific test
go test -v -run TestGetRestaurant
go test -v -run TestGetAllRestaurants

# Run tests in parallel
go test -v -p 1
```

## Test Output Explained

The test logs show:
- Docker connection details (version, API, OS)
- Container creation and startup events
- Wait strategy progress
- Test results with emoji indicators (🍱 for single restaurant, 🥘 for multiple)
- Cleanup events (terminate/stop)

## Benefits of Testcontainers

1. **Real Dependencies**: Tests run against actual Redis, not mocks
2. **Isolation**: Each test gets a fresh Redis instance
3. **Portability**: Tests work anywhere Docker is available
4. **Automatic Cleanup**: Containers are removed after tests complete
5. **Reproducibility**: Consistent behavior across environments
