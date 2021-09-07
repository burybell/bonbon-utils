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

	mapped := NewMap().Set("name", "张三").Set("age", 3).Mapped()
	fmt.Println(mapped)
}