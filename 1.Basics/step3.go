package main

import "fmt"

func main() {
	a := 10
	b := 20

	total := a + b
	fmt.Println(total)

	modul := a % b // Harqanday kichik sonni katta songa modulini olsangto'g'ridan
	fmt.Println(modul) //--> 10 // tog'ri kichiksonning o'zi o'tadi

	num1 := 63
	num2 := 10

	module := num1 % num2
	fmt.Println(module)
//________________________________________________________________________________________________

    num := 0
	fmt.Println(num)
	num += 10 // num = num + 10 --> 10
	fmt.Println(num)
	num -= 10 // num = num - 10 --> 0
	fmt.Println(num)
	num *= 10 // num = num * 10 --> 0
	fmt.Println(num)
	num /= 10 // num / num  10--> 0
	fmt.Println(num)
	//__________________________________________________________

box := 10
box -= 5 // box = 10 - 5 --> 5  
box += 4 // box = 5 + 4 --> 9
box *= 2 // box = 9 * 2 --> 18
box /= 3 // box = 18 / 3 --> 6
fmt.Println(box)
// ______________________________________________________________

	number := 10
	number++		// number += bu ham bo'ladi
	fmt.Println(number)
	
	
	number--             // number -= 1 bu ham bo'ladi
	fmt.Println(number) // -->
//_______________________________________________________________

number += 5number += 5
fmt.Println(number) // --> 15

number -= 3
fmt.Println(number) // --> 12

number /= 2
fmt.Println(number) // --> 6

number *= 3
fmt.Println(number) // --> 18
}