package config

import (
	"os"
	"strconv"
)

// Config 总配置结构
type Config struct {
	Server ServerConfig
	Log    LogConfig
	Etcd   EtcdConfig
	Redis  RedisConfig
	Model  ModelConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Address string
	Debug   bool
}

// LogConfig 日志配置
type LogConfig struct {
	Level  string
	Output string
}

// EtcdConfig ETCD配置
type EtcdConfig struct {
	Endpoints []string
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// ModelConfig 模型配置
type ModelConfig struct {
	ChatEndpointID      string
	EmbeddingEndpointID string
	APIKey              string
}

// MustLoad 加载配置
func MustLoad() *Config {
	// 服务配置
	port := os.Getenv("MALL_AGENT_PORT")
	if port == "" {
		port = "8080"
	}

	debug, _ := strconv.ParseBool(os.Getenv("MALL_AGENT_DEBUG"))

	// Redis配置
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	// ETCD配置
	etcdEndpoints := []string{"localhost:2379"}
	if ep := os.Getenv("ETCD_ENDPOINTS"); ep != "" {
		etcdEndpoints = []string{ep}
	}

	return &Config{
		Server: ServerConfig{
			Address: ":" + port,
			Debug:   debug,
		},
		Log: LogConfig{
			Level:  "info",
			Output: "stdout",
		},
		Etcd: EtcdConfig{
			Endpoints: etcdEndpoints,
		},
		Redis: RedisConfig{
			Addr:     redisAddr,
			Password: "",
			DB:       0,
		},
		Model: ModelConfig{
			ChatEndpointID:      os.Getenv("ARK_CHAT_MODEL"),
			EmbeddingEndpointID: os.Getenv("ARK_EMBEDDING_MODEL"),
			APIKey:              os.Getenv("ARK_API_KEY"),
		},
	}
}

// 在现有的Config结构中添加RPCServer字段

type RPCServerConfig struct {
	Address string
}

type Config struct {
	Server    ServerConfig
	RPCServer RPCServerConfig  // 新增RPC服务器配置
	Etcd      EtcdConfig
	Redis     RedisConfig
	// 其他配置...
}
