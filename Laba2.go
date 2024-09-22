package main

import (
	"fmt"
)

func main() {
	//even_Odd()

	//result := checkSimbol(0)
	//fmt.Println(result)

	//cycle()

	//result := lenghtString("Марина")
	//fmt.Println(result)

	//result := Rectangle{5, 5}
	//fmt.Println(areaRectangle(result))

	result := average(20, 10)
	fmt.Println(result)
}

func even_Odd() {
	var a int
	fmt.Println("Введите число: ")
	fmt.Scanf("%d", &a)

	if a%2 == 0 {
		fmt.Printf("Число %d четное", a)
	} else {
		fmt.Printf("Число %d нечетное", a)
	}
}

func checkSimbol(a float64) string {

	if a < 0 {
		return "Отрицательное"
	} else if a > 0 {
		return "Положительное"
	}
	return "Ноль"
}

func cycle() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func lenghtString(s string) int {
	lenght := 0
	for range s {
		lenght++
	}
	return lenght
}

func areaRectangle(r Rectangle) float64 {
	return r.Lenght * r.Height
}

type Rectangle struct {
	Lenght, Height float64
}

func average(x, y int) float64 {
	return float64(x+y) / 2
}
