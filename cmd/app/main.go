// Entrypoing of KVStore
package main

import (
	"fmt"

	"github.com/PhosFactum/KVStore/pkg/input"
)

func main() {
	fmt.Println("--- KVStore ---")
	fmt.Println("\n| Program for key-value store for data |")

	for {
		input, err := input.GetString()
		if err != nil {
			fmt.Println("error while writing command")
			continue
		}

		fmt.Println(input)
	}
}
