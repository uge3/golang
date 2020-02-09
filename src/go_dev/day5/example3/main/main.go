package main

import "fmt"

//二叉树结构体
type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

//遍历二叉树
func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	//fmt.Println(root)
	trans(root.right)
}

func main() {
	var root *Student = new(Student)
	root.Name = "root"
	root.Age = 39
	root.Score = 99.87
	fmt.Println(root)
	var left1 *Student = new(Student)
	left1.Name = "left1"
	left1.Age = 39
	left1.Score = 99.87

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "right1"
	right1.Age = 39
	right1.Score = 99.87

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "left2"
	left2.Age = 39
	left2.Score = 99.87

	left1.left = left2

	var left3 *Student = new(Student)
	left3.Name = "left3"
	left3.Age = 39
	left3.Score = 99.87

	left2.left = left3

	trans(root)
}
