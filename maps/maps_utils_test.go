package maps

import (
	"fmt"
	"testing"
)

func TestEasyMap(t *testing.T) {
	easyMap := EasyMap([][]interface{}{
		{"name", "张三"},
		{"age", 20},
	})
	fmt.Println(easyMap)
}

func TestNewMap(t *testing.T) {

	mapped := NewChainMap().Set("name", "张三").Set("age", 3).Gets()
	fmt.Println(mapped)
}

func TestCommonMap(t *testing.T) {

	commonMap := NewCommonMap(20)
	commonMap.Put("name","张三")
	commonMap.Put("age",30)
	fmt.Println(commonMap.Keys())
}


func TestLockCommonMap(t *testing.T) {

	commonMap := NewLockCommonMap(20)
	commonMap.Put("name","张三")
	commonMap.Put("age",30)
	fmt.Println(commonMap.Keys())
}