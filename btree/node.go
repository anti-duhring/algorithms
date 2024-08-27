package main

import "bytes"

const (
	// Higher degree means more keys can be stored in a node
	// degree (often denoted as m) is the min number of children a non-leaf node can point to
	degree      = 5
	maxChildren = degree * 2
	maxKeys     = maxChildren - 1
	minKeys     = degree - 1
)

type item struct {
	key   []byte
	value []byte
}

type node struct {
	// using fixed-size arrays instead of slices to avoid costly slice expasion operations during insertions
	// ref: https://go.dev/blog/slices
	items         [maxKeys]*item
	children      [maxChildren]*node
	itemsCount    int
	childrenCount int
}

func (n *node) isLeaf() bool {
	return n.childrenCount == 0
}

// binary search
// returns the index of the key and a boolean indicating if the key was found
// if the key was not found, the index returned is the index where the key should be
// we can use the returned position to continue traversing the tree until we eventually find a data item with the key
func (n *node) search(key []byte) (int, bool) {
	low, high := 0, n.itemsCount
	var mid int
	for low < high {
		mid = low + (high-low)/2
		cmp := bytes.Compare(n.items[mid].key, key)
		if cmp == 0 {
			return mid, true
		}
		if cmp < 0 {
			high = mid
		}
		if cmp > 0 {
			low = mid + 1
		}
	}
	return low, false
}

func (n *node) insertItemAt(pos int, i *item) {
	if pos < n.itemsCount {
		// make room for the new item if it's not being inserted at the end
		copy(n.items[pos+1:n.itemsCount+1], n.items[pos:n.itemsCount])
	}
	n.items[pos] = i
	n.itemsCount++
}

func (n *node) insertChildAt(pos int, c *node) {
	if pos < n.childrenCount {
		// make room for the new child if it's not being inserted at the end
		copy(n.children[pos+1:n.childrenCount+1], n.children[pos:n.childrenCount])
	}
	n.children[pos] = c
	n.childrenCount++
}

func (n *node) split() (*item, *node) {
	mid := minKeys
	midItem := n.items[mid]

	newNode := &node{}
	copy(newNode.items[:], n.items[mid+1:])

	if !n.isLeaf() {
		copy(newNode.children[:], n.children[mid+1:])
		newNode.childrenCount = minKeys + 1
	}

	// Remove data items and child pointers from the current node that were moved to the new node
	for i, l := mid, n.itemsCount; i < l; i++ {
		n.items[i] = nil
		n.itemsCount--

		if !n.isLeaf() {
			n.children[i+1] = nil
			n.childrenCount--
		}
	}

	return midItem, newNode
}

func (n *node) insert(item *item) bool {
	pos, found := n.search(item.key)

	// data exists, update it
	if found {
		n.items[pos] = item
		return false
	}

	// leaf node so theres room to insert
	if n.isLeaf() {
		n.insertItemAt(pos, item)
		return true
	}

	if n.children[pos].itemsCount >= maxKeys {
		midItem, midNode := n.children[pos].split()
		n.insertItemAt(pos, midItem)
		n.insertChildAt(pos+1, midNode)

		cmp := bytes.Compare(item.key, midItem.key)
		if cmp > 0 {
			pos++
		}
		if cmp == 0 {
			n.items[pos] = item
			return false
		}
	}

	return n.children[pos].insert(item)
}

func (n *node) removeItemAt(pos int) *item {
	removedItem := n.items[pos]
	n.items[pos] = nil

	// fill the gap left if the item was not the last one
	if lastPost := n.itemsCount - 1; pos < lastPost {
		copy(n.items[pos:lastPost], n.items[pos+1:])
		n.items[lastPost] = nil
	}
	n.itemsCount--

	return removedItem
}

func (n *node) removeChildAt(pos int) *node {
	removedChild := n.children[pos]
	n.children[pos] = nil

	// fill the gap left if the child was not the last one
	if lastPost := n.childrenCount - 1; pos < lastPost {
		copy(n.children[pos:lastPost], n.children[pos+1:])
		n.children[lastPost] = nil
	}
	n.childrenCount--

	return removedChild
}

