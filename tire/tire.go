package tire

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type View struct {
	value []rune
	level int
}

func (v View) String() string {
	return fmt.Sprintf("value = %s, level = %d", string(v.value), v.level)
}

// Tire 前缀树
type Tire struct {
	self rune
	flag bool
	value []rune
	level int
	child map[rune]*Tire
}

// NewTire 初始化Root前缀树
func NewTire() *Tire {
	return &Tire{
		self:  0,
		flag: false,
		value: nil,
		level: 0,
		child: make(map[rune]*Tire),
	}
}

// GetTire 获取值
func (tree *Tire) GetTire(raw []rune,index int, flag bool, level int) *Tire {
	if flag {
		return &Tire{
			self:  raw[index],
			flag:  true,
			value: raw[0:index + 1],
			level: level,
			child: make(map[rune]*Tire),
		}
	} else {
		return &Tire{
			self:  raw[index],
			flag:  false,
			value: nil,
			level: level,
			child: make(map[rune]*Tire),
		}
	}
}

func (tree *Tire) Index(raw []rune)  {
	var hand *Tire = tree
	for i := range raw {
		var tire *Tire
		if len(raw) - 1 > i {
			tire = tree.GetTire(raw, i, false, i)
		} else {
			tire = tree.GetTire(raw, i, true, i)
		}
		// 存在
		if _,ok := hand.child[raw[i]];ok {
			// 最后一个
			if len(raw) - 1 == i {
				hand.child[raw[i]].flag = true
				hand.child[raw[i]].value = raw
			}
		} else {
			// 不存在
			hand.child[raw[i]] = tire
		}
		hand = hand.child[raw[i]]
	}
}

// Values 从某节点获取值
func (tree *Tire) Values(node *Tire) []View {
	if node != nil {
		values := make([]View,0)
		if node.child != nil && len(node.child) > 0 {
			for key := range node.child {
				if node.child[key].flag {
					values = append(values, View{
						value: node.child[key].value,
						level: node.child[key].level,
					})
				}
				views := tree.Values(node.child[key])
				values = append(values, views...)
			}
		}
		return values
	} else {
		return nil
	}
}

func (tree *Tire) Search(raw []rune) []View {
	var hand *Tire = tree
	for i := range raw {
		if _,ok := hand.child[raw[i]];ok {
			hand = hand.child[raw[i]]
		} else {
			return nil
		}
	}

	views := make([]View, 0)
	if hand != nil && hand.flag {
		views = append(views, View{
			value: hand.value,
			level: hand.level,
		})
	}
	values := tree.Values(hand)
	views = append(views, values...)
	return views
}

func (tree *Tire) Analysis(raw []rune) []View {
	var hand *Tire = tree
	views := make([]View, 0)
	for i := range raw {
		if _,ok := hand.child[raw[i]];ok {
			if hand.flag {
				views = append(views, View{
					value: hand.value,
					level: hand.level,
				})
			}
			hand = hand.child[raw[i]]
		} else {
			return views
		}
	}
	if hand != nil && hand.flag {
		views = append(views, View{
			value: hand.value,
			level: hand.level,
		})
	}
	return views
}

// IkAnalysis Ik分词器
type Searcher struct {
	tree Tire

}

func NewSearcher(dices ...string) *Searcher {
	tree := NewTire()
	for di := range dices {
		file, err := ioutil.ReadFile(dices[di])
		if err != nil {
			return nil
		}
		split := strings.Split(string(file), "\n")

		for i := range split {
			tree.Index([]rune(strings.Trim(split[i],"\r \n\t")))
		}
	}
	return &Searcher{tree: *tree}
}

func (ik *Searcher) Search(word string) []string {
	values := make([]string, 0)
	search := ik.tree.Search([]rune(word))
	for i := range search {
		values = append(values, string(search[i].value))
	}
	return values
}

func (ik *Searcher) Analysis(word string) []string {
	values := make([]string, 0)
	search := ik.tree.Analysis([]rune(word))
	for i := range search {
		values = append(values, string(search[i].value))
	}
	return values
}