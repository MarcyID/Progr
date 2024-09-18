package main

import (
	"fmt"
	"time"
)

func main() {
	datetime()
	//variables()
	//shortVariables()
	//intOperations()
	//floatOperations()
	//average()
}

func datetime() {
	time := time.Now()
	format := time.Format("02.01.2006 15:04:05")
	fmt.Println("Дата и время: ", format)
}

func variables() {
	var a int = 4
	var b float64 = 4.62432
	var c string = "Строка"
	var d bool = true
	fmt.Println(a, b, c, d)
}

func shortVariables() {
	a := 4
	b := 4.62432
	c := "Строка"
	d := true
	fmt.Println(a, b, c, d)
}

func intOperations() {
	var x int
	var y int
	fmt.Println("x=")
	fmt.Scanf("%d", &x)
	fmt.Println("y=")
	fmt.Scanf("%d", &y)
	fmt.Println("x+y=", x+y)
	fmt.Println("x-y=", x-y)
	fmt.Println("x*y=", x*y)
	fmt.Println("x/y=", x/y)
}

func floatOperations() {
	var x float64
	var y float64
	fmt.Println("x=")
	fmt.Scanf("%f", &x)
	fmt.Println("y=")
	fmt.Scanf("%f", &y)
	fmt.Println("x+y=", x+y)
	fmt.Println("x-y=", x-y)
}

func average() {
	var x, y, z float64
	fmt.Println("x=")
	fmt.Scanf("%f", &x)
	fmt.Println("y=")
	fmt.Scanf("%f", &y)
	fmt.Println("z=")
	fmt.Scanf("%f", &z)
	fmt.Printf("Среднее значение: %.3f", ((x + y + z) / 3))
}
