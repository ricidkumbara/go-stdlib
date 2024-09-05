package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestContainerList(t *testing.T) {
	data := list.New()
	data.PushBack("Ricid")
	data.PushBack("Kumbara")
	data.PushBack("Fulan")

	// var head *list.Element = data.Front()
	// fmt.Println(head.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
