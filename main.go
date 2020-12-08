package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(n int) string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {

	for {
		currentTime := time.Now()
		fmt.Println(currentTime.Format("2006-01-02 15:04:05.000000000"), " - ", RandomString(75))
	}
}
