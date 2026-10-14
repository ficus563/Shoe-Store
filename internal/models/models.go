package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Gender      string    `json:"gender"`
	Size        string    `json:"size"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	Season      string    `json:"season"`
	Price       float64   `json:"price"`
}

type Photo struct {
	ID        uuid.UUID `json:"id"`
	Filename  string    `json:"filename"`
	ProductID uuid.UUID `json:"product_id"`
}

type Role struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type User struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`
	RoleID uuid.UUID `json:"role_id"`
}

type Cart struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
	Count     int       `json:"count"`
}

type Order struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Status string    `json:"status"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Price     float64   `json:"price"`
}
