package mallagent

import (
	"time"
)

// Product 商品信息
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
	CreatedAt   int64   `json:"created_at"`
}

// CartItem 购物车项
type CartItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}

// Cart 购物车
type Cart struct {
	UserID    string     `json:"user_id"`
	Items     []CartItem `json:"items"`
	Total     float64    `json:"total"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// OrderItem 订单项
type OrderItem struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}

// Order 订单
type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"` // pending, paid, shipped, delivered, cancelled
	Address   string      `json:"address"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// User 用户信息
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserMessage 用户消息
type UserMessage struct {
	UserID  string            `json:"user_id"`
	Query   string            `json:"query"`
	Context map[string]string `json:"context"`
}

// 响应结果
type ResponseResult struct {
	Message string
	Actions []AgentAction
}

// 客服动作
type AgentAction struct {
	ActionType string
	TargetID   string
	Result     string
}

// 会话消息
type ChatMessage struct {
	Role      string
	Content   string
	Timestamp int64
}

// 地址信息
type Address struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Street    string `json:"street"`
	Detail    string `json:"detail"`
	Recipient string `json:"recipient"`
	Phone     string `json:"phone"`
	IsDefault bool   `json:"is_default"`
}
