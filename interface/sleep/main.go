package main

import (
	"flag"
	"fmt"
	"time"
)

type Celsius float64
type celsiusFlag struct{ Celsius }

func (c Celsius) String() string {
	return fmt.Sprintf("%f", c)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		//f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var period *time.Duration = flag.Duration("period", 1*time.Second, "sleep period")
var c *Celsius = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	//返回*time.Duration指针 下main参数接受的是实际值
	//flag.Parse()
	//fmt.Printf("Sleeping for %v...", *period)
	//fmt.Println(period)
	//time.Sleep(*period)
	//fmt.Println()
	flag.Parse()
	fmt.Println(c)
}
