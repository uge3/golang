package link

import "fmt"

//节点 结构体
type LinkNode struct {
	data interface{}
	next *LinkNode
}

//链表  结构体
type Link struct {
	head *LinkNode
	tail *LinkNode
}

//方法 链表 头部插入法
func (l *Link) IntsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if l.tail == nil && l.head == nil {
		l.tail = node
		l.head = node
		return
	}
	node.next = l.head
	l.head = node
}

//方法 链表 尾部插入法
func (l *Link) IntsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if l.tail == nil && l.head == nil {
		l.tail = node
		l.head = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *Link) Tarns() {
	p := l.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}

}
