# boilerplate_go_websocket

## Overview
This project is a Go-based WebSocket server boilerplate designed for real-time chat applications. It includes:
- **Authentication** with JWT.
- **Room management**: Create, fetch, and close chat rooms.
- **Real-time chat messaging** using WebSockets and a hub/manager pattern.
- **ORM generation** using GORM Gen.

## Architecture
- **main.go**  
  Bootstraps the application, loads environment variables, initializes the database, sets up Echo with middleware and routes, and starts the HTTP server.

- **cmd/**  
  Contains tools like `gentool.go` to generate ORM code.

- **config/**  
  Contains the database schema (`schema.sql`).

- **internal/**  
  - **controller/**: Handles HTTP and WebSocket requests (e.g. [`auth-controller.go`](d:/git/boilerplate_go_websocket/internal/controller/auth-controller.go), [`chat-controller.go`](d:/git/boilerplate_go_websocket/internal/controller/chat-controller.go), [`room-controller.go`](d:/git/boilerplate_go_websocket/internal/controller/room-controller.go)).
  - **core/**: Contains WebSocket logic including the Hub, Client and connection Upgrader (see [`hub.go`](d:/git/boilerplate_go_websocket/internal/core/hub.go), [`client.go`](d:/git/boilerplate_go_websocket/internal/core/client.go)).
  - **database/**: Database initialization (see [`dbc.go`](d:/git/boilerplate_go_websocket/internal/database/dbc.go)).
  - **dto/**: Data Transfer Objects for API responses (e.g. [`room-dto.go`](d:/git/boilerplate_go_websocket/internal/dto/room-dto.go), [`user-dto.go`](d:/git/boilerplate_go_websocket/internal/dto/user-dto.go)).
  - **gorm_gen/**: Automatically generated ORM code using GORM Gen.
  - **middleware/**: JWT authentication middleware (see [`jwt-middleware.go`](d:/git/boilerplate_go_websocket/internal/middleware/jwt-middleware.go)).
  - **model/**: GORM models mapping to database tables.
  - **router/**: Route definitions binding endpoints to controllers (e.g. [`auth-router.go`](d:/git/boilerplate_go_websocket/internal/router/auth-router.go), [`chat-router.go`](d:/git/boilerplate_go_websocket/internal/router/chat-router.go), [`room-router.go`](d:/git/boilerplate_go_websocket/internal/router/room-router.go)).
  - **usecase/**: Business logic for authentication, room management, and chat functionality.
  - **utils/**: Utility functions for JWT, password hashing, room ID generation, etc.

- **.env.example**  
  Provides a sample configuration for the database and JWT settings.

- **go.mod**  
  Lists project dependencies, including Echo, Logrus, GORM, and Gorilla WebSocket.

## Setup & Installation

1. **Clone the Repository**
   ```sh
   git clone <repo-url>
   cd boilerplate_go_websocket
   ```

2. **Configure Environment**
   - Copy `.env.example` to `.env` and update the configuration:
     ```sh
     cp .env.example .env
     ```
   - Adjust your database credentials and JWT settings as needed.

3. **Database Schema**
   - Run the SQL statements in [config/schema.sql](d:/git/boilerplate_go_websocket/config/schema.sql) against your MySQL database to create the required tables.

4. **Generate ORM Code**
   - Execute the generator tool:
     ```sh
     go run cmd/gentool.go
     ```

5. **Build & Run**
   - Build the project:
     ```sh
     go build
     ```
   - Run the application:
     ```sh
     ./boilerplate_go_websocket
     ```
   - The server listens on port 8080 by default.

## API Endpoints

- **Authentication**
  - `POST /api/auth/login` – Authenticate a user. (See [`auth-controller.go`](d:/git/boilerplate_go_websocket/internal/controller/auth-controller.go))
  
- **Room Management**
  - `POST /api/room` – Create a new chat room.
  - `GET /api/room/:room_id` – Retrieve room details.
  - `PATCH /api/room/:room_id/close` – Close an existing room.

- **Chat WebSocket**
  - `GET /api/ws/:hubID` – Connect to a WebSocket for real-time messaging. (See [`chat-router.go`](d:/git/boilerplate_go_websocket/internal/router/chat-router.go))

## WebSocket Communication
Clients can connect to the `/api/ws/:hubID` endpoint to join a chat room. The connection is managed by a Hub that broadcasts messages to registered clients. JWT authentication is required for protected endpoints.

## Logging & Metrics
- **Logrus**: Used for structured JSON logging.
- **Prometheus**: Integrated for monitoring via Echo middleware.

## Contributing
Contributions are welcome! Fork this repository and submit pull requests for improvements or new features.

## License
Specify your license here if needed.

## Acknowledgements
Thanks to the open-source community for providing excellent libraries like Echo, GORM, Gorilla WebSocket, and others that make this project possible.