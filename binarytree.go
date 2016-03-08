package main

import (
  "fmt"
)

type Node struct {
  value int
  left, right *Node
}

type Tree struct {
  root *Node
}

func (t *Tree) Insert(value int) {
  currentNode := t.root

  if currentNode == nil {
    t.root = &Node{value, nil, nil}
    return
  }

  for {
    if value > currentNode.value {
      if currentNode.right != nil {
        currentNode = currentNode.right
      } else {
        currentNode.right = &Node{value, nil, nil}
        return
      }
    } else {
      if currentNode.left != nil {
        currentNode = currentNode.left
      } else {
        currentNode.left = &Node{value, nil, nil}
        return
      }
    }
  }
}

func (t *Tree) Print() {
  if t.root == nil {
    fmt.Println("Tree is empty.")
    return
  }

  nodes := []*Node{t.root}

  var children []*Node

  for {
    if len(nodes) == 0 {
      fmt.Println()
      if len(children) == 0 {
        return
      }
      nodes = children
      children = []*Node{}
    }

    node_ptr := nodes[0]
    node := *node_ptr

    if len(nodes) > 1 {
      nodes = nodes[1:]
    } else {
      nodes = []*Node{}
    }

    fmt.Printf("(%d)", node.value)

    if node.left != nil {
      children = append(children, node.left)
    }

    if node.right != nil {
      children = append(children, node.right)
    }
  }
}

func main() {
  t := Tree{nil}
  t.Insert(5)
  t.Insert(4)
  t.Insert(6)
  t.Insert(3)
  t.Insert(7)
  t.Insert(1)
  t.Insert(0)
  t.Insert(2)
  t.Insert(8)
  t.Print()
}
