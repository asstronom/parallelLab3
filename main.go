package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"example.com/lab3/blocking"
	"example.com/lab3/nonblocking"
	"example.com/lab3/single"
)

const (
	sizeMultiplier = 100000
	N              = 23
	multipleOf     = 17
	test           = 0
)

var (
	arraySize = sizeMultiplier * N
)

func FillArray(array []int64) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i := range array {
		array[i] = r.Int63()
	}
}

func main() {
	fmt.Println("start")
	if test == 1 {
		arraySize = 30
	}
	array := make([]int64, arraySize)
	FillArray(array)
	if test == 1 {
		fmt.Println(array)
	}

	start := time.Now()
	fmt.Println("single: ", single.CountMultiples(multipleOf, array))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("blocking: ", blocking.CountMultiples(multipleOf, array))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("nonblocking: ", nonblocking.CountMultiples(multipleOf, array))
	fmt.Println(time.Since(start))
	if single.CountMultiples(multipleOf, array) != nonblocking.CountMultiples(multipleOf, array) {
		log.Fatalln("error when comparing (nonblocking): ", single.CountMultiples(multipleOf, array), " != ", nonblocking.CountMultiples(multipleOf, array))
	}
	if single.CountMultiples(multipleOf, array) != blocking.CountMultiples(multipleOf, array) {
		log.Fatalln("error when comparing (blocking): ", single.CountMultiples(multipleOf, array), " != ", blocking.CountMultiples(multipleOf, array))
	}
}
