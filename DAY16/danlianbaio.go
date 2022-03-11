package main

import "fmt"

type object interface{}
type node struct {
	Data object
	Next *node
}
type List struct {
	HeadNode *node
	LastNode *node
}

func (list *List) addhead(p object) {
	var node = &node{}
	node.Data = p
	if list.HeadNode != nil {
		return
	} else {
		list.HeadNode = node
	}
}
func main() {
	var a = List{}
	a.addhead(13)
	fmt.Println(a.HeadNode)

}
