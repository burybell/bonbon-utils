package maps

/**
[
	["key","value"]
]

 */

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

type Map struct {
	data map[interface{}]interface{}
}

// NewMap 创建一个Map
func NewMap() * Map{
	return &Map{data: make(map[interface{}]interface{})}
}

// Set 链式调用，设置值
func (m *Map) Set(key interface{},value interface{}) * Map {
	m.data[key] = value
	return m
}

// Get 获取值
func (m *Map) Get(key interface{}) interface{} {
	return m.data[key]
}

// Mapped 获取原生map
func (m *Map) Mapped() map[interface{}]interface{} {
	return m.data
}