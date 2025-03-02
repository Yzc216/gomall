package discovery

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// EtcdConfig ETCD配置
type EtcdConfig struct {
	Endpoints []string
}

// NewEtcdRegistry 创建ETCD注册中心
func NewEtcdRegistry(config EtcdConfig) (registry.Registry, error) {
	return etcd.NewEtcdRegistry(config.Endpoints)
}

// NewEtcdResolver 创建ETCD解析器
func NewEtcdResolver(config EtcdConfig) (discovery.Resolver, error) {
	return etcd.NewEtcdResolver(config.Endpoints)
}
