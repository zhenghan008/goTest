package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type List interface {
	NewList(data ...interface{})
	PrintList()
	SetList(index int, value interface{})
}

type List1 struct {
	listPtr []interface{}
}

type Cat struct {
	Name string
	Age  int
}

func (l *List1) NewList(data ...interface{}) {
	for i := 0; i < len(data); i++ {
		l.listPtr = append(l.listPtr, &data[i])

	}
}

func (l *List1) PrintList() {
	for _, value := range l.listPtr {
		fmt.Println(reflect.ValueOf(value).Elem())

	}
}

func (l *List1) SetList(index int, value interface{}) {
	l.listPtr[index] = &value

}

func main() {
	//var l1 List = &List1{}
	//a := "123"
	//b := 1
	//c := Cat{"cat", 1}
	//d := 10.23
	//l1.NewList(a, b, c)
	//l1.PrintList()
	//l1.SetList(1, d)
	//l1.PrintList()
	a := []string{"Add", "Sub"}
	c := make(map[string]int, 0)
	for i := 0; i <= 100; i++ {
		var i Int = &MyInt{
			self: rand.Intn(10),
		}
		callFun := reflect.ValueOf(i).MethodByName(a[rand.Intn(2)])
		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(rand.Intn(10))
		s := fmt.Sprintf("%s", callFun.Call(params)[0])
		c[s] = 1
		//fmt.Println(reflect.ValueOf(i).MethodByName(a[rand.Intn(2)]).Call([]reflect.Value{})[0])
	}
	for key, _ := range c {
		fmt.Println(key)
	}

}
