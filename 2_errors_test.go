package main

import (
	"errors"
	"fmt"
	"testing"
)

var (
	ErrValidationError = errors.New("validation error")
	ErrNotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ErrValidationError
	}

	if id != "ricid" {
		return ErrNotFoundError
	}

	return nil
}

func TestErrors(t *testing.T) {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ErrValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, ErrNotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
