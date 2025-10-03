# Chat App (Go + React Native)

🚀 A real-time chat application built with **Golang (backend)** and **React Native (frontend)**.  
This project is designed for learning purposes and is open source.

---

## Features (Planned / Implemented)

- ⏳ User authentication (JWT + REST)
- ⏳ Real-time messaging (WebSockets)
- ⏳ Typing indicators
- ⏳ Online / offline presence
- ⏳ Message read receipts
- ⏳ Group and direct chats
- ⏳ Reactions / polls (future)
- ⏳ Push notifications (future)

---

## Tech Stack

### Backend

- **Go (Golang)**
- **PostgreSQL** – store users, chats, messages
- **Redis** – caching + pub/sub (future scaling)
- **WebSockets** – real-time messaging
- **Gin/Fiber** – REST endpoints for auth + history

### Frontend

- **React Native (Expo)**
- **Redux Toolkit** for state management
- **WebSocket client** for live messages
- **AsyncStorage** for session persistence

---

## Architecture

- **HTTP (REST)** → login, registration, fetching old messages
- **WebSocket** → new messages, typing indicators, presence

```text
+-------------+        HTTP (auth, history)       +-------------+
|  React      | <--------------------------------> |   Go API    |
|  Native App |                                    |  (Gin/Fiber)|
+-------------+                                    +-------------+
       |                                                   |
       |                 WebSocket (real-time)             |
       +---------------------------------------------------+
```
