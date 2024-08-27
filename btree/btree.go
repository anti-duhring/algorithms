package main

import "fmt"

type BTree struct {
	root *node
}

func NewBTree() *BTree {
	return &BTree{root: &node{}}
}

func (t *BTree) Find(key []byte) ([]byte, error) {
	for next := t.root; next != nil; {
		pos, found := next.search(key)
		if found {
			return next.items[pos].value, nil
		}
		if next.isLeaf() {
			break
		}
		next = next.children[pos]
	}

	return nil, fmt.Errorf("key not found")
}

func (t *BTree) splitRoot() {
	newRoot := &node{}
	midItem, newNode := t.root.split()
	newRoot.insertItemAt(0, midItem)
	newRoot.insertChildAt(0, t.root)
	newRoot.insertChildAt(1, newNode)
	t.root = newRoot
}

func (t *BTree) Insert(key, value []byte) {
	i := &item{key, value}

	// tree is empty, create a new root
	if t.root == nil {
		t.root = &node{}
	}

	// root is full, split it
	if t.root.itemsCount >= maxKeys {
		t.splitRoot()
	}

	t.root.insert(i)
}

func (t *BTree) Delete(key []byte) bool {
	if t.root == nil {
		return false
	}
	deletedItem := t.root.delete(key, false)

	// root is empty, remove it
	if t.root.itemsCount == 0 {
		if t.root.isLeaf() {
			t.root = nil
		} else {
			t.root = t.root.children[0]
		}
	}

	if deletedItem != nil {
		return true
	}
	return false
}
