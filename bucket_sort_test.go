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
	Sort(slice)
	t.Log(slice)
}
