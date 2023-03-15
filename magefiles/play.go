package main

import (
	"fmt"
	"todo/randx"

	"github.com/google/uuid"
)

func UUID() {
	fmt.Println(uuid.NewString())
}

func Rand() error {
	r, err := randx.String(32)
	if err != nil {
		return err
	}
	fmt.Println(r)

	return nil
}
