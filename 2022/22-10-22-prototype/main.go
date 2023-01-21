package main

import (
	"bytes"
	"fmt"
)

type Node interface {
	String() string
	Parent() Node
	SetParent(node Node)
	Children() []Node
	AddChild(child Node)
	Clone() Node
}

type Element struct {
	text     string
	parent   Node
	children []Node
}

func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]Node, 0),
	}
}

func (e *Element) Parent() Node {
	return e.parent
}

func (e *Element) SetParent(node Node) {
	e.parent = node
}

func (e *Element) Children() []Node {
	return e.children
}

func (e *Element) AddChild(child Node) {
	cp := child.Clone()
	cp.SetParent(e)
	e.children = append(e.children, cp)
}

func (e *Element) Clone() Node {
	cp := &Element{
		text:     e.text,
		parent:   nil,
		children: make([]Node, 0),
	}
	for _, child := range e.children {
		cp.AddChild(child)
	}
	return cp
}

func (e *Element) String() string {
	buf := bytes.NewBufferString(e.text)
	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buf, "\n %s", text)
	}
	return buf.String()
}

func main() {
	// 职级节点----总监
	directorNode := NewElement("Director of Engineering")
	// 职级节点----研发经理
	engManagerNode := NewElement("Engineering Manager")
	engManagerNode.AddChild(NewElement("Lead Software Engineer"))

	// 研发经理是总监的下级
	directorNode.AddChild(engManagerNode)
	directorNode.AddChild(engManagerNode)

	// 办公室经理也是总监的下级
	officeManagerNode := NewElement("Office Manager")
	directorNode.AddChild(officeManagerNode)
	println()
	println("# Company Hierarchy")
	fmt.Print(directorNode)
	println()

	// 从研发经理节点可溶出一颗新的树
	println("# Team Hierarchy")
	fmt.Print(engManagerNode.Clone())
}
