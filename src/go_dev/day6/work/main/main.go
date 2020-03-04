package main

import (
	"fmt"
	model "go_dev/day6/work/example1/model"
)

//创建图书
func BookCreate(m map[string]model.Books) {
	var (
		name       string
		total      int
		author     string
		createtime string
	)
	fmt.Println("请输入书名,数量,作者,日期(用空格隔开):")
	fmt.Scanf("%s %d %s %s", &name, &total, &author, &createtime)
	fmt.Println("=================")
	book1 := model.CreateBook(name, total, author, createtime)
	fmt.Println(book1)
	m[name] = book1
	fmt.Println(m)
	return
}

//借书方法
func CanBoook(m map[string]model.Books) {
	var keys []string
	for k, _ := range m {
		//fmt.Println(k)
		keys = append(keys, k)
	}
	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i])
	}
}

//图书管理方法
func BookMang(m map[string]model.Books) {
	var a int
	fmt.Println("欢迎使用GO编写的图书管理系统-图书管理员端")

	for {
		fmt.Println("请选择项目:创建书本-1,查看书本-2,借书-3,还书-4,退出-5")
		fmt.Scanln(&a)
		if a == 1 {
			BookCreate(m)
			fmt.Println("------------------")
			fmt.Println(m)
			a = 0
			continue
		} else if a == 5 {
			return
		} else if a == 2 {
			fmt.Println(m)
			a = 0
		} else if a == 3 {
			fmt.Println("借书")
			CanBoook(m)
			a = 0
		} else if a == 4 {
			fmt.Println("还书")
			a = 0
		} else {
			fmt.Println("请输入正确的选项!")
			continue
		}

	}

}

func main() {
	var ListBooks []*model.Books
	var MM = new(model.Books)
	MM.Name = "好书名"
	MM.Total = 3
	MM.Author = "作者AA"
	fmt.Println("---->", MM)
	ListBooks = append(ListBooks, MM)
	fmt.Println("---ListBooks->", ListBooks)

	for i := 0; i < len(ListBooks); i++ {
		Book := ListBooks[i]
		fmt.Println("---Name->", Book.Name)
		fmt.Println("---Total->", Book.Total)
		fmt.Println("---Author->", Book.Author)
	}
	var (
		n int
		m map[string]model.Books
	)
	m = make(map[string]model.Books)
	fmt.Println("欢迎使用GO语言编写程序!")
	fmt.Println("请选择界面:1-管理员界面,2-学生界面,3-退出.")
	fmt.Scanf("%d", &n)
	for {
		fmt.Println("请选择界面:1-管理员界面,2-学生界面,3-退出.")
		fmt.Scanf("%d", &n)
		if n == 1 {
			BookMang(m)
		} else if n == 2 {
			fmt.Println("欢迎使用GO编写的图书管理系统-学生端")
		} else if n == 3 {
			break
		} else {
			fmt.Println("请选择正确的选项!")
			n = 0
			continue
		}

	}

}
