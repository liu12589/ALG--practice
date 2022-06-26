package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxLevel    = 16
	maxRand     = 65535.0
	ZSKIPLIST_P = 0.25
)

type SkipList struct {
	header *SkipListNode
	tail   *SkipListNode
	length int64
	level  int
}
type SkipListNode struct {
	ele      string //实际上使用sds数据结构
	score    float64
	backward *SkipListNode
	level    []SkipListLevel
}
type SkipListLevel struct {
	forward *SkipListNode
	span    int64
}

func main() {
	fruit := createSkipList()
	fruit.Add(10, "苹果")
	fruit.Add(5, "香蕉")
	fruit.Add(1, "西瓜")
	fmt.Println(fruit.Search(5, "香蕉"))
}

// 创造一个跳表节点
func createNode(level int, score float64, ele string) *SkipListNode {
	var arr = make([]SkipListLevel, level)
	return &SkipListNode{
		level: arr,
		score: score,
		ele:   ele,
	}
}

// Constructor 初始化一个跳表
func createSkipList() *SkipList {
	var zsl *SkipList
	zsl = new(SkipList)
	// 注意要分配内存后才能够使用
	zsl.level = 1
	zsl.length = 0

	//创建一个null的zslCreateNode 跳表节点，作为header，后续的节点加入都在null之后
	zsl.header = createNode(maxLevel, 0, "")
	for i := 0; i < maxLevel; i++ {
		zsl.header.level[i].forward = nil
		zsl.header.level[i].span = 0
	}
	zsl.header.backward = nil
	zsl.tail = nil
	return zsl
}

func zslRandomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for rand.Float64() < ZSKIPLIST_P && level < maxLevel {
		level++
	}
	return level
}

func (zsl *SkipList) Add(score float64, ele string) {
	// 创造一个update和rank辅助数组，负责处理跳表结构
	var update = make([]*SkipListNode, maxLevel)
	var x *SkipListNode
	var rank = make([]int64, maxLevel)
	var level int

	x = zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		if i == zsl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					strings.Compare(x.level[i].forward.ele, ele) < 0)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}
	level = zslRandomLevel()
	// 如果得到的随机level比现在的level最大层高还要大的话，需要设置最大层为最新level
	if level > zsl.level {
		for i := zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = zsl.length
		}
		zsl.level = level
	}

	//初始化新的节点，开始插入操作
	x = createNode(level, score, ele)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		//根据update数组更新span跨度
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = rank[0] - rank[i] + 1
	}
	for i := level; i < zsl.level; i++ {
		update[i].level[i].span++
	}
	if update[0] == zsl.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		zsl.tail = x
	}
	zsl.length++
}

func (zsl *SkipList) Search(score float64, ele string) int64 {
	var x *SkipListNode
	var rank int64 = 0
	x = zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					strings.Compare(x.level[i].forward.ele, ele) == 0)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}

		if x.ele != "" && x.score == score && strings.Compare(x.ele, ele) == 0 {
			return rank
		}
	}
	return 0
}

//
//func (zsl *SkipList) Delete(score float64, ele string) bool {
//
//}
