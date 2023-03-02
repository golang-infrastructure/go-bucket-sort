package main

import (
	"fmt"
	bucket_sort "github.com/golang-infrastructure/go-bucket-sort"
	"math/rand"
)

type user struct {
	name string
	age  int
}

func main() {
	userSlice := make([]*user, 100)
	for i := 0; i < 100; i++ {
		userSlice[i] = &user{
			name: fmt.Sprintf("name-%d", i),
			age:  rand.Intn(99) + 1,
		}
	}

	err := bucket_sort.SortByFunc(userSlice, func(index int, value *user) int {
		return value.age
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, user := range userSlice {
		fmt.Println(fmt.Sprintf("age: %d, name: %s", user.age, user.name))
	}
	// Output:
	// age: 1, name: name-21
	// age: 1, name: name-75
	// age: 2, name: name-76
	// age: 4, name: name-96
	// age: 6, name: name-19
	// age: 6, name: name-32
	// age: 7, name: name-9
	// age: 8, name: name-6
	// age: 8, name: name-51
	// age: 8, name: name-55
	// age: 10, name: name-24
	// age: 11, name: name-23
	// age: 11, name: name-93
	// age: 12, name: name-13
	// age: 13, name: name-38
	// age: 14, name: name-4
	// age: 16, name: name-5
	// age: 16, name: name-61
	// age: 17, name: name-26
	// age: 17, name: name-46
	// age: 18, name: name-7
	// age: 18, name: name-15
	// age: 19, name: name-16
	// age: 21, name: name-2
	// age: 21, name: name-99
	// age: 22, name: name-12
	// age: 22, name: name-57
	// age: 23, name: name-30
	// age: 23, name: name-91
	// age: 24, name: name-0
	// age: 24, name: name-33
	// age: 24, name: name-59
	// age: 25, name: name-28
	// age: 26, name: name-85
	// age: 27, name: name-60
	// age: 29, name: name-27
	// age: 29, name: name-64
	// age: 31, name: name-34
	// age: 32, name: name-77
	// age: 33, name: name-22
	// age: 36, name: name-56
	// age: 37, name: name-94
	// age: 37, name: name-97
	// age: 38, name: name-83
	// age: 39, name: name-17
	// age: 41, name: name-54
	// age: 42, name: name-87
	// age: 43, name: name-92
	// age: 45, name: name-35
	// age: 45, name: name-65
	// age: 47, name: name-44
	// age: 48, name: name-3
	// age: 48, name: name-69
	// age: 48, name: name-90
	// age: 49, name: name-49
	// age: 52, name: name-31
	// age: 53, name: name-98
	// age: 54, name: name-67
	// age: 55, name: name-81
	// age: 56, name: name-47
	// age: 58, name: name-63
	// age: 60, name: name-10
	// age: 60, name: name-73
	// age: 61, name: name-53
	// age: 61, name: name-80
	// age: 62, name: name-37
	// age: 62, name: name-40
	// age: 62, name: name-78
	// age: 64, name: name-52
	// age: 65, name: name-89
	// age: 66, name: name-84
	// age: 68, name: name-14
	// age: 72, name: name-72
	// age: 72, name: name-79
	// age: 75, name: name-18
	// age: 75, name: name-20
	// age: 76, name: name-29
	// age: 77, name: name-95
	// age: 78, name: name-68
	// age: 79, name: name-1
	// age: 79, name: name-36
	// age: 79, name: name-74
	// age: 81, name: name-88
	// age: 82, name: name-39
	// age: 82, name: name-43
	// age: 82, name: name-45
	// age: 82, name: name-50
	// age: 83, name: name-42
	// age: 83, name: name-48
	// age: 83, name: name-70
	// age: 89, name: name-11
	// age: 89, name: name-25
	// age: 89, name: name-58
	// age: 89, name: name-82
	// age: 90, name: name-66
	// age: 91, name: name-71
	// age: 93, name: name-86
	// age: 94, name: name-62
	// age: 95, name: name-8
	// age: 95, name: name-41
}
