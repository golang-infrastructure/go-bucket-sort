package main

import (
	"fmt"
	bucket_sort "github.com/golang-infrastructure/go-bucket-sort"
	"math/rand"
)

func main() {
	slice := make([]int, 10)
	for index := range slice {
		slice[index] = rand.Intn(100)
		if rand.Int()%2 == 0 {
			slice[index] *= -1
		}
		//slice[index] *= -1
	}
	err := bucket_sort.Sort(slice, 1000)
	fmt.Println(err)
	fmt.Println(slice)
	// Output:
	// <nil>
	// [-81 -62 -25 11 28 37 47 56 81 94]
}
