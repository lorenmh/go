package main

import (
  "fmt"
)

// binary tree node
type Node struct {
  value int
  left, right *Node
}

// binary tree
type Tree struct {
  root *Node
}

// insert an int into the tree
func (t *Tree) Insert(value int) {
  currentNode := t.root

  // if root is nil, insert the value at the root
  if currentNode == nil {
    t.root = &Node{value, nil, nil}
    return
  }

  // find the correct spot to insert the value and insert it
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

func (t *Tree) Find(value int) *Node {
  currentNode := t.root

  for {
    if currentNode == nil {
      return nil
    } else if currentNode.value == value {
      return currentNode
    }

    if value > currentNode.value {
      currentNode = currentNode.right
    } else {
      currentNode = currentNode.left
    }
  }
}

// Breadth first print
// using slices as a queue
func (t *Tree) Print() {
  if t.root == nil {
    fmt.Println("Tree is empty.")
    return
  }

  // initialize nodes slice with root
  nodes := []*Node{t.root}

  // declare children slice
  var children []*Node

  for {
    // take first element from nodes slice
    node := nodes[0]

    // print out this node's value
    fmt.Printf("(v: %d", node.value)

    // print out left child's value
    if node.left != nil {
      fmt.Printf(", l: %d", node.left.value)
    }

    // print out right child's value
    if node.right != nil {
      fmt.Printf(", r: %d", node.right.value)
    }

    // close parens
    fmt.Print(")")

    // append left child to childrens slice
    if node.left != nil {
      children = append(children, node.left)
    }

    // append right child to childrens slice
    if node.right != nil {
      children = append(children, node.right)
    }

    // if len slice gt 1, perform shift operation on slice
    if len(nodes) > 1 {
      nodes = nodes[1:]
    } else {
      // else, nodes is empty, print new line
      fmt.Println()

      // if children is empty, we have nothing left to print, return
      if len(children) == 0 {
        return
      }

      // move to the next row of children
      nodes = children
      children = []*Node{}
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
  fmt.Println(t.Find(2))
}
