package main

import (
	"fmt"
	"sort"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSort(t *testing.T) {
	users := []User{
		{"Ricid", 25},
		{"Kumbara", 26},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
