package utils

import "io"

type Operation func(tree Node, items ...Node) bool

type AcceptedValues interface {
	float32 | float64 | string | int | int32 | int64 | int8
}

type Node interface {
	GetValue() any
	HasChildren() bool
	GetChildren() []Node
	AddChild(item Node, operation Operation) Node
	DeleteChild(item Node, operation Operation) bool
	ReplaceChild(newChildren []Node, operation Operation) bool
	Print(writer io.Writer)
}

type TreeNode[T AcceptedValues] struct {
	Val T
}

func (t *TreeNode[T]) GetValue() any {
	return t.Val
}

func (t *TreeNode[T]) ToTree() *Tree {
	return &Tree{Val: t.GetValue(), Children: []Node{}}
}

type Tree struct {
	Val      any
	Children []Node
}

func (t *Tree) GetValue() any {
	if t.Val == nil {
		return nil
	}
	return t.Val
}

func (t *Tree) HasChildren() bool {
	return len(t.Children) > 0
}

func (t *Tree) GetChildren() []Node {
	return t.Children
}

func (t *Tree) AddChild(item Node, operation Operation) Node {
	if operation == nil {
		t.Children = append(t.Children, item)
	} else {
		operation(t, item)
	}
	return t
}

func (t *Tree) ReplaceChild(newChildren []Node, operation Operation) bool {
	if operation == nil {
		t.Children = newChildren
		return true
	}
	return operation(t, newChildren...)
}

func (t *Tree) DeleteChild(item Node, operation Operation) bool {
	if operation == nil {
		return findAndDelete(t, item)
	}
	return operation(t, item)
}

func findAndDelete(parent Node, item Node) bool {
	if !parent.HasChildren() {
		return false
	}

	children := parent.GetChildren()
	for pos, val := range children {

		if val == item {
			parent.ReplaceChild(append(children[:pos], children[pos+1:]...), nil)
			return true
		}

		if parentChild, ok := val.(Node); ok {
			if parentChild.HasChildren() && findAndDelete(parentChild, item) {
				return true
			}
		}
	}

	return false
}

func (t *Tree) Print(writer io.Writer) {
	PrintNode(writer, t, "", true, true)
}
