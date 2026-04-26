# FRONTEND.md

Frontend Architecture (React.js) for UMKM Chatbot Dashboard

---

## 🎯 PURPOSE

Build a **clean, scalable, and user-friendly dashboard** for UMKM owners to:

* Manage menus
* View and process orders
* Monitor chatbot activity
* Control store settings

👉 This is **NOT a customer-facing app** — it is an **admin dashboard**

---

## 🧠 CORE PRINCIPLES

* **Simple UI/UX** → designed for non-technical users (UMKM)
* **Reusable Components** → avoid duplication
* **Separation of Concerns**
* **Clean Code Structure**
* **API-driven UI**

---

## 🧱 FRONTEND ARCHITECTURE

```text
[ React UI ]
     ↓
[ Service Layer (API Calls) ]
     ↓
[ Go Backend API ]
```

---

## ⚙️ TECH STACK

* React.js (recommended: Vite)
* Tailwind CSS (for fast styling)
* Axios / Fetch (API calls)
* React Router (routing)
* Zustand / Context API (state management, optional)

---

## 📁 PROJECT STRUCTURE (RECOMMENDED)

```bash
/src
  /components     → reusable UI components
  /pages          → page-level components
  /layouts        → layout (sidebar, navbar)
  /services       → API calls
  /hooks          → custom hooks
  /types          → TypeScript types (if used)
  /utils          → helper functions
  /constants      → static values
```

---

## 🧩 KEY PAGES

### 🏠 Dashboard

* Summary stats (optional)
* Recent orders
* Quick overview

---

### 🍔 Menu Management

* List all menus
* Add menu
* Edit menu
* Delete menu

---

### 🧾 Orders Page

* List orders
* Show details
* Update status:

  * pending
  * confirmed
  * done

---

### ⚙️ Settings (optional)

* Store info
* Bot configuration

---

## 🔌 API INTEGRATION

Centralize API calls inside `/services`

Example:

```js
// services/menuService.js
export const getMenus = async () => {
  const res = await fetch("/menus");
  return res.json();
};
```

---

## 🧠 COMPONENT DESIGN

### ✅ Reusable Components

Examples:

* Button
* Input
* Table
* Modal
* Card

```jsx
// components/Button.jsx
export default function Button({ children, onClick }) {
  return (
    <button onClick={onClick} className="px-4 py-2 bg-green-500 text-white rounded">
      {children}
    </button>
  );
}
```

---

## 🔄 STATE MANAGEMENT

### Use simple approach first:

* `useState`
* `useEffect`

### Upgrade later:

* Zustand (recommended)
* Context API

---

## 🧠 DATA FLOW

```text
Page → Service → API → Response → State → UI
```

---

## ⚡ PERFORMANCE BEST PRACTICES

* Avoid unnecessary re-render
* Use key in list rendering
* Lazy load pages if needed
* Debounce search input

---

## 🎨 UI/UX GUIDELINES

* Clean and minimal
* Mobile-friendly (important for UMKM)
* Use consistent spacing
* Clear call-to-action buttons

---

## 🔐 AUTHENTICATION (FUTURE)

* Login page
* JWT token storage
* Protected routes

---

## 🚀 DEVELOPMENT PHASES

### Phase 1 (MVP)

* Menu CRUD
* Orders list
* Basic UI

---

### Phase 2

* Order detail view
* Status update
* Better UI/UX

---

### Phase 3

* Real-time updates (WebSocket)
* Analytics dashboard

---

## 💣 ANTI-PATTERNS (AVOID)

* ❌ API calls inside components directly (no service layer)
* ❌ Large components (split them)
* ❌ Hardcoded data
* ❌ Inline styles everywhere
* ❌ No loading/error state

---

## 🧠 FINAL INSIGHT

A good frontend is:

* simple to use
* easy to maintain
* consistent

Not:

* overly complex
* full of duplicated code
* hard to navigate

---

## 💬 FINAL NOTE

* Frontend = control panel for UMKM
* Backend = business logic
* Chatbot = sales channel

👉 Focus on usability, not complexity
