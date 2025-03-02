package mallagent

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// 商品服务接口
type ProductService interface {
	QueryProduct(ctx context.Context, name string) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	CheckStock(ctx context.Context, id string) (int, error)
}

// 购物车服务接口
type CartService interface {
	AddToCart(ctx context.Context, userID, productID string, quantity int) (*Cart, error)
	GetCart(ctx context.Context, userID string) (*Cart, error)
	UpdateCart(ctx context.Context, userID, productID string, quantity int) (*Cart, error)
	RemoveFromCart(ctx context.Context, userID, productID string) (*Cart, error)
}

// 订单服务接口
type OrderService interface {
	CreateOrder(ctx context.Context, userID string, items []OrderItem, address string) (*Order, error)
	GetOrder(ctx context.Context, orderID string) (*Order, error)
	CancelOrder(ctx context.Context, orderID string) (*Order, error)
	RequestRefund(ctx context.Context, orderID string, reason string) (bool, error)
}

// 用户服务接口
type UserService interface {
	GetUserInfo(ctx context.Context, userID string) (*User, error)
	UpdateUserInfo(ctx context.Context, userID string, info map[string]interface{}) (*User, error)
	GetUserAddress(ctx context.Context, userID string) ([]Address, error)
	AddUserAddress(ctx context.Context, userID string, address *Address) (*Address, error)
}

// 商品服务工具
func NewProductTools(svc ProductService) []tool.BaseTool {
	return []tool.BaseTool{
		// 查询商品工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "search_product",
			Description: "搜索商品信息",
			Parameters: []*schema.Parameter{
				{Name: "name", Type: "string", Required: true, Description: "商品名称关键词"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.QueryProduct(ctx, args["name"].(string))
		}),

		// 获取商品详情工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "get_product",
			Description: "获取商品详细信息",
			Parameters: []*schema.Parameter{
				{Name: "id", Type: "string", Required: true, Description: "商品ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.GetProduct(ctx, args["id"].(string))
		}),

		// 检查库存工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "check_stock",
			Description: "检查商品库存",
			Parameters: []*schema.Parameter{
				{Name: "product_id", Type: "string", Required: true, Description: "商品ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			stock, err := svc.CheckStock(ctx, args["product_id"].(string))
			return map[string]interface{}{
				"product_id": args["product_id"].(string),
				"stock":      stock,
				"available":  stock > 0,
			}, err
		}),
	}
}

// 购物车工具
func NewCartTools(svc CartService) []tool.BaseTool {
	return []tool.BaseTool{
		// 添加到购物车工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "add_to_cart",
			Description: "将商品添加到购物车",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
				{Name: "product_id", Type: "string", Required: true, Description: "商品ID"},
				{Name: "quantity", Type: "integer", Required: true, Description: "数量"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			quantity := int(args["quantity"].(float64))
			return svc.AddToCart(ctx, args["user_id"].(string), args["product_id"].(string), quantity)
		}),

		// 查看购物车工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "get_cart",
			Description: "查看购物车内容",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.GetCart(ctx, args["user_id"].(string))
		}),

		// 更新购物车工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "update_cart",
			Description: "更新购物车商品数量",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
				{Name: "product_id", Type: "string", Required: true, Description: "商品ID"},
				{Name: "quantity", Type: "integer", Required: true, Description: "新数量"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			quantity := int(args["quantity"].(float64))
			return svc.UpdateCart(ctx, args["user_id"].(string), args["product_id"].(string), quantity)
		}),

		// 从购物车移除商品工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "remove_from_cart",
			Description: "从购物车移除商品",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
				{Name: "product_id", Type: "string", Required: true, Description: "商品ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.RemoveFromCart(ctx, args["user_id"].(string), args["product_id"].(string))
		}),
	}
}

