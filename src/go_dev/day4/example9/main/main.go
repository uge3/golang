package main

import (
	"fmt"
	"sort"
)

func testSort() {
	s := [...]int{1, 8, 656, 999, 5, 7, 6}
	fmt.Println("排序前:", s)
	sort.Ints(s[:])
	fmt.Println("排序后:", s)
}
func testSortf() {
	s := [...]float64{1.56, 8.25, 656.568, 999.254, 5.987, 7.03, 6.007}
	fmt.Println("排序前:", s)
	sort.Float64s(s[:])
	fmt.Println("排序后:", s)
}
func testSortstinge() {
	s := [...]string{"abc", "oop", "tt", "bb", "zzz", "ASD", "UUU"}
	fmt.Println("排序前:", s)
	sort.Strings(s[:])
	fmt.Println("排序后:", s)
}

func testSortsting() {
	s := [...]string{"中", "国", "人", "无", "夺", "柘城", "哦哦地方"}
	fmt.Println("排序前:", s)
	sort.Strings(s[:])
	fmt.Println("排序后:", s)
}

//查找
func testSearch() {
	s := [...]int{1, 8, 656, 999, 5, 7, 6}
	fmt.Println("排序前:", s)
	sort.Ints(s[:])
	fmt.Println("排序后:", s)
	index := sort.SearchInts(s[:], 7)
	fmt.Println(index)
}

func main() {
	testSortstinge()
	testSortsting()
	testSort()
	testSortf()
	testSearch()
}
