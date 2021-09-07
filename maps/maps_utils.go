package maps

import "sync"

type Tuple struct {
	Key interface{}
	Value interface{}
}

type CommonTuple struct {
	Key string
	Value interface{}
}

// EasyMap 二维数字转map
func EasyMap(kvs [][]interface{}) map[interface{}]interface{}{
	mp := make(map[interface{}]interface{})
	for _,kv := range kvs {
		if len(kv) == 2 {
			mp[kv[0]] = kv[1]
		}
	}
	return mp
}

type ChainMap struct {
	data map[interface{}]interface{}
}

// NewChainMap 创建一个Map
func NewChainMap(params ...int) * ChainMap{
	if len(params) > 0 && params[0] > 0{
		return &ChainMap{data: make(map[interface{}]interface{}, params[0])}
	} else {
		return &ChainMap{data: make(map[interface{}]interface{})}
	}
}

// Set 链式调用，设置值
func (m *ChainMap) Set(key interface{},value interface{}) * ChainMap {
	m.data[key] = value
	return m
}

// Get 获取值
func (m *ChainMap) Get(key interface{}) interface{} {
	return m.data[key]
}

// Gets 获取原生map
func (m *ChainMap) Gets() map[interface{}]interface{} {
	return m.data
}

type Map interface {
	Get(key string) interface{}
	Put(key string,value interface{})
	GetOrPut(key string,value interface{}) interface{}
	Keys() []string
	Remove(key string) bool
	Clear()
	Iter() <-chan CommonTuple
}

// CommonMap 键为string的map
type CommonMap struct {
	data map[string]interface{}
}

// NewCommonMap 创建map
func NewCommonMap(params ...int) * CommonMap {
	if len(params) > 0 && params[0] > 0{
		return &CommonMap{data: make(map[string]interface{}, params[0])}
	} else {
		return &CommonMap{data: make(map[string]interface{})}
	}
}

// Get get值
func (cm *CommonMap) Get(key string) interface{} {
	if raw, ok := cm.data[key];ok {
		return raw
	} else {
		return nil
	}
}

// Put put值
func (cm *CommonMap) Put(key string,value interface{})  {
	cm.data[key] = value
}

// GetOrPut 首先get，get不到就设置值
func (cm *CommonMap) GetOrPut(key string,value interface{}) interface{} {
	if raw, ok := cm.data[key];ok {
		return raw
	} else {
		cm.data[key] = value
		return value
	}
}

// Keys 获取所有key
func (cm *CommonMap) Keys() []string {
	keys := make([]string,0,len(cm.data))
	for key := range cm.data {
		keys = append(keys, key)
	}
	return keys
}

// Remove 移除key
func (cm *CommonMap) Remove(key string) bool  {
	if _, ok := cm.data[key];ok {
		delete(cm.data, key)
		return true
	}
	return false
}

// Clear 清空
func (cm *CommonMap) Clear() {
	cm.data = map[string]interface{}{}
}

// Iter 迭代器
func (cm *CommonMap) Iter() <-chan CommonTuple{
	ch := make(chan CommonTuple)
	go func() {
		for key,value := range cm.data {
			ch <- CommonTuple{
				Key:   key,
				Value: value,
			}
		}
		close(ch)
	}()
	return ch
}

// LockCommonMap 带锁的map
type LockCommonMap struct {
	CommonMap
	sync.RWMutex
}

// NewLockCommonMap 初始化
func NewLockCommonMap(params ...int) * LockCommonMap {
	return &LockCommonMap{*NewCommonMap(params...), sync.RWMutex{}}
}

// Get 获取带锁
func (lcm *LockCommonMap) Get(key string) interface{} {
	lcm.Lock()
	defer lcm.Unlock()
	return lcm.CommonMap.Get(key)
}

// Put put值
func (lcm *LockCommonMap) Put(key string,value interface{})  {
	lcm.Lock()
	defer lcm.Unlock()
	lcm.CommonMap.Put(key,value)
}

// GetOrPut 首先get，get不到就设置值
func (lcm *LockCommonMap) GetOrPut(key string,value interface{}) interface{} {
	lcm.Lock()
	defer lcm.Unlock()
	return lcm.CommonMap.GetOrPut(key,value)
}

// Keys 获取所有key
func (lcm *LockCommonMap) Keys() []string {
	lcm.Lock()
	defer lcm.Unlock()
	return lcm.CommonMap.Keys()
}

// Remove 移除key
func (lcm *LockCommonMap) Remove(key string) bool {
	lcm.Lock()
	defer lcm.Unlock()
	return lcm.CommonMap.Remove(key)
}

func (lcm *LockCommonMap) Clear() {
	lcm.Lock()
	defer lcm.Unlock()
	lcm.CommonMap.Clear()
}

// Iter 迭代器
func (lcm *LockCommonMap) Iter() <-chan CommonTuple {
	ch := make(chan CommonTuple)
	go func() {
		lcm.Lock()
		for key,value := range lcm.data {
			ch <- CommonTuple{
				Key:   key,
				Value: value,
			}
		}
		close(ch)
		lcm.Unlock()
	}()
	return ch
}