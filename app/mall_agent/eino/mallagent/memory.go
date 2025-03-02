package mallagent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cloudwego/eino/components/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/go-redis/redis/v8"
)

// Memory 会话内存管理
type Memory struct {
	redis *redis.Client
}

// NewMemory 创建内存管理实例
func NewMemory(addr string) (*Memory, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("connect to redis failed: %v", err)
	}

	return &Memory{redis: client}, nil
}

// Save 保存会话历史
func (m *Memory) Save(ctx context.Context, sessionID string, messages []*schema.Message) error {
	data, err := json.Marshal(messages)
	if err != nil {
		return fmt.Errorf("marshal messages failed: %v", err)
	}

	// 设置24小时过期
	key := fmt.Sprintf("chat:%s", sessionID)
	if err := m.redis.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("save to redis failed: %v", err)
	}

	return nil
}

// Get 获取会话历史
func (m *Memory) Get(ctx context.Context, sessionID string) [][]*schema.Message {
	key := fmt.Sprintf("chat:%s", sessionID)
	data, err := m.redis.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return make([][]*schema.Message, 0)
		}
		return make([][]*schema.Message, 0)
	}

	var messages []*schema.Message
	if err := json.Unmarshal(data, &messages); err != nil {
		return make([][]*schema.Message, 0)
	}

	return [][]*schema.Message{messages}
}

// 全局内存实例
var defaultMemory *Memory

// InitMemory 初始化内存管理
func InitMemory(addr string) error {
	var err error
	defaultMemory, err = NewMemory(addr)
	return err
}

// GetMemory 获取内存管理实例
func GetMemory() *Memory {
	return defaultMemory
}

// MemoryCache 基于内存的缓存实现
type MemoryCache struct {
	sessions     map[string][]string  // 会话历史记录
	agentMemory  map[string]AgentData // 代理数据
	mutex        sync.RWMutex
	maxCacheSize int
}

// AgentData 代理数据结构
type AgentData struct {
	Thoughts []string // 代理思考过程
	Actions  []string // 代理执行的操作
	LastUsed time.Time
}

var (
	// 全局会话缓存实例
	memoryCache *MemoryCache
	once        sync.Once
)

// NewMemoryCache 创建新的内存缓存
func NewMemoryCache() *MemoryCache {
	once.Do(func() {
		memoryCache = &MemoryCache{
			sessions:     make(map[string][]string),
			agentMemory:  make(map[string]AgentData),
			maxCacheSize: 1000, // 最大缓存条目数
		}
	})
	return memoryCache
}

// GetCache 实现 react.AgentCache 接口
func (m *MemoryCache) GetCache(ctx context.Context, key string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	data, ok := m.agentMemory[key]
	if !ok {
		return "", fmt.Errorf("缓存项未找到: %s", key)
	}

	// 简单的缓存格式化
	result := fmt.Sprintf("思考: %v\n操作: %v", data.Thoughts, data.Actions)
	return result, nil
}

// SetCache 实现 react.AgentCache 接口
func (m *MemoryCache) SetCache(ctx context.Context, key string, value *react.AgentMemory) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// 如果缓存已满，清理最老的条目
	if len(m.agentMemory) >= m.maxCacheSize {
		m.cleanOldestEntries()
	}

	// 构建代理数据
	data := AgentData{
		Thoughts: value.Thoughts,
		Actions:  value.Actions,
		LastUsed: time.Now(),
	}

	m.agentMemory[key] = data
	return nil
}

// cleanOldestEntries 清理最老的缓存条目
func (m *MemoryCache) cleanOldestEntries() {
	// 移除四分之一的最老条目
	removeCount := m.maxCacheSize / 4
	if removeCount <= 0 {
		removeCount = 1
	}

	// 收集所有条目
	type keyTime struct {
		key  string
		time time.Time
	}
	entries := make([]keyTime, 0, len(m.agentMemory))
	for k, v := range m.agentMemory {
		entries = append(entries, keyTime{k, v.LastUsed})
	}

	// 按最后使用时间排序
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].time.After(entries[j].time) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}

	// 删除最老的条目
	for i := 0; i < removeCount && i < len(entries); i++ {
		delete(m.agentMemory, entries[i].key)
	}
}

// AddSessionMessage 添加会话消息
func (m *MemoryCache) AddSessionMessage(userID, message string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.sessions[userID]; !ok {
		m.sessions[userID] = make([]string, 0)
	}

	m.sessions[userID] = append(m.sessions[userID], message)

	// 限制历史记录长度，保留最新的20条
	maxHistory := 20
	if len(m.sessions[userID]) > maxHistory {
		m.sessions[userID] = m.sessions[userID][len(m.sessions[userID])-maxHistory:]
	}
}

// GetSessionHistory 获取会话历史
func GetSessionHistory(ctx context.Context, userID string, sessionID string) ([]Message, error) {
	cache := NewMemoryCache()
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	key := fmt.Sprintf("%s:%s", userID, sessionID)
	history, ok := cache.sessions[key]
	if !ok {
		return []Message{}, nil
	}

	// 将字符串历史记录转换为Message结构
	messages := make([]Message, 0, len(history))
	for _, msg := range history {
		// 简单解析消息格式，实际项目中可能需要更复杂的解析
		role := "user"
		if len(msg) > 0 && msg[0] == '>' {
			role = "assistant"
			msg = msg[1:]
		}

		messages = append(messages, Message{
			Role:      role,
			Content:   msg,
			Timestamp: time.Now().Unix(), // 这里使用当前时间，实际应该存储原始时间戳
		})
	}

	return messages, nil
}

// Message 定义消息结构
type Message struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

// ClearSession 清除会话历史
func ClearSession(ctx context.Context, userID string) error {
	cache := NewMemoryCache()
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if _, ok := cache.sessions[userID]; !ok {
		return errors.New("会话不存在")
	}

	delete(cache.sessions, userID)
	return nil
}
