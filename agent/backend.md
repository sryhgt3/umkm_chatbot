# BACKEND.md

Production-Ready Backend Guide (Go + FastAPI + PostgreSQL)

---

## 🎯 PURPOSE

Build a **scalable, maintainable, and reusable backend** for:

* Chat-based ordering system (Telegram / WhatsApp)
* UMKM dashboard (menu & order management)
* AI-assisted chat parsing

---

## 🧠 CORE PRINCIPLES

* **Clean Architecture**
* **Separation of Concerns**
* **Reusable Code**
* **Scalable Design**
* **Minimal AI Dependency (use only where needed)**

---

## 🧱 SYSTEM ARCHITECTURE

```text
[ Client (Telegram / React) ]
            ↓
        [ Go Backend ]
            ↓
      [ FastAPI (AI Service) ]
            ↓
       [ PostgreSQL ]
```

---

## 🟢 GO BACKEND (CORE SYSTEM)

### 🎯 RESPONSIBILITIES

* Handle HTTP API
* Handle chatbot webhook
* Business logic (order, menu, customer)
* Validation & data processing
* Communicate with AI service
* Manage database operations

---

## 📁 PROJECT STRUCTURE (CLEAN ARCHITECTURE)

```bash
/internal/
  /config        → env & config loader
  /database      → DB connection
  /model         → database entities
  /repository    → database queries
  /service       → business logic
  /handler       → HTTP / webhook
  /dto           → request & response structs
  /middleware    → auth, logging
  /utils         → reusable helpers
/cmd/server/main.go
```

---

## 🧩 LAYER RESPONSIBILITIES

### 🔹 Handler Layer

* Accept request (HTTP / webhook)
* Validate input
* Call service
* Return response

```go
func CreateOrder(c *gin.Context) {
    var req dto.CreateOrderRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result, err := orderService.CreateOrder(req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, result)
}
```

❌ No business logic here

---

### 🔹 Service Layer (CORE LOGIC)

```go
func (s *OrderService) CreateOrder(req dto.CreateOrderRequest) (*model.Order, error) {
    // validate items
    // calculate total
    // call repository
}
```

✔ All business logic here

---

### 🔹 Repository Layer

```go
func (r *OrderRepository) Create(order *model.Order) error {
    return r.db.Create(order).Error
}
```

✔ Only DB access

---

## 🧩 DTO (DATA TRANSFER OBJECT)

```go
type CreateOrderRequest struct {
    CustomerID int64              `json:"customer_id"`
    Items      []OrderItemRequest `json:"items"`
}
```

---

## 🔌 DEPENDENCY INJECTION

Initialize dependencies in `main.go`:

```go
orderRepo := repository.NewOrderRepository(db)
orderService := service.NewOrderService(orderRepo)
orderHandler := handler.NewOrderHandler(orderService)
```

---

## 🔄 DATABASE TRANSACTION

```go
tx := db.Begin()

// insert order
// insert order_items

if err != nil {
    tx.Rollback()
} else {
    tx.Commit()
}
```

---

## 📊 LOGGING

Log:

* incoming requests
* errors
* order creation

---

## ⚡ PERFORMANCE

* Use indexing (store_id, chat_id)
* Avoid unnecessary queries
* Implement pagination

---

## 🔐 SECURITY

* Validate all inputs
* Use bcrypt for password hashing
* Protect endpoints (JWT in future)
* Rate limit webhook

---

## 🟡 FASTAPI (AI SERVICE)

### 🎯 PURPOSE

* Parse user chat
* Convert text → structured data
* Handle NLP

---

### 📁 STRUCTURE

```bash
/app
  /routes
  /services
  /models
main.py
```

---

### 🔌 ENDPOINT

#### POST /parse-order

Request:

```json
{
  "message": "nasi goreng 2, es teh 1"
}
```

Response:

```json
{
  "items": [
    {"name": "nasi goreng", "qty": 2},
    {"name": "es teh", "qty": 1}
  ]
}
```

---

## ⚠️ AI RULES

* ✅ Use only for parsing & NLP
* ❌ No database operations
* ❌ No business logic

---

## 🔄 GO ↔ FASTAPI FLOW

1. User sends message
2. Go receives webhook
3. Go checks intent
4. If needed → send to FastAPI
5. FastAPI returns structured data
6. Go processes order

---

## 🧠 BUSINESS FLOW (ORDER)

1. Receive message
2. Detect order
3. Parse items (AI or rule-based)
4. Validate menu
5. Calculate total
6. Save order
7. Send confirmation

---

## 🚀 DEVELOPMENT PHASES

### Phase 1 (MVP)

* Menu API
* Order API
* Telegram webhook
* Manual parsing

---

### Phase 2

* FastAPI integration
* Auto parsing

---

### Phase 3

* Scaling (Redis, queue)
* Multi-store optimization

---

## 💣 ANTI-PATTERNS (AVOID)

* ❌ Business logic in handler
* ❌ Direct DB query in handler
* ❌ AI for everything
* ❌ Hardcoded config
* ❌ Duplicate code

---

## 🧠 BEST PRACTICES

* Use interfaces for repository
* Keep functions small & focused
* Use environment variables
* Write modular code
* Make everything reusable

---

## 💬 FINAL INSIGHT

A good backend is:

* structured
* predictable
* easy to scale

Not:

* messy
* over-engineered
* AI-dependent

---

## 🚀 FINAL STACK

* Go → core backend
* FastAPI → AI parsing
* PostgreSQL → database

👉 Focus on building a **chat-driven ordering system**, not just a chatbot
