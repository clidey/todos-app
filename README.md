# Futuristic Todo App

A modern todo application built with Go (Gin + GORM) backend and React + Tailwind CSS frontend.

## Features

- Clean and modern UI with animations
- SQLite database for data persistence
- RESTful API endpoints
- Real-time updates
- Responsive design

## Prerequisites

- Go 1.21 or later
- Node.js 16 or later
- npm or yarn

## Setup

### Backend

1. Navigate to the root directory
2. Install Go dependencies:
   ```bash
   go mod tidy
   ```
3. Run the backend server:
   ```bash
   go run main.go
   ```

The backend server will start on http://localhost:8080

### Frontend

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```

The frontend will be available at http://localhost:5173

## API Endpoints

- GET /api/todos - Get all todos
- POST /api/todos - Create a new todo
- DELETE /api/todos/:id - Delete a todo

## Technologies Used

- Backend:
  - Go
  - Gin Web Framework
  - GORM
  - SQLite

- Frontend:
  - React
  - Tailwind CSS
  - Framer Motion
  - Axios
  - Heroicons