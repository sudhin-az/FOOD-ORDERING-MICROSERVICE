package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID              string         `gorm:"primaryKey"`
	UserID          string         `gorm:"not null"`
	ItemIDs         pq.StringArray `gorm:"type:text[]"`
	Quantities      []int          `json:"quantities"`
	Status          string         `gorm:"default:'pending'"`
	DeliveryAddress string         `gorm:"type:text"`
	TotalAmount     float64        `gorm:"type:decimal(10,2)"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
}

type OrderRepository struct {
	conn *gorm.DB
}

func NewOrderRepository(dbURL string) (*OrderRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}
	if err := conn.AutoMigrate(&Order{}); err != nil {
		return nil, fmt.Errorf("failed to migrate order table: %v", err)
	}
	return &OrderRepository{conn: conn}, nil
}

func (r *OrderRepository) PlaceOrder(ctx context.Context, userID string, itemIDs []string, quantities []int, deliveryAddress string, totalAmount float64) (Order, error) {
	orderID := uuid.NewString()

	
	newOrder := Order{
		ID:              orderID,
		UserID:          userID,
		ItemIDs:         itemIDs,
		Quantities:      quantities,
		Status:          "pending",
		DeliveryAddress: deliveryAddress,
		TotalAmount:     totalAmount,
	}

	err := r.conn.WithContext(ctx).Create(&newOrder).Error
	if err != nil {
		return Order{}, fmt.Errorf("failed to place order: %v", err)
	}

	log.Printf("Order Placed Successfully: %+v", newOrder)
	return newOrder, nil
}

func (r *OrderRepository) GetOrderByUser(ctx context.Context, userID string) ([]Order, error) {
	var order []Order

	err := r.conn.WithContext(ctx).Where("user_id = ?", userID).Find(&order).Error
	if err != nil {
		return []Order{}, fmt.Errorf("failed to get orders for user %s: %v", userID, err)
	}
	log.Printf("Orders Retrieved Successfully for user %s: %d orders", userID, len(order))
	return order, nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID string) (Order, error) {
	var order Order

	err := r.conn.WithContext(ctx).Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return Order{}, fmt.Errorf("failed to get orders %s: %v", orderID, err)
	}
	return order, nil
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, orderID, status string) error {
	err := r.conn.WithContext(ctx).Model(&Order{}).Where("id = ?", orderID).Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update order status: %v", err)
	}

	log.Printf("Order %s status updated to %s", orderID, status)
	return nil
}

func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	err := r.conn.WithContext(ctx).Find(&orders).Error
	if err != nil {
		return []Order{}, fmt.Errorf("failed to get all orders: %v", err)
	}
	log.Printf("All Orders Retrieved Successfully: %d orders", len(orders))
	return orders, nil
}
