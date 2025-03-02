package mallagent

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/strings77wzq/gomall/app/mall_agent/pkg/clients"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/cart"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/order"
	"github.com/strings77wzq/gomall/rpc_gen/kitex_gen/product"
)

// ProductServiceImpl 实现ProductService接口
type ProductServiceImpl struct {
	client *clients.RPCClients
}

// CartServiceImpl 实现CartService接口
type CartServiceImpl struct {
	client *clients.RPCClients
}

// OrderServiceImpl 实现OrderService接口
type OrderServiceImpl struct {
	client *clients.RPCClients
}

// 所有服务的包装结构
type ServiceWrappers struct {
	ProductSvc ProductService
	CartSvc    CartService
	OrderSvc   OrderService
}

// 创建所有服务适配器
func NewServiceWrappers(rpcClients *clients.RPCClients) *ServiceWrappers {
	return &ServiceWrappers{
		ProductSvc: &ProductServiceImpl{client: rpcClients},
		CartSvc:    &CartServiceImpl{client: rpcClients},
		OrderSvc:   &OrderServiceImpl{client: rpcClients},
	}
}

// ProductService 实现

// QueryProduct 查询商品
func (p *ProductServiceImpl) QueryProduct(ctx context.Context, name string) (*Product, error) {
	req := &product.SearchProductRequest{
		Keyword: name,
	}

	resp, err := p.client.ProductClient.SearchProduct(ctx, req)
	if err != nil {
		log.Printf("查询商品失败: %v", err)
		return nil, fmt.Errorf("查询商品失败: %v", err)
	}

	if len(resp.Products) == 0 {
		return nil, fmt.Errorf("未找到相关商品")
	}

	// 返回第一个匹配的商品
	item := resp.Products[0]
	return &Product{
		ID:          item.Id,
		Name:        item.Name,
		Description: item.Description,
		Price:       float64(item.Price),
		Stock:       int(item.Stock),
		Category:    item.Category,
		ImageURL:    item.ImageUrl,
	}, nil
}

// GetProduct 获取商品详情
func (p *ProductServiceImpl) GetProduct(ctx context.Context, id string) (*Product, error) {
	req := &product.GetProductRequest{
		Id: id,
	}

	resp, err := p.client.ProductClient.GetProduct(ctx, req)
	if err != nil {
		log.Printf("获取商品详情失败: %v", err)
		return nil, fmt.Errorf("获取商品详情失败: %v", err)
	}

	return &Product{
		ID:          resp.Product.Id,
		Name:        resp.Product.Name,
		Description: resp.Product.Description,
		Price:       float64(resp.Product.Price),
		Stock:       int(resp.Product.Stock),
		Category:    resp.Product.Category,
		ImageURL:    resp.Product.ImageUrl,
	}, nil
}

// CheckStock 检查库存
func (p *ProductServiceImpl) CheckStock(ctx context.Context, id string) (int, error) {
	req := &product.GetProductRequest{
		Id: id,
	}

	resp, err := p.client.ProductClient.GetProduct(ctx, req)
	if err != nil {
		log.Printf("检查库存失败: %v", err)
		return 0, fmt.Errorf("检查库存失败: %v", err)
	}

	return int(resp.Product.Stock), nil
}

// CartService 实现

// AddToCart 添加商品到购物车
func (c *CartServiceImpl) AddToCart(ctx context.Context, userID, productID string, quantity int) (*Cart, error) {
	req := &cart.AddToCartRequest{
		UserId:    userID,
		ProductId: productID,
		Quantity:  int32(quantity),
	}

	resp, err := c.client.CartClient.AddToCart(ctx, req)
	if err != nil {
		log.Printf("添加到购物车失败: %v", err)
		return nil, fmt.Errorf("添加到购物车失败: %v", err)
	}

	return convertCartResponse(resp.Cart), nil
}

