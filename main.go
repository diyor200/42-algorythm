package main

import "fmt"

// [[],[1],[2],[1],[3],[2],[2],[2],[2]]
func main() {
	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(1))
	fmt.Println(myHashSet.Contains(3))
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(2))
	myHashSet.Remove(2)
	fmt.Println(myHashSet.Contains(2))

}

type Node struct {
	Value int
	Next  *Node
}

type MyHashSet struct {
	Head *Node
}

func Constructor() *MyHashSet {
	return &MyHashSet{Head: nil}
}

func (mh *MyHashSet) Add(key int) {
	if mh.isEmpty() {
		node := &Node{Value: key}
		mh.Head = node
		return
	}

	if mh.find(key) == nil {
		node := &Node{Value: key, Next: mh.Head}
		mh.Head = node
	}

	return
}

func (mh *MyHashSet) Remove(key int) {
	if mh.isEmpty() {
		return
	}

	if mh.Head.Value == key {
		mh.Head = mh.Head.Next
		return
	}

	node := mh.Head
	for node.Next != nil {
		if node.Next.Value == key {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
}

func (mh *MyHashSet) Contains(key int) bool {
	if mh.isEmpty() {
		return false
	}

	res := mh.find(key)
	return res != nil
}

func (mh *MyHashSet) find(value int) *Node {
	node := mh.Head
	for node != nil {
		if node.Value == value {
			return node
		}
		node = node.Next
	}

	return nil
}

func (mh *MyHashSet) isEmpty() bool {
	return mh.Head == nil
}
