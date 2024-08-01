package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var rand int = rand.Intn(10)
	var input int = 0
	var laps int = 1
	fmt.Print("Aklımdan 1-10 arasında bir sayı tuttum, bakalım bilebilecek misin (ipucu icin 0 yazın)\n")
	for true {
		fmt.Print("Bir sayı seç : ")
		fmt.Scan(&input)
		if input == 0 {
			if rand <= 5 {
				fmt.Printf("Sayı 5 ten küçüktür\n")
			} else {
				fmt.Printf("Sayı 5 ten büyüktür\n")
			}
		} else if input == rand {
			fmt.Printf("Tebrikler sayıyı %d adımda bildin\n", laps)
			break
		} else {
			fmt.Printf("Tekrar Dene\n")
		}
		laps++
	}
}
