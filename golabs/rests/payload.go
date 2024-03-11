package main

import "strings"

const (
	sizeSmall  uint = 100
	sizeMedium uint = 1000
	sizeLarge  uint = 10000
)

func generateText(size uint) string {
	var sb strings.Builder
	for i := uint(0); i < size; i++ {
		_, err := sb.WriteString("a")
		if err != nil {
			panic(err)
		}
	}
	return sb.String()
}
