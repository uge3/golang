package model

import "errors"

//定义各类提示
var (
	ErrStockNotEnough = errors.New("库存不足!")
)

//书籍 结构体
type Books struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

//书本创建方法
func CreateBook(name string, total int, author string, createtime string) (b *Books) {
	*b = Books{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createtime,
	}
	return b
}

//结构方法-借书数量判断
func (b *Books) canBorrow(c int) bool {
	return b.Total >= c
}

//接口-借书
func (b *Books) Borrow(c int) (err error) {
	if b.canBorrow(c) == false {
		err = ErrStockNotEnough
		return
	}
	b.Total -= c //数量充足减去借出数量
	return
}

//接口-还书
func (b *Books) Back(c int) {

	b.Total += c //加上还回的数量
	return
}
