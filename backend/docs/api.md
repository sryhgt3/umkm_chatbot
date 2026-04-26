# UMKM Chatbot API Documentation

This document describes the available API endpoints for the UMKM Chatbot backend.

## Base URL
Default local development URL: `http://localhost:8080`

---

## Public Endpoints

### 1. Health Check
Checks if the server is running.
- **URL**: `/health`
- **Method**: `GET`
- **Success Response**:
    - **Code**: 200 OK
    - **Content**: `{"status": "healthy"}`

### 2. Telegram Webhook
Endpoint for receiving updates from the Telegram Bot API.
- **URL**: `/webhook/telegram`
- **Method**: `POST`
- **Payload**: Standard Telegram Update JSON.
- **Success Response**:
    - **Code**: 200 OK
    - **Content**: `{"status": "ok"}`

### 3. User Registration
Register a new user (Store Owner or Admin).
- **URL**: `/auth/register`
- **Method**: `POST`
- **Payload**:
    ```json
    {
      "name": "User Name",
      "email": "user@example.com",
      "password": "strongpassword",
      "role": "store_owner"
    }
    ```
- **Success Response**:
    - **Code**: 201 Created
    - **Content**: `{"message": "user registered successfully"}`

### 4. User Login
Authenticate and receive a JWT token.
- **URL**: `/auth/login`
- **Method**: `POST`
- **Payload**:
    ```json
    {
      "email": "user@example.com",
      "password": "strongpassword"
    }
    ```
- **Success Response**:
    - **Code**: 200 OK
    - **Content**: `{"token": "JWT_TOKEN_HERE"}`

---

## Protected Endpoints (Requires JWT)
All endpoints below require the `Authorization: Bearer <TOKEN>` header.

### 1. Get Current User info
- **URL**: `/api/auth/me`
- **Method**: `GET`
- **Success Response**:
    - **Code**: 200 OK
    - **Content**: User object details.

### 2. Super Admin Dashboard
- **URL**: `/api/admin/dashboard`
- **Method**: `GET`
- **Requirement**: Must have `super_admin` role.
- **Success Response**:
    - **Code**: 200 OK
    - **Content**: `{"message": "Welcome to Super Admin Dashboard"}`

### 3. Store Owner Dashboard
- **URL**: `/api/store/dashboard`
- **Method**: `GET`
- **Requirement**: Must have `store_owner` role.
- **Success Response**:
    - **Code**: 200 OK
    - **Content**: `{"message": "Welcome to Store Owner Dashboard"}`
