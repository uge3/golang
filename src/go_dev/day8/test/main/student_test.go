package main

import (
	"testing"
	"time"
)

//单元测试
func TestSave(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}
	err := stu.Save()
	if err != nil {
		t.Fatalf("save studnet error%v", err)
	}
}

func TestLoad(t *testing.T) {
	stu := &student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}
	err := stu.Save()
	if err != nil {
		t.Fatalf("save studnet error%v", err)
	}
	stu2 := &student{}
	time.Sleep(time.Second * 10)
	err2 := stu2.Load()
	if err2 != nil {
		t.Fatalf("load student failed,error%v", err)
	}
	if stu.Name != stu2.Name {
		t.Fatalf("laod student failed,error name")
	}
	if stu.Sex != stu2.Sex {
		t.Fatalf("laod student failed,error sex")
	}
	if stu.Age != stu2.Age {
		t.Fatalf("laod student failed,error age")
	}
}
