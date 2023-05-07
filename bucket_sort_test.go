package bucket_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

// ------------------------------------------------- --------------------------------------------------------------------

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
	assert.Nil(t, err)
	assert.True(t, IsSorted(slice))

	//t.Log(slice)
}

// ------------------------------------------------- --------------------------------------------------------------------

func TestSortDesc(t *testing.T) {
	slice := make([]int, 10)
	for index := range slice {
		slice[index] = rand.Intn(100)
		if rand.Int()%2 == 0 {
			slice[index] *= -1
		}
		//slice[index] *= -1
	}
	err := SortDesc(slice, 1000)
	assert.Nil(t, err)
	assert.True(t, IsSortedDesc(slice))

	//t.Log(slice)
}

// ------------------------------------------------ ---------------------------------------------------------------------

type user struct {
	name string
	age  int
}

func TestSortByFunc(t *testing.T) {
	userSlice := make([]*user, 1000)
	for i := 0; i < 1000; i++ {
		userSlice[i] = &user{
			name: fmt.Sprintf("name-%d", i),
			age:  rand.Intn(99) + 1,
		}
	}

	err := SortByFunc(userSlice, func(index int, value *user) int {
		return value.age
	})
	assert.Nil(t, err)
	for index := 1; index < len(userSlice); index++ {
		assert.True(t, userSlice[index].age >= userSlice[index-1].age)
	}

	//for _, user := range userSlice {
	//	t.Log(fmt.Sprintf("age: %d, name: %s\n", user.age, user.name))
	//}

}

// ------------------------------------------------ ---------------------------------------------------------------------

func IsSorted(slice []int) bool {
	for index := 1; index < len(slice); index++ {
		if slice[index] < slice[index-1] {
			return false
		}
	}
	return true
}

func IsSortedDesc(slice []int) bool {
	for index := 1; index < len(slice); index++ {
		if slice[index] > slice[index-1] {
			return false
		}
	}
	return true
}

// ------------------------------------------------- --------------------------------------------------------------------
