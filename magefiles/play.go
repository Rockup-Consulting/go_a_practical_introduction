package main

import (
	"fmt"
	"todo/core/randx"
)

func Rand() error {
	r, err := randx.String(32)
	if err != nil {
		return err
	}
	fmt.Println(r)

	return nil
}
