package model

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Role       string    `json:"role"`
	StoreID    *int      `json:"store_id"`
	TelegramID *string   `json:"telegram_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

const (
	RoleSuperAdmin = "SUPER_ADMIN"
	RoleStoreOwner = "STORE_OWNER"
	RoleCustomer   = "CUSTOMER"
)
