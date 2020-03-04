package main

import (
	"fmt"
	model "go_dev/day6/work/example1/model"
)

//创建图书
func BookCreate(l *[]*model.Books) {
	var (
		name       string
		total      int
		author     string
		createtime string
	)
	fmt.Println("请输入书名,数量,作者,日期(用空格隔开):")
	fmt.Scanf("%s %d %s %s", &name, &total, &author, &createtime)
	fmt.Println("=================")
	//book1 := model.CreateBook(name, total, author, createtime)
	//fmt.Println(book1)
	// var book1 = new(model.Books)
	var book1 = model.Books{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createtime,
	}

	*l = append(*l, &book1)
	fmt.Println(l)

	return
}

//借书 还书方法
func CanBoook(l *[]*model.Books, n int) {
	k := *l
	//ListBooks(k)
	for i := 0; i < len(k); i++ {
		temp := k[i]
		fmt.Printf("编号:[%d],书名:[%s].", i, temp.Name)
	}
	var (
		ns    int
		count int
	)
	fmt.Scanf("选择书的编号和数量,以空格隔开:", &ns, &count)
	if n == 3 {
		u := k[ns].Borrow(count)
		if u != nil {
			fmt.Println(u)
		} else {
			fmt.Printf("书名:[%s],剩余数量:[%d]", k[ns].Name, k[ns].Total)
		}
	} else if n == 4 { //还书方法
		k[ns].Back(count)
	}
}

//图书管理方法
func BookMang(l *[]*model.Books) {
	var a int
	fmt.Println("欢迎使用GO编写的图书管理系统-图书管理员端")

	for {
		fmt.Println("请选择项目:创建书本-1,查看书本-2,借书-3,还书-4,创建学生-5,退出-6")
		fmt.Scanln(&a)
		if a == 1 {
			BookCreate(l)
			fmt.Println("------------------")
			// fmt.Println(m)
			a = 0
			continue
		} else if a == 6 {
			return
		} else if a == 2 {
			k := *l
			for i := 0; i < len(k); i++ {
				temp := k[i]
				fmt.Println(temp.Name)
			}
			a = 0
		} else if a == 3 {
			fmt.Println("借书")
			CanBoook(l, a)
			a = 0
		} else if a == 4 {
			fmt.Println("还书")
			CanBoook(l, a)
			a = 0
		} else if a == 5 {
			fmt.Println("创建学生!")
		} else {
			fmt.Println("请输入正确的选项!")
			continue
		}
	}

}

//学员界面
func stumag(l *[]*model.Books) {
	var a int
	fmt.Println("欢迎使用GO编写的图书管理系统-学生端")
	for {
		fmt.Println("请选择项目:借书-1,还书-2,退出-3")
		fmt.Scanln(&a)
		if a == 1 {

		}
	}

}

func main() {
	var (
		n        int
		ListBook []*model.Books
	)

	fmt.Println("欢迎使用GO语言编写程序!")
	fmt.Println("请选择界面:1-管理员界面,2-学生界面,3-退出.")
	fmt.Scanf("%d", &n)
	for {
		fmt.Println("请选择界面:1-管理员界面,2-学生界面,3-退出.")
		fmt.Scanf("%d", &n)
		if n == 1 {
			BookMang(&ListBook)
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