// 订单工具
func NewOrderTools(svc OrderService) []tool.BaseTool {
	return []tool.BaseTool{
		// 创建订单工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "create_order",
			Description: "创建新订单",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
				{Name: "items", Type: "array", Required: true, Description: "订单商品列表"},
				{Name: "address", Type: "string", Required: true, Description: "收货地址"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			itemsData := args["items"].([]interface{})
			items := make([]OrderItem, 0, len(itemsData))

			for _, item := range itemsData {
				itemMap := item.(map[string]interface{})
				items = append(items, OrderItem{
					ProductID:   itemMap["product_id"].(string),
					ProductName: itemMap["product_name"].(string),
					Price:       itemMap["price"].(float64),
					Quantity:    int(itemMap["quantity"].(float64)),
					Subtotal:    itemMap["subtotal"].(float64),
				})
			}

			return svc.CreateOrder(ctx, args["user_id"].(string), items, args["address"].(string))
		}),

		// 查询订单工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "get_order",
			Description: "查询订单状态",
			Parameters: []*schema.Parameter{
				{Name: "order_id", Type: "string", Required: true, Description: "订单ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.GetOrder(ctx, args["order_id"].(string))
		}),

		// 取消订单工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "cancel_order",
			Description: "取消订单",
			Parameters: []*schema.Parameter{
				{Name: "order_id", Type: "string", Required: true, Description: "订单ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.CancelOrder(ctx, args["order_id"].(string))
		}),

		// 申请退款工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "request_refund",
			Description: "申请订单退款",
			Parameters: []*schema.Parameter{
				{Name: "order_id", Type: "string", Required: true, Description: "订单ID"},
				{Name: "reason", Type: "string", Required: true, Description: "退款原因"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			success, err := svc.RequestRefund(ctx, args["order_id"].(string), args["reason"].(string))
			return map[string]interface{}{
				"order_id": args["order_id"].(string),
				"success":  success,
			}, err
		}),
	}
}

// 用户服务工具
func NewUserTools(svc UserService) []tool.BaseTool {
	return []tool.BaseTool{
		// 获取用户信息工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "get_user_info",
			Description: "获取用户基本信息",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.GetUserInfo(ctx, args["user_id"].(string))
		}),

		// 更新用户信息工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "update_user_info",
			Description: "更新用户信息",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
				{Name: "info", Type: "object", Required: true, Description: "需要更新的用户信息"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.UpdateUserInfo(ctx, args["user_id"].(string), args["info"].(map[string]interface{}))
		}),

		// 获取用户地址工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "get_user_address",
			Description: "获取用户保存的地址列表",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			return svc.GetUserAddress(ctx, args["user_id"].(string))
		}),

		// 添加用户地址工具
		tool.NewTool(&schema.ToolInfo{
			Name:        "add_user_address",
			Description: "为用户添加新地址",
			Parameters: []*schema.Parameter{
				{Name: "user_id", Type: "string", Required: true, Description: "用户ID"},
				{Name: "address", Type: "object", Required: true, Description: "地址信息"},
			},
		}, func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			addressMap := args["address"].(map[string]interface{})
			address := &Address{
				Province:  addressMap["province"].(string),
				City:      addressMap["city"].(string),
				District:  addressMap["district"].(string),
				Street:    addressMap["street"].(string),
				Detail:    addressMap["detail"].(string),
				Recipient: addressMap["recipient"].(string),
				Phone:     addressMap["phone"].(string),
				IsDefault: addressMap["is_default"].(bool),
			}
			return svc.AddUserAddress(ctx, args["user_id"].(string), address)
		}),
	}
}

// GetTools 获取所有工具
func GetTools(ctx context.Context) ([]tool.BaseTool, error) {
	// 获取服务客户端
	clients := GetClients()
	if clients == nil {
		return nil, fmt.Errorf("service clients not initialized")
	}

	// 合并所有工具
	var tools []tool.BaseTool

	// 添加商品服务工具
	productTools := NewProductTools(clients.ProductClient)
	tools = append(tools, productTools...)

	// 添加购物车工具
	cartTools := NewCartTools(clients.CartClient)
	tools = append(tools, cartTools...)

	// 添加订单工具
	orderTools := NewOrderTools(clients.OrderClient)
	tools = append(tools, orderTools...)

	// 添加用户工具
	if clients.UserClient != nil {
		userTools := NewUserTools(clients.UserClient)
		tools = append(tools, userTools...)
	}

	return tools, nil
}
