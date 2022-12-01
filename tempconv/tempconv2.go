package tempc

import "fmt"

//Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个名叫String的方法，该方法返回该类型对象c带着°C温度单位的字符串：
//c 为调用该方法的对象
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
func (ft Feet) String() string {
	return fmt.Sprintf("%gft", ft)
}
func (m Meter) String() string {
	return fmt.Sprintf("%g", m)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func MToFT(m Meter) Feet {
	return Feet(m * 3.2808398950131)
}
func FTToM(ft Feet) Meter {
	return Meter(0.3048 * ft)
}
func PToKG(p Pounds) KG {
	return KG(0.45359237 * p)
}
func KGToP(kg KG) Pounds {
	return Pounds(2.20462262 * kg)
}