func (n *node) borrowFromPrev(pos int) {
	left, right := n.children[pos-1], n.children[pos]
	// take the item from the parent and insert it at the left-most position of the right node
	copy(right.items[1:right.itemsCount+1], right.items[:right.itemsCount])
	right.items[0] = n.items[pos-1]
	right.itemsCount++

	// for non-leaf nodes make the right-most child of the left node
	// the new left-most child of the right node
	if !right.isLeaf() {
		right.insertChildAt(0, left.removeChildAt(left.childrenCount-1))
	}

	// borrow the right-most item from the left node and replace the parent item
	n.items[pos-1] = left.removeItemAt(left.itemsCount - 1)
}

func (n *node) borrowFromNext(pos int) {
	left, right := n.children[pos], n.children[pos+1]
	// take the item from the parent and insert it at the right-most position of the left node
	left.items[left.itemsCount] = n.items[pos]
	left.itemsCount++
	// for non-leaf nodes make the left-most child of the right node
	// the new right-most child of the left node
	if !left.isLeaf() {
		left.insertChildAt(left.childrenCount, right.removeChildAt(0))
	}

	// borrow the left-most item from the right node and replace the parent item
	n.items[pos] = right.removeItemAt(0)
}

// borrow and merge
func (n *node) fillChildAt(pos int) {
	// borrow the right-most item from the left sibling if possible
	if pos > 0 && n.children[pos-1].itemsCount > minKeys {
		n.borrowFromPrev(pos)
		return
	}
	// borrow the left-most item from the right sibling if possible
	if pos < n.childrenCount-1 && n.children[pos+1].itemsCount > minKeys {
		n.borrowFromNext(pos)
		return
	}

	// there are no nodes to borrow items from, so merge
	// if we are at the right-most child, merge with the left sibling
	// otherwise merge with the right sibling just to keep it simple
	if pos >= n.itemsCount {
		pos = n.itemsCount - 1
	}

	left, right := n.children[pos], n.children[pos+1]
	// borrow an item from parent node and place it at the right-most position of the left node
	left.items[left.itemsCount] = n.items[pos]
	left.itemsCount++
	// copy all items from the right node to the left node
	copy(left.items[left.itemsCount:], right.items[:right.itemsCount])
	// for non-leaf nodes copy all children from the right node to the left node
	if !left.isLeaf() {
		copy(left.children[left.childrenCount:], right.children[:right.childrenCount])
		left.childrenCount += right.childrenCount
	}

	// remove the child pointer from the parent to the right node and discard the right node
	n.removeChildAt(pos + 1)
	right = nil
}

func (n *node) delete(key []byte, isSeekingSuccessor bool) *item {
	pos, found := n.search(key)
	var next *node

	// we found a node with the key
	if found {
		// if the node is a leaf node, we can just remove the item
		if n.isLeaf() {
			return n.removeItemAt(pos)
		}

		// not a leaf node, we need to find the in-order successor
		next, isSeekingSuccessor = n.children[pos+1], true
	} else {
		// the key was not found, continue searching in the child node
		next = n.children[pos]
	}

	// reached the leaf node containing the inoder successor, so remove this successor from the leaf
	if next.isLeaf() && isSeekingSuccessor {
		return next.removeItemAt(0)
	}

	// item not found
	if next == nil {
		return nil
	}

	deletedItem := next.delete(key, isSeekingSuccessor)

	// found the in-order successor
	// replace the item in the current node with the in-order successor
	if found && isSeekingSuccessor {
		n.items[pos] = deletedItem
	}

	// check if underflow occurred after deletion
	if next.itemsCount < minKeys {
		if found && isSeekingSuccessor {
			next.fillChildAt(pos + 1)
		} else {
			n.fillChildAt(pos)
		}
	}

	// propagate the deleted item back to the previous stack frame
	return deletedItem
}
