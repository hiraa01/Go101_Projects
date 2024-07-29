package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Rastgele bir sayı üretiyoruz
	randomNumber := rand.Intn(100) + 1
	var tahmin int

	fmt.Print("Rastgele sayıyı tahmin etme oyununa hoş geldiniz!\n1 ile 100 arasında bir sayı tahmin edin. ")
	for {
		fmt.Print("Tahmin: ")
		fmt.Scan(&tahmin)
		if tahmin < randomNumber {
			fmt.Println("Yanlış cevap, daha büyük!")
		} else if tahmin > randomNumber {
			fmt.Println("Yanlış cevap, daha küçük!")
		} else {
			fmt.Println("Doğru cevap!")
			break
		}
	}
}
