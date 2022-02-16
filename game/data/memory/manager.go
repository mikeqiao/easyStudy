package memory

import "sync"

type Manager struct {
	dataList sync.Map
}

//从管理列表中获取数据
func (m *Manager) GetData(key interface{}) interface{} {
	d, ok := m.dataList.Load(key)
	if ok && nil != d {
		return d
	}
	return nil
}

//向管理列表中添加数据
func (m *Manager) AddData(key, value interface{}) bool {
	d, ok := m.dataList.Load(key)
	if ok && nil != d {
		return false
	}
	m.dataList.Store(key, value)
	return true
}

//从管理列表中删除数据
func (m *Manager) DelData(key interface{}) bool {
	_, ok := m.dataList.Load(key)
	if ok {
		m.dataList.Delete(key)
		return true
	}
	return false
}
