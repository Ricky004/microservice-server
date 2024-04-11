package model

import (
	"time"

	"github.com/google/uuid"
)

// Defines a Go struct for managing orders with order ID, customer ID, line items, and timestamps.
type Order struct {
	OrderID     uint64
	CustomerID  uuid.UUID
	LineItem    []LineItem
	CreatedAt   *time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
}



type LineItem struct {
	ItemID   uuid.UUID
	Quantity uint
	Price    uint
}
