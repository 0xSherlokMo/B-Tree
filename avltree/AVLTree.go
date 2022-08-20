package avltree

import "github.com/0xSherlokMo/B-Tree/helpers"

type AVLNode struct {
	key int

	depth int
	left  *AVLNode
	right *AVLNode
}

func (n *AVLNode) add(key int) *AVLNode {
	if n == nil {
		return &AVLNode{key, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key)
	} else if key > n.key {
		n.right = n.right.add(key)
	} else {
		return n
	}

	return n.rebalance()
}

func (n *AVLNode) rebalance() *AVLNode {
	if n == nil {
		return nil
	}
	n.updateDepth()
	balance := n.balanceFactor()

	if balance > 1 {
		n.rotateRight()
	} else if balance < -1 {
		n.rotateLeft()
	}

	return n
}

func (n *AVLNode) rotateLeft() *AVLNode {
	temp := n.right
	n.right = temp.right
	temp.right = n

	n.updateDepth()
	temp.updateDepth()

	return temp
}

func (n *AVLNode) rotateRight() *AVLNode {

	temp := n.left
	n.left = temp.right
	temp.right = n

	n.updateDepth()
	temp.updateDepth()

	return temp
}

func (n *AVLNode) balanceFactor() int {
	return n.left.getDepth() - n.right.getDepth()
}

func (n *AVLNode) updateDepth() {
	n.depth = 1 + helpers.Max(n.left.getDepth(), n.right.getDepth())
}

func (n *AVLNode) getDepth() int {
	return n.depth
}
