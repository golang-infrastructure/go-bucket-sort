package bucket_sort

import (
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	slice := make([]int, 10)
	for index := range slice {
		slice[index] = rand.Intn(100)
		if rand.Int()%2 == 0 {
			slice[index] *= -1
		}
		//slice[index] *= -1
	}
	err := Sort(slice, 1000)
	t.Log(err)
	t.Log(slice)
}
