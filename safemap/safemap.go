/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     tofuutils
 * @date        2018-01-25 19:19
 */
package tofuutils

import (
	"sync"
	"github.com/liuyongshuai/goutils/elem"
	"fmt"
)

type SafeMap struct {
	lock *sync.RWMutex
	data map[elem.ItemElem]elem.ItemElem
}

//获取实例
func NewSafeMap() *SafeMap {
	return &SafeMap{
		lock: new(sync.RWMutex),
		data: make(map[elem.ItemElem]elem.ItemElem),
	}
}

//提取值
func (m *SafeMap) Get(k elem.ItemElem) (elem.ItemElem, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.data[k]; ok {
		return val, nil
	}
	return elem.ItemElem{}, fmt.Errorf("not exists")
}

//设置值
func (m *SafeMap) Set(k elem.ItemElem, v elem.ItemElem) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.data[k]; !ok {
		m.data[k] = v
	} else if val != v {
		m.data[k] = v
	} else {
		return false
	}
	return true
}

//检查是否存在
func (m *SafeMap) Check(k elem.ItemElem) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.data[k]
	return ok
}

//干掉一个值
func (m *SafeMap) Delete(k elem.ItemElem) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.data, k)
}

//返回所有的值
func (m *SafeMap) Items() map[elem.ItemElem]elem.ItemElem {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make(map[elem.ItemElem]elem.ItemElem)
	for k, v := range m.data {
		r[k] = v
	}
	return r
}

//统计数量
func (m *SafeMap) Count() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.data)
}
