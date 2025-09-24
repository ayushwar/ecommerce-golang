package models

import (
	"time"
)

type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Role string `gorm:"type:enum('user','admin');default:'user'" json:"role"`

	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email"` // Unique index for fast email lookups
	Password  string    `gorm:"size:255;not null" json:"-"`                 // Omit password in API responses
	Orders    []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`  // One-to-many relation
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Description string    `gorm:"size:255" json:"description"`
	Price       int64     `gorm:"not null" json:"price"`           // Stored in paise/cents to avoid float issues
	Stock       uint      `gorm:"not null;default:0" json:"stock"` // Stock cannot be negative
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type OrderItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   uint      `gorm:"not null;index" json:"order_id"` // Indexed for faster lookup
	Order     Order     `gorm:"foreignKey:OrderID" json:"-"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product omitempty"` // Relation for easy product fetch
	Quantity  uint      `gorm:"not null;default:1" json:"quantity"`            // Always at least 1
	Price     int64     `gorm:"not null" json:"price"`                         // Snapshotted price (paise/cents)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Order struct {
	ID        uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint        `gorm:"not null;index" json:"user_id"` // Indexed for fast lookups
	User      User        `gorm:"foreignKey:UserID" json:"-"`
	Total     int64       `gorm:"not null" json:"total"` // Stored in paise/cents
	Status    string      `gorm:"size:50;default:'Pending'" json:"status"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"` // One-to-many relation
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
type CartItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  uint      `gorm:"not null;default:1" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
