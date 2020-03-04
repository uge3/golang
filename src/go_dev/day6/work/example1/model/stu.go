package model

import "errors"

//定义各类提示
var (
	ErrNotFontEnough = errors.New("书名出错!")
)

//学生结构体
type Student struct {
	Name  string
	Age   int
	Id    string
	Grand string
	Sex   string
	Book  []*BorrowItem
}

//学生借书结构体  书名  数量
type BorrowItem struct {
	Book *Books
	num  int
}

//创建学生
func CreateStu(name string, age int, id string, grand string, sex string) *Student {
	stu := &Student{
		Name:  name,
		Age:   age,
		Id:    id,
		Grand: grand,
		Sex:   sex,
	}
	return stu
}

//学生借书
func (s *Student) AddBook(b *BorrowItem) {
	s.Book = append(s.Book, b) //借书列表进行添加
}

//学生还书
func (s *Student) DelBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.Book); i++ { //循环借书列表
		if s.Book[i].Book.Name == b.Book.Name { //查询到书名
			if b.num == s.Book[i].num { //所还数量如果相等
				front := s.Book[0:i]           //所还书本的左边
				left := s.Book[i+1:]           //所还书本的右边
				front = append(front, left...) //更新借书的列表
				s.Book = front
				return
			}
			s.Book[i].num -= b.num //还部分,直接减去书本的数量
			return
		}

	}
	err = ErrNotFontEnough
	return
}

//学生借书的列表
func (s *Student) GrandList() []*BorrowItem {
	return s.Book
}
