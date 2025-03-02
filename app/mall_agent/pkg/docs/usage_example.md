# 电商智能客服工作流程示例

本文档通过具体示例演示mall_agent智能客服系统的完整工作流程。

## 场景：查询商品、加入购物车并下单

以下示例展示了用户从查询商品到完成下单的整个交互过程。

### 1. 用户输入

用户输入如下请求：
```
请帮我查询一下最新款的iPhone手机，然后帮我加入购物车，再帮我完成下单操作
```

### 2. 解析意图

mall_agent会通过ChatModel理解用户意图，识别出三个主要任务：
- 查询iPhone手机商品信息
- 将商品加入购物车
- 创建订单

### 3. 执行任务流程

#### 步骤1：查询商品

**调用工具**：`search_product`

**请求参数**：
```json
{
  "name": "iPhone"
}
```

**微服务调用**：商品服务 (ProductService)

**服务返回结果**：
```json
{
  "products": [
    {
      "id": "p10001",
      "name": "iPhone 15 Pro Max",
      "description": "Apple最新旗舰手机，搭载A17芯片",
      "price": 9999.00,
      "category": "手机数码",
      "image_url": "https://example.com/images/iphone15.jpg",
      "created_at": 1634567890
    },
    {
      "id": "p10002",
      "name": "iPhone 15",
      "description": "Apple全新iPhone，128GB存储",
      "price": 6999.00,
      "category": "手机数码",
      "image_url": "https://example.com/images/iphone15normal.jpg",
      "created_at": 1634567890
    }
  ]
}
```

**客服回复**：
```
我找到了两款最新的iPhone手机：

1. iPhone 15 Pro Max - ¥9999.00
   Apple最新旗舰手机，搭载A17芯片

2. iPhone 15 - ¥6999.00
   Apple全新iPhone，128GB存储

请问您对哪一款感兴趣，需要我帮您加入购物车吗？
```

#### 步骤2：用户选择商品

**用户回复**：
```
我想要iPhone 15 Pro Max，帮我查一下库存
```

**调用工具**：`check_stock`

**请求参数**：
```json
{
  "product_id": "p10001"
}
```

**微服务调用**：商品服务 (ProductService)

**服务返回结果**：
```json
{
  "product_id": "p10001",
  "stock": 42,
  "available": true
}
```

**客服回复**：
```
iPhone 15 Pro Max目前有货，库存数量为42台。您要将它加入购物车吗？
```

#### 步骤3：加入购物车

**用户回复**：
```
好的，请帮我加入购物车
```

**调用工具**：`add_to_cart`

**请求参数**：
```json
{
  "user_id": "u2001",
  "product_id": "p10001",
  "quantity": 1
}
```

**微服务调用**：购物车服务 (CartService)

**服务返回结果**：
```json
{
  "user_id": "u2001",
  "items": [
    {
      "product_id": "p10001",
      "product_name": "iPhone 15 Pro Max",
      "price": 9999.00,
      "quantity": 1,
      "subtotal": 9999.00
    }
  ],
  "total": 9999.00,
  "updated_at": "2023-10-20T14:30:45Z"
}
```

**客服回复**：
```
已成功将iPhone 15 Pro Max加入购物车！

您的购物车现在有：
- 1台 iPhone 15 Pro Max，单价：¥9999.00，小计：¥9999.00

购物车总计：¥9999.00

请问您需要继续购物还是现在就下单？
```

#### 步骤4：创建订单

**用户回复**：
```
我要下单，使用我的默认地址
```

**调用工具1**：`get_user_address`

**请求参数**：
```json
{
  "user_id": "u2001"
}
```

**微服务调用**：用户服务 (UserService)

**服务返回结果**：
```json
[
  {
    "id": "addr001",
    "user_id": "u2001",
    "province": "广东省",
    "city": "深圳市",
    "district": "南山区",
    "street": "科技园路",
    "detail": "某某大厦10楼1001室",
    "recipient": "张三",
    "phone": "13812345678",
    "is_default": true
  },
  {
    "id": "addr002",
    "user_id": "u2001",
    "province": "北京市",
    "city": "北京市",
    "district": "海淀区",
    "street": "中关村大街",
    "detail": "创新大厦5层",
    "recipient": "张三",
    "phone": "13812345678",
    "is_default": false
  }
]
```

**调用工具2**：`get_cart`

**请求参数**：
```json
{
  "user_id": "u2001"
}
```

**微服务调用**：购物车服务 (CartService)

**服务返回结果**：
```json
{
  "user_id": "u2001",
  "items": [
    {
      "product_id": "p10001",
      "product_name": "iPhone 15 Pro Max",
      "price": 9999.00,
      "quantity": 1,
      "subtotal": 9999.00
    }
  ],
  "total": 9999.00,
  "updated_at": "2023-10-20T14:30:45Z"
}
```

**调用工具3**：`create_order`

**请求参数**：
```json
{
  "user_id": "u2001",
  "items": [
    {
      "product_id": "p10001",
      "product_name": "iPhone 15 Pro Max",
      "price": 9999.00,
      "quantity": 1,
      "subtotal": 9999.00
    }
  ],
  "address": "广东省深圳市南山区科技园路某某大厦10楼1001室"
}
```

**微服务调用**：订单服务 (OrderService)

**服务返回结果**：
```json
{
  "id": "o30001",
  "user_id": "u2001",
  "items": [
    {
      "product_id": "p10001",
      "product_name": "iPhone 15 Pro Max",
      "price": 9999.00,
      "quantity": 1,
      "subtotal": 9999.00
    }
  ],
  "total": 9999.00,
  "status": "pending",
  "address": "广东省深圳市南山区科技园路某某大厦10楼1001室",
  "created_at": "2023-10-20T14:35:10Z",
  "updated_at": "2023-10-20T14:35:10Z"
}
```

**客服回复**：
```
订单已成功创建！

订单号：o30001
商品：iPhone 15 Pro Max x 1台
总金额：¥9999.00
收货地址：广东省深圳市南山区科技园路某某大厦10楼1001室
收件人：张三
联系电话：138****5678

订单当前状态为：待付款

您可以通过"我的订单"查看详情并完成支付。请在24小时内完成支付，逾期订单将自动取消。

还有什么可以帮到您的吗？
```

## 整体流程分析

上述示例展示了mall_agent如何通过以下步骤完成用户任务：

1. **理解用户意图**：通过NLP技术解析用户输入，提取关键信息和操作意图
2. **调用相关工具**：根据意图选择合适的工具，准备参数
3. **微服务交互**：通过RPC与相应的微服务进行通信
4. **结果处理**：将服务返回的数据转换为用户友好的自然语言回复
5. **引导用户**：在流程的每一步提供明确指引，并询问下一步操作

整个过程中，mall_agent充当了用户与各个微服务之间的智能中介，使用户能够通过自然语言完成复杂的电商操作，提供了无缝且高效的购物体验。 