// GetCart 获取购物车
func (c *CartServiceImpl) GetCart(ctx context.Context, userID string) (*Cart, error) {
	req := &cart.GetCartRequest{
		UserId: userID,
	}

	resp, err := c.client.CartClient.GetCart(ctx, req)
	if err != nil {
		log.Printf("获取购物车失败: %v", err)
		return nil, fmt.Errorf("获取购物车失败: %v", err)
	}

	return convertCartResponse(resp.Cart), nil
}

// UpdateCart 更新购物车
func (c *CartServiceImpl) UpdateCart(ctx context.Context, userID, productID string, quantity int) (*Cart, error) {
	req := &cart.UpdateCartRequest{
		UserId:    userID,
		ProductId: productID,
		Quantity:  int32(quantity),
	}

	resp, err := c.client.CartClient.UpdateCart(ctx, req)
	if err != nil {
		log.Printf("更新购物车失败: %v", err)
		return nil, fmt.Errorf("更新购物车失败: %v", err)
	}

	return convertCartResponse(resp.Cart), nil
}

// RemoveFromCart 从购物车中移除商品
func (c *CartServiceImpl) RemoveFromCart(ctx context.Context, userID, productID string) (*Cart, error) {
	req := &cart.RemoveFromCartRequest{
		UserId:    userID,
		ProductId: productID,
	}

	resp, err := c.client.CartClient.RemoveFromCart(ctx, req)
	if err != nil {
		log.Printf("从购物车中移除商品失败: %v", err)
		return nil, fmt.Errorf("从购物车中移除商品失败: %v", err)
	}

	return convertCartResponse(resp.Cart), nil
}

// OrderService 实现

// CreateOrder 创建订单
func (o *OrderServiceImpl) CreateOrder(ctx context.Context, userID string, items []OrderItem, address string) (*Order, error) {
	orderItems := make([]*order.OrderItem, len(items))
	for i, item := range items {
		orderItems[i] = &order.OrderItem{
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			Price:       float32(item.Price),
			Quantity:    int32(item.Quantity),
		}
	}

	req := &order.CreateOrderRequest{
		UserId:  userID,
		Items:   orderItems,
		Address: address,
	}

	resp, err := o.client.OrderClient.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("创建订单失败: %v", err)
		return nil, fmt.Errorf("创建订单失败: %v", err)
	}

	return convertOrderResponse(resp.Order), nil
}

// GetOrder 获取订单详情
func (o *OrderServiceImpl) GetOrder(ctx context.Context, orderID string) (*Order, error) {
	req := &order.GetOrderRequest{
		OrderId: orderID,
	}

	resp, err := o.client.OrderClient.GetOrder(ctx, req)
	if err != nil {
		log.Printf("获取订单详情失败: %v", err)
		return nil, fmt.Errorf("获取订单详情失败: %v", err)
	}

	return convertOrderResponse(resp.Order), nil
}

// CancelOrder 取消订单
func (o *OrderServiceImpl) CancelOrder(ctx context.Context, orderID string) (*Order, error) {
	req := &order.CancelOrderRequest{
		OrderId: orderID,
	}

	resp, err := o.client.OrderClient.CancelOrder(ctx, req)
	if err != nil {
		log.Printf("取消订单失败: %v", err)
		return nil, fmt.Errorf("取消订单失败: %v", err)
	}

	return convertOrderResponse(resp.Order), nil
}

// RequestRefund 申请退款
func (o *OrderServiceImpl) RequestRefund(ctx context.Context, orderID string, reason string) (bool, error) {
	req := &order.RequestRefundRequest{
		OrderId: orderID,
		Reason:  reason,
	}

	resp, err := o.client.OrderClient.RequestRefund(ctx, req)
	if err != nil {
		log.Printf("申请退款失败: %v", err)
		return false, fmt.Errorf("申请退款失败: %v", err)
	}

	return resp.Success, nil
}

