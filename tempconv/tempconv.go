package main

import (
	"fmt"
	"reflect"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = 273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("华氏温度：%f\t", CToF(1))
	fmt.Printf("摄氏温度：%f\n", FToC(70))
	//显式类型转换将AbsoluteZeroC转换为Fahrenheit类型
	a := Fahrenheit(AbsoluteZeroC)
	fmt.Println(a)

	//将指针类型转换,需要用小括弧包裹(T)(*b)
	b := FreezingC
	fmt.Println(reflect.TypeOf(&b))
	p := (*Fahrenheit)(&b)
	fmt.Println(reflect.TypeOf(p))

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	//报错 类型不同无法相减
	//fmt.Printf("%g\n", boilingF-FreezingC)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) //true
	fmt.Println(f == 0) // true
	//fmt.Println(c == f) //报错
	fmt.Println(c == Celsius(f)) //true 类型转换不会改变值

	fmt.Println("--------------")
	//许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法
	d := FToC(212.0)
	fmt.Println(d.String()) // "100°C"
	fmt.Printf("%v\n", d)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", d)   // "100°C"
	fmt.Println(d)          // "100°C"
	fmt.Printf("%g\n", d)   // "100"; does not call String
	fmt.Println(float64(d)) // "100"; does not call String
}

//Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个名叫String的方法，该方法返回该类型对象c带着°C温度单位的字符串：
//c 为调用该方法的对象
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
