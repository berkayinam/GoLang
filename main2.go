package main

import (
	"fmt"
)

func main() {
	var sonuc int = 0
	var a int = 0
	var b int = 0
	var isaret string = ""
	results := make([]int, 3)
	var i int = 0

	for i < 3 {
		fmt.Println(i+1, ". İçin Bir sayı girin")
		fmt.Scanln(&a)
		fmt.Println(i+1, ". İçin ikincisini girin")
		fmt.Scanln(&b)
		fmt.Println("isaret girin + - * /")
		fmt.Scanln(&isaret)

		if isaret == "+" {
			sonuc = a + b
			results[i] = sonuc
			fmt.Printf("sonuc = %d\n", sonuc)
		} else if isaret == "-"{
			sonuc = a - b
			results[i] = sonuc
			fmt.Printf("Sonuc = %d\n", sonuc)
		}
		switch isaret {
		case "*":		
			sonuc = a * b
			results[i] = sonuc
			fmt.Printf("Sonuc = %d\n", sonuc)
		case "/":
			sonuc = a / b
			results[i] = sonuc
			fmt.Printf("Sonuc = %d\n", sonuc)
		}
		i++
	}
	i = 0;
	fmt.Print(results)
	for i < len(results) - 1 {
		results[0] += results[i + 1]
		i++
	}
	fmt.Print(results[0], " = Tüm girilen sayıların toplamı")
}
