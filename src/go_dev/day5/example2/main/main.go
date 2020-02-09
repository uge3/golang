package main

import (
	"fmt"
	"math/rand"
)

//结构体-链表
type Stuedt struct {
	Name  string
	Age   int
	Score float32
	next  *Stuedt
}

//循环打印
func trans(p *Stuedt) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

//尾部插入法
func insertTail(p *Stuedt) {
	var tail = p
	for i := 0; i <= 10; i++ {
		var temp = Stuedt{ //
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		// temp := &Stuedt{
		// 	Name:  fmt.Sprintf("stu%d", i),
		// 	Age:   rand.Intn(100),
		// 	Score: rand.Float32() * 100,
		// }
		// tail.next = temp
		// tail = temp
		tail.next = &temp
		tail = &temp
	}
}

//头部插入
func insertHead(head **Stuedt) {
	for i := 0; i <= 10; i++ {
		var temp = Stuedt{ //
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		temp.next = *head
		*head = &temp

	}
}

//删除节点  不完善,无法删除头节点
func delTail(p *Stuedt, name string) {
	var perv *Stuedt = p
	for p.next != nil {
		if p.Name == name {
			perv.next = p.next
			break
		}
		perv = p
		p = p.next
	}
}

//链表中增加插入节点      新节点指针      要插入节点定位
func addNode(p *Stuedt, newNode *Stuedt, name string) {
	for p != nil {
		if p.Name == name {
			newNode.next = p.next //新节点下一个指向,当前节点的指向
			p.next = newNode      //当前节点指向这个新节点
			break
		}
		p = p.next
	}
}

func main() {
	var stu Stuedt
	stu.Name = "颜靖靖"
	stu.Age = 39
	stu.Score = 99.87

	var stu2 Stuedt
	stu2.Age = 40
	stu2.Name = "大规模"
	stu2.Score = 88
	stu2.next = nil
	stu.next = &stu2
	trans(&stu)
	insertTail(&stu) //尾部插入

	trans(&stu) //循环打印
	var head *Stuedt = new(Stuedt)
	head.Name = "g"
	head.Age = 17
	head.Score = 99
	insertHead(&head) //头部插入
	fmt.Println("---------------------")
	trans(head)
	fmt.Println("---------------------")
	delTail(head, "stu6") //删除一个
	trans(head)
	fmt.Println("---------------------")

	addNode(head, &stu2, "g") //插入一下节点
	trans(head)
}
