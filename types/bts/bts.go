package bts

import (
	"github.com/stuffgo/libs/types"
)

type node[T types.Item[T]] struct {
	val   types.Item[T]
	left  *node[T]
	right *node[T]
}

func New[T types.Item[T]]() types.BTS[T] {
	return &node[T]{}
}

// Get get node value
func (tn *node[T]) Get() types.Item[T] {
	return tn.val
}

// GetMin - get minimal value
func (tn *node[T]) GetMin() types.Item[T] {
	if tn.left == nil {
		return tn.Get()
	}

	return tn.left.GetMin()
}

// GetMax - get maximum value
func (tn *node[T]) GetMax() types.Item[T] {
	if tn.right == nil {
		return tn.Get()
	}

	return tn.right.GetMax()
}

// Insert - insert value into bts
func (tn *node[T]) Insert(val types.Item[T]) {
	if tn == nil {
		*tn = node[T]{
			val: val,
		}
	}

	switch {
	case tn.val == nil:
		tn.val = val
	case tn.val.Equal(val):
		return
	case tn.val.Less(val):
		if tn.right == nil {
			tn.right = &node[T]{val: val}
			return
		}

		tn.right.Insert(val)
	case val.Less(tn.val):
		if tn.left == nil {
			tn.left = &node[T]{val: val}
			return
		}

		tn.left.Insert(val)
	default:
		tn.val = val
	}
}

// Delete - remove value and node fom bts
func (tn *node[T]) Delete(val types.Item[T]) {
	*tn = *tn.remove(val)
}

// Find - search value in tree
func (tn *node[T]) Find(val types.Item[T]) types.BTS[T] {
	if tn == nil {
		return nil
	}

	switch {
	case val.Equal(tn.val):
		return tn
	case val.Less(tn.val):
		return tn.Find(val)
	default:
		return tn.right.Find(val)
	}
}

// Len - returning length of bts
func (tn *node[T]) Len() int {
	l := 0

	tn.traverse(
		func(node *node[T]) {
			l++
		},
	)

	return l
}

// Iter - returning bts iterator
func (tn *node[T]) Iter() <-chan types.BTS[T] {
	iter := make(chan types.BTS[T])

	go func() {
		tn.traverse(
			func(node *node[T]) {
				iter <- node
			},
		)
		close(iter)
	}()

	return iter
}

func (tn *node[T]) traverse(traverseFunc func(node *node[T])) {
	if tn != nil {
		tn.left.traverse(traverseFunc)
		traverseFunc(tn)
		tn.right.traverse(traverseFunc)
	}
}

func (tn *node[T]) inorderShift() *node[T] {
	cur := tn
	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func (tn *node[T]) remove(val types.Item[T]) *node[T] {
	if tn == nil {
		return nil
	}

	switch {
	case val.Less(tn.val):
		tn.left = tn.left.remove(val)
	case tn.val.Less(val):
		tn.right = tn.right.remove(val)
	default:
		switch {
		case tn.left == nil:
			return tn.right
		case tn.right == nil:
			return tn.left
		default:
			t := tn.right.inorderShift()
			t.left = tn.left

			return tn.right
		}
	}

	return tn
}
