# *** Project

## introduce

- Use the [Kitex](https://github.com/cloudwego/kitex/) framework
- Generating the base code for unit tests.
- Provides basic config functions
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| handler.go  | Used for request processing return of response. |
| kitex_gen  | kitex generated code |
| biz/service  | The actual business logic. |
| biz/dal  | Logic for operating the storage layer |

## How to run

```shell
sh build.sh
sh output/bootstrap.sh
```

- 多级缓存
- 对象存储
- 创建时发送消息
- 调用库存服务并缓存
  - 库存服务不可用时（降级）：
  - 返回最近一次缓存值
  - 显示"库存充足"兜底文案
  - 标记库存状态为"查询中"
- 灾备方案：1. 消息重放机制 2. 数据一致性验证