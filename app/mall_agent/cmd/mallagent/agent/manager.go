package agent

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/strings77wzq/gomall/app/mall_agent/eino/mallagent"
)

// Manager 管理全局Agent实例
type Manager struct {
	agent  *mallagent.ReactAgent
	memory *mallagent.Memory
	once   sync.Once
}

var (
	globalManager *Manager
	managerMu     sync.Mutex
)

// InitManager 初始化全局Manager
func InitManager(ctx context.Context, registry discovery.Registry) error {
	managerMu.Lock()
	defer managerMu.Unlock()

	if globalManager != nil {
		return nil
	}

	manager := &Manager{}
	if err := manager.init(ctx, registry); err != nil {
		return err
	}

	globalManager = manager
	return nil
}

// GetManager 获取全局Manager实例
func GetManager() *Manager {
	return globalManager
}

func (m *Manager) init(ctx context.Context, registry discovery.Registry) error {
	var initErr error
	m.once.Do(func() {
		// 初始化内存管理
		if err := mallagent.InitMemory(); err != nil {
			initErr = fmt.Errorf("init memory failed: %v", err)
			return
		}
		m.memory = mallagent.GetMemory()

		// 初始化Agent
		agent, err := mallagent.NewReactAgent(ctx)
		if err != nil {
			initErr = fmt.Errorf("create agent failed: %v", err)
			return
		}
		m.agent = agent
	})
	return initErr
}
