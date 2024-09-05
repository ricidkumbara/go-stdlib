package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"testing"
)

func TestContainerRing(t *testing.T) {
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
