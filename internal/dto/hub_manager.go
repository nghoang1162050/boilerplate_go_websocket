package dto

import "sync"

type HubManager struct {
	hubs map[string]*Hub
	mu   sync.RWMutex
}

func NewHubManager() *HubManager {
	return &HubManager{
		hubs: make(map[string]*Hub),
	}
}

func (m *HubManager) InitHub(id string) *Hub {
    m.mu.Lock()
    defer m.mu.Unlock()
    if hub, ok := m.hubs[id]; ok {
        return hub
    }
    hub := NewHub()
    m.hubs[id] = hub
    go hub.Run()
    return hub
}

func (m *HubManager) GetHub(id string) (*Hub, bool) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    hub, ok := m.hubs[id]
    return hub, ok
}

func (m *HubManager) CloseHub(id string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    if hub, ok := m.hubs[id]; ok {
        hub.Close()
        delete(m.hubs, id)
    }
}
