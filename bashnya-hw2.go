package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите ваше число: ")
	fmt.Scan(&num)
	for i := 0; num < 12307; i++ {
		if num < 0 {
			num *= -1
		}
		if num%7 == 0 {
			num *= 39
		}
		if num%9 == 0 {
			num *= 13
			num += 1
		} else {
			num += 2
			num *= 3
		}
		if num%13 == 0 {
			if num%9 == 0 {
				fmt.Println("service error")
				break
			}
		} else {
			num += 1
		}
	}
	fmt.Println("Итоговое число: ", num)
}
