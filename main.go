package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	print("Hello world")

	uuid := uuid.New().String()

	fmt.Println(uuid)
}
