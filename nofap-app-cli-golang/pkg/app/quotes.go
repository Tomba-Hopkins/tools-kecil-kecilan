package app

import (
	"math/rand"
	"time"
)

func QuotesBiarSemangat() string {

	rand.Seed(time.Now().UnixNano())

	arr := []string{
		"Teruslah berjuang king lanjutkan usahamu itu☕",
		"Lanjutkann brother apresiasi dirimu🥤",
		"Mantap dan semangat terus bro🍛",
		"Kau tau, bahkan burger bisa lebih enak🍔",
		"Lanjutkan. pergi ke pluto jika kau menyerah! 🚀",
	}

	randIndex := rand.Intn(len(arr))

	res := arr[randIndex]
	// fmt.Printf("Result: %v", res)

	return res
}