// 工具函数 - 转换Cart响应
func convertCartResponse(cartResp *cart.Cart) *Cart {
	items := make([]CartItem, len(cartResp.Items))
	for i, item := range cartResp.Items {
		items[i] = CartItem{
			ProductID:   item.ProductId,
			ProductName: item.ProductName,
			Price:       float64(item.Price),
			Quantity:    int(item.Quantity),
		}
	}

	return &Cart{
		UserID:     cartResp.UserId,
		Items:      items,
		TotalPrice: float64(cartResp.TotalPrice),
	}
}

// 工具函数 - 转换Order响应
func convertOrderResponse(orderResp *order.Order) *Order {
	items := make([]OrderItem, len(orderResp.Items))
	for i, item := range orderResp.Items {
		items[i] = OrderItem{
			ProductID:   item.ProductId,
			ProductName: item.ProductName,
			Price:       float64(item.Price),
			Quantity:    int(item.Quantity),
		}
	}

	return &Order{
		OrderID:    orderResp.OrderId,
		UserID:     orderResp.UserId,
		Items:      items,
		TotalPrice: float64(orderResp.TotalPrice),
		Status:     orderResp.Status,
		CreateTime: orderResp.CreateTime,
		Address:    orderResp.Address,
	}
}

// ActionResult 动作结果
type ActionResult struct {
	ActionType string
	TargetID   string
	Result     string
}

// ProcessResult 处理结果
type ProcessResult struct {
	Message string
	Actions []ActionResult
}

// MessageHistory 消息历史
type MessageHistory struct {
	Role      string
	Content   string
	Timestamp int64
}

// ProcessUserMessage 处理用户消息
func ProcessUserMessage(ctx context.Context, msg *UserMessage) (*ProcessResult, error) {
	// 创建结果
	result := &ProcessResult{
		Actions: make([]ActionResult, 0),
	}

	// 获取会话ID
	sessionID := msg.Context["session_id"]
	if sessionID == "" {
		sessionID = fmt.Sprintf("session_%d", time.Now().Unix())
	}

	// 获取内存管理器
	memory := GetMemory()
	if memory == nil {
		return nil, fmt.Errorf("memory manager not initialized")
	}

	// 获取历史消息
	history := memory.Get(sessionID)
	messages := make([]*schema.Message, 0)

	// 添加历史消息
	for _, h := range history {
		messages = append(messages, h...)
	}

	// 添加新消息
	messages = append(messages, schema.UserMessage(msg.Query))

	// 创建Agent
	agent, err := NewAgent(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("create agent failed: %v", err)
	}

	// 调用Agent处理
	resp, err := agent.agent.Run(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("agent run failed: %v", err)
	}

	// 保存会话历史
	memory.Add(sessionID, []*schema.Message{
		schema.UserMessage(msg.Query),
		schema.AssistantMessage(resp.Content),
	})

	// 设置结果
	result.Message = resp.Content

	// 解析动作
	for _, tool := range resp.Tools {
		result.Actions = append(result.Actions, ActionResult{
			ActionType: tool.Name,
			TargetID:   tool.Arguments,
			Result:     tool.Output,
		})
	}

	return result, nil
}

// GetSessionHistory 获取会话历史
func GetSessionHistory(ctx context.Context, userID, sessionID string) ([]MessageHistory, error) {
	// 获取内存管理器
	memory := GetMemory()
	if memory == nil {
		return nil, fmt.Errorf("memory manager not initialized")
	}

	// 获取历史消息
	history := memory.Get(sessionID)
	if len(history) == 0 {
		return nil, nil
	}

	// 转换为MessageHistory
	result := make([]MessageHistory, 0)
	for _, msgs := range history {
		for _, msg := range msgs {
			result = append(result, MessageHistory{
				Role:      msg.Role,
				Content:   msg.Content,
				Timestamp: time.Now().Unix(), // 这里应该使用实际的时间戳，但示例中简化处理
			})
		}
	}

	return result, nil
}
