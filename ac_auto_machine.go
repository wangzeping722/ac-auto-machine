package ac_auto_machine

import (
	"unicode/utf8"
)

type AcAutoMachine struct {
	root *acNode
}

type acNode struct {
	child      map[rune]*acNode
	r          rune
	existWords map[int]bool
	failPtr    *acNode
	father     *acNode
}

func NewAcAutoMachine() *AcAutoMachine {
	return &AcAutoMachine{
		root: &acNode{
			child:      map[rune]*acNode{},
			existWords: map[int]bool{},
		},
	}
}

func (acm *AcAutoMachine) AddWord(word string) {
	curNode := acm.root
	for _, w := range word {
		if _, ok := curNode.child[w]; !ok {
			curNode.child[w] = &acNode{
				r:          w,
				father:     curNode,
				child:      map[rune]*acNode{},
				existWords: map[int]bool{},
			}
		}
		curNode = curNode.child[w]
	}

	curNode.existWords[utf8.RuneCountInString(word)] = true
}

func (acm *AcAutoMachine) Build() {
	root := acm.root
	queue := make([]*acNode, 0)
	queue = append(queue, root)
	// bfs
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		for _, child := range curNode.child {
			queue = append(queue, child)
		}
		if curNode == root {
			continue
		}
		if curNode.father == root {
			curNode.failPtr = root
			continue
		}
		if failNode, ok := curNode.father.failPtr.child[curNode.r]; ok {
			curNode.failPtr = failNode
			for k := range failNode.existWords {
				curNode.existWords[k] = true
			}
		} else {
			curNode.failPtr = root
		}
	}
}

func (acm *AcAutoMachine) Search(word string) []string {
	root := acm.root
	curNode := acm.root
	matched := make(map[string]bool)
	for i, r := range word {
		if node, ok := curNode.child[r]; !ok {
			if curNode.failPtr == nil {
				curNode = root
				continue
			}
			// 如果 child 不存在,那么说明需要利用 failptr 进行回退, 直到回退到 root 节点, 或者找到 child 为止
			for curNode != root{
				curNode = curNode.failPtr
				if curNode.child[r] != nil {
					curNode = curNode.child[r]
					break
				}
			}
		} else {
			curNode = node
			if len(curNode.existWords) != 0 {
				// 如果存在 existWords, 那么 matched
				for l := range curNode.existWords {
					matched[word[i-l+1:i+1]] = true
				}
			}
		}
	}
	res := make([]string, 0)
	for s := range matched {
		res = append(res, s)
	}
	return res
}
