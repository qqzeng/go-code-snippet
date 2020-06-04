package main

import "fmt"

//go:noinline
func Add(a int32, b int32) int32 {
	return a + b
}

type Calculator struct {
	val int32
}

//go:noinline
func (c Calculator) Value() int32 {
	return c.val
}

//go:noinline
func (c *Calculator) SetValue(val int32) {
	c.val = val
}

//go:noinline
func (c *Calculator) Add(incr int32) {
	c.val += incr
}

func main() {
	Add(10, 21) // 0. 函数调用

	cal := Calculator{val: 0}
	cal.Value()    // 1. 值类型变量调用值类型接收者的方法
	(&cal).Value() // 2. 指针类型变量调用值类型接收者的方法

	(&cal).SetValue(10)              // 3. 指针类型变量调用指针类型接收者的方法
	cal.SetValue(10)                 // 4. 值类型变量调用指针类型接收者的方法
	(*Calculator).SetValue(&cal, 10) // 5. 直接调用方法所对应的一个函数以调用对应的方法

	fmt.Println(Calculator.Value(cal)) // 10
}
