package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	tempc "studygolang-example/tempconv"
)

var m = flag.String("m", "a", "选择一个换算模式，比如：温度a,米英尺b,磅公斤c")
var v = flag.String("v", "", "需要换算的值")

func main() {
	flag.Parse()
	t, err := strconv.ParseFloat(*v, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf:%v\n", err)
		os.Exit(1)
	}
	//if len(os.Args) > 1 {
	//	for _, arg := range os.Args[1:] {
	//		t, err = strconv.ParseFloat(arg, 64)
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
	//			os.Exit(1)
	//		}
	//
	//	}
	//}
	//else {
	//	scaner := bufio.NewScanner(os.Stdin)
	//	if scaner.Scan() {
	//		t, err = strconv.ParseFloat(scaner.Text(), 64)
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
	//			os.Exit(1)
	//		}
	//
	//	}
	//}
	if *m == "a" {
		f := tempc.Fahrenheit(t)
		c := tempc.Celsius(t)
		fmt.Printf("%s = %s, %s = %s \n",
			f, tempc.FToC(f), c, tempc.CToF(c))
	} else if *m == "b" {
		m := tempc.Meter(t)
		ft := tempc.Feet(t)
		fmt.Printf("%s = %s,%s = %s \n",
			m, tempc.MToFT(m), ft, tempc.FTToM(ft))
	} else if *m == "c" {
		p := tempc.Pounds(t)
		kg := tempc.KG(t)
		fmt.Printf("%v = %v,%v = %v \n",
			p, tempc.PToKG(p), kg, tempc.KGToP(kg))
	} else {
		fmt.Fprintf(os.Stderr, "not support")
	}

}
