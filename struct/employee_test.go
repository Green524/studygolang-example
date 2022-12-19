package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmp(t *testing.T) {
	fmt.Printf("%p%v\n", &dilbert, dilbert)
	dilbert.Salary -= 5000
	fmt.Printf("%p%v\n", &dilbert, dilbert)
	position := &dilbert.Position
	fmt.Println(reflect.TypeOf(position)) //*string
	*position = "Senior " + *position
	fmt.Printf("%p%v\n", &dilbert, dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	//等价于 (*employeeOfTheMonth).Position += "(proactive team player)"
	fmt.Printf("%p%v\n", &dilbert, dilbert)

	id := dilbert.ID
	a := EmployeeByID(id)
	fmt.Printf("%p%v\n", &a, a)
	b := &dilbert
	fmt.Printf("%p%v\n", &b, dilbert)

	//为什么golang的结构体成员中不能包含自己？但是可以包含自己的地址 *s
}
